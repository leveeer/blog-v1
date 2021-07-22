package service

import (
	pb "blog-go-gin/go_proto"
)

type IMessageService interface {
	GetMessages() ([]*pb.Message, error)
	AddMessage(*pb.Message) error
	GetAdminMessages(c *pb.CsCondition) (*pb.ScAdminMessages, error)
	DeleteMessage(ids *pb.CsDeleteMessages) error
}
