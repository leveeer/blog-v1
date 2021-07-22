package impl

import (
	"blog-go-gin/dao"
	pb "blog-go-gin/go_proto"
	"blog-go-gin/models/model"
	"blog-go-gin/models/page"
	"gorm.io/gorm"
	"sync"
)

type MessageServiceImpl struct {
	wg sync.WaitGroup
}

func (m *MessageServiceImpl) DeleteMessage(ids *pb.CsDeleteMessages) error {
	err := dao.SqlTransaction(dao.Db.Begin(), func(tx *gorm.DB) error {
		_, err := model.DeleteMessage(tx, "id in (?)", ids.MessageIdList)
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

func (m *MessageServiceImpl) GetAdminMessages(c *pb.CsCondition) (*pb.ScAdminMessages, error) {
	messages, err := model.GetMessagesByConditionWithPage(c.GetKeywords(), &page.IPage{Current: int(c.GetCurrent()), Size: int(c.GetSize())}, "%"+c.GetKeywords()+"%")
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
			CreateTime:     message.CreateTime,
		})
	}
	messageCount, err := model.GetMessagesCountByCondition(c.GetKeywords(), "%"+c.GetKeywords()+"%")
	if err != nil {
		return nil, err
	}
	return &pb.ScAdminMessages{
		MessageList: messageSlice,
		Count:       messageCount,
	}, nil
}

func NewMessageServiceImpl() *MessageServiceImpl {
	return &MessageServiceImpl{}
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
