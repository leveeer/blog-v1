package router

import (
	"blog-go-gin/common"
	pb "blog-go-gin/go_proto"
	"blog-go-gin/logging"
	"blog-go-gin/models/enum"
	"github.com/golang/protobuf/proto"
	"github.com/gorilla/websocket"
	"github.com/segmentio/ksuid"
	"strconv"
	"time"
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 2046
)

var ClientMgr = NewClientManager()

// ClientManager 客户端管理
type ClientManager struct {
	//客户端 map 储存并管理所有的长连接client，在线的为true，不在的为false
	clients map[*Session]bool
	//web端发送来的的message我们用broadcast来接收，并最后分发给所有的client
	broadcast chan *pb.ResponsePkg
	//新创建的长连接client
	register chan *Session
	//新注销的长连接client
	unregister chan *Session
}

func NewClientManager() *ClientManager {
	return &ClientManager{
		clients:    make(map[*Session]bool),
		broadcast:  make(chan *pb.ResponsePkg),
		register:   make(chan *Session),
		unregister: make(chan *Session),
	}
}

func (manager *ClientManager) Register() {
	for {
		select {
		//如果有新的连接接入,就通过channel把连接传递给conn
		case session := <-manager.register:
			logging.Logger.Debug("收到新的连接：", session)
			//把客户端的连接设置为true
			manager.clients[session] = true
			logging.Logger.Debug(len(manager.clients))
			manager.send(&pb.ResponsePkg{
				ServerTime: time.Now().Unix(),
				ScChat: &pb.ScChat{
					Type:   uint32(enum.OnlineCount.GetChatType()),
					Online: uint32(len(manager.clients)),
				},
				Code: pb.ResultCode_SuccessOK,
			})
		//如果连接断开了
		case session := <-manager.unregister:
			//判断连接的状态，如果是true,就关闭send，删除连接client的值
			if _, ok := manager.clients[session]; ok {
				close(session.sendWs)
				delete(manager.clients, session)
			}
		//广播
		case message := <-manager.broadcast:
			//遍历已经连接的客户端，把消息发送给他们
			manager.send(message)
		}
	}
}

//定义客户端管理的send方法
func (manager *ClientManager) send(message *pb.ResponsePkg) {
	for session := range manager.clients {
		select {
		case session.sendWs <- message:
		default:
			close(session.sendWs)
			delete(manager.clients, session)
		}
	}
}

// Session is a middleman between the websocket connection and the hub.
type Session struct {
	Id string
	// 小心，为了减少copy，此对象不能复用。
	sendWs chan *pb.ResponsePkg
	// stopCh is an additional signal channel.
	// Its sender is the moderator goroutine shown
	// below, and its receivers are all senders
	// and receivers of dataCh.
	stopSendWs chan struct{}
	// The channel toStop is used to notify the
	// moderator to close the additional signal
	// channel (stopCh). Its senders are any senders
	// and receivers of dataCh, and its receiver is
	// the moderator goroutine shown below.
	// It must be a buffered channel.
	toStopSendWs chan string
	// The websocket connection.
	conn *websocket.Conn
	// client real ip
	clientIP string
}

func NewSession(conn *websocket.Conn) *Session {
	client := &Session{
		Id: ksuid.New().String(),
		//Router:       router,
		sendWs:       make(chan *pb.ResponsePkg, 64), // magic number
		stopSendWs:   make(chan struct{}),
		toStopSendWs: make(chan string, 1),
		conn:         conn,
		//clientIP:     common.GetIPAddress(r),
	}
	logging.Logger.Infof("upgrade a new client %v", client)

	// Allow collection of memory referenced by the caller by doing all work in
	// new goroutines.
	go client.moderator()
	go client.writePump()
	go client.readPump()
	logging.Logger.Infof("upgrade finish %v", client)
	ClientMgr.register <- client
	return client
}

func (p *Session) kick() {
	// 只有read goroutine 可以关闭
	p.CloseWs("kick")
}

func (p *Session) closeAndCleanup() {
	p.CloseWs("closeAndCleanup")
}

func (p *Session) readPump() {
	defer func(extras ...interface{}) {
		if err := recover(); err != nil {
			common.PrintPanicStack(extras)
			logging.Logger.Error("readPump recover from panic!!!", err)
		}
	}()
	defer p.closeAndCleanup()
	p.conn.SetReadLimit(maxMessageSize)
	_ = p.conn.SetReadDeadline(time.Now().Add(pongWait))
	p.conn.SetPongHandler(func(payload string) error {
		_ = p.conn.SetReadDeadline(time.Now().Add(pongWait))
		return nil
	})
	for {
		messageType, data, err := p.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				logging.Logger.Error("IsUnexpectedCloseError error : ", err)
				return
			}
			logging.Logger.Error("readPump ReadMessage error ", err)
			return
		}
		switch messageType {
		case websocket.BinaryMessage:
			MessageHandler(p, data)
		case websocket.TextMessage:
			//MessageHandler(p, (data))
			logging.Logger.Error("readPump receive TextMessage ", data)
		case websocket.CloseMessage:
			logging.Logger.Error("readPump receive CloseMessage ")
			return
		}
	}
}

