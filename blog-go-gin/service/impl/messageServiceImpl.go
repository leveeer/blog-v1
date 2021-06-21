package impl

import (
	"blog-go-gin/dao"
	pb "blog-go-gin/go_proto"
	"blog-go-gin/models/model"
	"gorm.io/gorm"
	"sync"
)

type MessageServiceImpl struct {
	wg sync.WaitGroup
}

func (m *MessageServiceImpl) GetMessages() ([]*pb.Message, error) {
	messages, err := model.GetMessages("1 = 1")
	if err != nil {
		return nil, err
	}
	var messageSlice []*pb.Message
	for _, message := range messages {
		messageSlice = append(messageSlice, &pb.Message{
			Id:             int32(message.ID),
			IpAddress:      message.IPAddress,
			IpSource:       message.IPSource,
			Nickname:       message.Nickname,
			Avatar:         message.Avatar,
			MessageContent: message.MessageContent,
			Time:           int32(message.Time),
			CreateTime:     message.CreateTime,
		})
	}
	return messageSlice, err
}

func (m *MessageServiceImpl) AddMessage(message *pb.Message) error {
	m1 := &model.Message{
		ID:             int(message.Id),
		IPAddress:      message.IpAddress,
		IPSource:       message.IpSource,
		Nickname:       message.Nickname,
		Avatar:         message.Avatar,
		MessageContent: message.MessageContent,
		Time:           int8(message.Time),
		CreateTime:     message.CreateTime,
	}
	err := dao.SqlTransaction(dao.Db.Begin(), func(tx *gorm.DB) error {
		err := model.AddMessage(tx, m1)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}
