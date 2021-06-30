package logic

import (
	"blog-go-gin/chat/router"
	"blog-go-gin/chat/utils"
	"blog-go-gin/common"
	pb "blog-go-gin/go_proto"
	"blog-go-gin/logging"
	"blog-go-gin/models/model"
	"errors"
	"fmt"
	"time"
)

type OnlineRouter struct {
	Users           map[uint32]*model.UserInfo
	UserBySessionId map[string]*model.UserInfo
}

var globalRouter OnlineRouter

func init() {
	logging.Logger.Debug("chat logic init...")
	router.WorldMessageChan = make(chan *router.ClientMessage, 64)
	router.ClosedChan = make(chan struct{})
	globalRouter.Users = make(map[uint32]*model.UserInfo)
	globalRouter.UserBySessionId = make(map[string]*model.UserInfo)
	common.GracefulWorkerAdd(1)
	go MessageDispatcher()
}

func MessageDispatcher() {
	// 初始化顺序不能乱
	defer func() {
		common.GracefulWorkerDone()
	}()
	logging.Logger.Info("[MessageDispatcher] running...")

	for {
		select {
		case msg := <-router.WorldMessageChan:
			err := CmdHandler(msg.Conn, msg.Cmd)
			if err != nil {
				logging.Logger.Error("cmd handler error", err)
				if msg.Conn != nil {
					msg.Conn.Send(&pb.ResponsePkg{
						ServerTime: time.Now().Unix(),
						Code:       pb.ResultCode_Fail,
						Message:    err.Error(),
					})
				}
			}
		case e := <-utils.TimerTick():
			utils.TimerExe(e)
		case <-router.ClosedChan:
			return
		default:
			time.Sleep(time.Millisecond * 32)
		}
	}
}

var handlers = map[pb.CsId]func(session *router.Session, p *model.UserInfo, pkg *pb.RequestPkg) error{
	pb.CsId_CsBeginIndex: func(session *router.Session, p *model.UserInfo, pkg *pb.RequestPkg) error {
		return nil
	},
	pb.CsId_CsChat: func(session *router.Session, p *model.UserInfo, pkg *pb.RequestPkg) error {
		return Chat(session, pkg.CsChatMessage)
	},
}

func CmdHandler(session *router.Session, pkg *pb.RequestPkg) (err error) {
	defer func(extras ...interface{}) {
		if err := recover(); err != nil {
			common.PrintPanicStack(extras)
			logging.Logger.Debug()
			logging.Logger.Error("cmdHandler recover from panic!!!", err)
			err = errors.New("服务器错误,请稍后重试~")
		}
	}()

	logging.Logger.Info("CmdHandler ", session.Id, pkg.CmdId)
	p := globalRouter.UserBySessionId[session.Id]
	f, ok := handlers[pkg.CmdId]
	if ok {
		return f(session, p, pkg)
	}
	return fmt.Errorf("CmdHandler missing %d", pkg.CmdId)
}

func Chat(session *router.Session, chatMessage *pb.CsChatMessage) error {
	logging.Logger.Debug("处理聊天信息", chatMessage)
	return nil
}