func (p *Session) Send(message *pb.ResponsePkg) {
	for {
		// The try-receive operation here is to
		// try to exit the sender goroutine as
		// early as possible. Try-receive and
		// try-send select blocks are specially
		// optimized by the standard Go
		// compiler, so they are very efficient.
		select {
		case <-p.stopSendWs:
			return
		default:
		}
		// Even if stopCh is closed, the first
		// branch in this select block might be
		// still not selected for some loops
		// (and for ever in theory) if the send
		// to dataCh is also non-blocking. If
		// this is unacceptable, then the above
		// try-receive operation is essential.
		select {
		case <-p.stopSendWs:
			return
		case p.sendWs <- message:
			return
		}
	}
}

func (p *Session) CloseWs(who string) {
	select {
	case p.toStopSendWs <- who:
	default:
	}
}

// writePump pumps messages from the hub to the websocket connection.
//
// A goroutine running writePump is started for each connection. The
// application ensures that there is at most one writer to a connection by
// executing all writes from this goroutine.
func (p *Session) writePump() {
	defer func(extras ...interface{}) {
		if err := recover(); err != nil {
			common.PrintPanicStack(extras)
			logging.Logger.Error("writePump recover from panic!!!", err)
		}
	}()
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		_ = p.conn.Close()
		p.CloseWs("writePump")
	}()
	for {
		// Same as the sender goroutine, the
		// try-receive operation here is to
		// try to exit the receiver goroutine
		// as early as possible.
		select {
		case <-p.stopSendWs:
			return
		default:
		}

		// Even if stopCh is closed, the first
		// branch in this select block might be
		// still not selected for some loops
		// (and forever in theory) if the receive
		// from dataCh is also non-blocking. If
		// this is not acceptable, then the above
		// try-receive operation is essential.
		select {
		case <-p.stopSendWs:
			return
		case message, ok := <-p.sendWs:
			err := p.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok || err != nil {
				// The hub closed the channel.
				_ = p.conn.WriteMessage(websocket.CloseMessage, []byte{})
				logging.Logger.Error(err)
				return
			}
			buf, err := proto.Marshal(message)
			if err != nil {
				logging.Logger.Error(err)
				return
			}
			err = p.conn.WriteMessage(websocket.BinaryMessage, buf)
		case <-ticker.C:
			if err := p.conn.SetWriteDeadline(time.Now().Add(writeWait)); err != nil {
				// The hub closed the channel.
				logging.Logger.Error("client close SetWriteDeadline ", err)
				_ = p.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			payload := []byte(strconv.Itoa(int(time.Now().Unix())))
			logging.Logger.Info("发送心跳包:", payload)
			if err := p.conn.WriteMessage(websocket.PingMessage, payload); err != nil {
				logging.Logger.Error("client close WriteMessage PingMessage ", err)
				return
			}
		}
	}
}

func (p *Session) moderator() {
	stoppedBy := <-p.toStopSendWs
	logging.Logger.Infof("%v stop by %v", p.clientIP, stoppedBy)
	WorldMessageChan <- &ClientMessage{
		Conn: p,
		Cmd:  &pb.RequestPkg{CmdId: pb.CsId_CsLogout},
	}
	close(p.stopSendWs)
}

type ClientMessage struct {
	Conn *Session
	Cmd  *pb.RequestPkg
}

var WorldMessageChan chan *ClientMessage
var ClosedChan chan struct{}

func MessageHandler(p *Session, data []byte) {
	//ClosedChan = make(chan struct{})
	//WorldMessageChan = make(chan *ClientMessage, 64)
	defer func(extras ...interface{}) {
		if err := recover(); err != nil {
			common.PrintPanicStack(extras)
			logging.Logger.Error("cmdHandler recover from panic!!!", err)
		}
	}()

	pkg := &pb.RequestPkg{}
	err := proto.Unmarshal(data, pkg)
	if err != nil {
		logging.Logger.Error("MessageHandler pkg error", err)
		return
	}

	logging.Logger.Infof("MessageHandler - %s %d", p.Id, pkg.CmdId)
	switch pkg.CmdId {
	case pb.CsId_CsBeginIndex:
		logging.Logger.Debug(pkg)
	//case pb.CsId_CsChat:
	//	logging.Logger.Debug(pkg)
	default:
		select {
		case <-ClosedChan:
			logging.Logger.Info("closed ")
		case WorldMessageChan <- &ClientMessage{
			Conn: p,
			Cmd:  pkg,
		}:
			//logging.Logger.Debug("收到客户端的消息：", <-WorldMessageChan)
		}
	}
}
