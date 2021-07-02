package router

import (
	"blog-go-gin/common"
	"blog-go-gin/config"
	"blog-go-gin/logging"
	"github.com/gorilla/websocket"
	"net/http"
	"strconv"
	"sync"
	"sync/atomic"
)

type GameInstance struct {
	stubId     string
	stubGameId string
	GameId     string
	GameType   string
}

type ChatServerInfo struct {
	ID        string
	RoomCount int32
	Ws        *config.WS
}

const (
	StatusNormal int = iota
	StatusShutdown
)

var ServerStatus = atomic.Value{}

type Router struct {
	*sync.RWMutex
	config      config.Config
	Sessions    map[string]*Session
	MessageChan chan *ClientMessage
}

//var globalRouter *Router

func init() {
	//globalRouter = &Router{
	//	RWMutex:           &sync.RWMutex{},
	//	Sessions:          make(map[string]*Session),
	//	MessageChan:       make(chan *ClientMessage, 128),
	//}
	ServerStatus.Store(StatusNormal)
}

var Upgrader = &websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	//允许跨域
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func RegisterGameServer(ws *config.WS) {
	info := ChatServerInfo{
		ID:        strconv.Itoa(int(ws.ID)),
		RoomCount: 0,
		Ws:        ws,
	}
	err := common.GetRedisUtil().HashSet(common.ChatServer, strconv.Itoa(int(ws.ID)), common.Serialization(info))
	if err != nil {
		logging.Entry.Fatal("RegisterGameServer error ", err)
		return
	}
}

func UnregisterGameServer(ws *config.WS) {
	err := common.GetRedisUtil().HashDel(common.ChatServer, strconv.Itoa(int(ws.ID)))
	if err != nil {
		logging.Entry.Error("UnregisterGameServer error ", err)
		return
	}
}

// HandleWs
// serveWs handles websocket requests from the peer.
func HandleWs(w http.ResponseWriter, r *http.Request) {
	if ServerStatus.Load().(int) == StatusShutdown {
		logging.Entry.Errorf("reject upgrade client %v", StatusShutdown)
		return
	}
	conn, err := Upgrader.Upgrade(w, r, nil)
	if err != nil {
		logging.Entry.Error(err)
		return
	}
	NewSession(conn)
	NewClientManager()
	//globalRouter.Lock()
	//globalRouter.Sessions[session.Id] = session
	//globalRouter.Unlock()
}
