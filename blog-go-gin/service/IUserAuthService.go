package service

import (
	pb "blog-go-gin/go_proto"
)

type IUserAuthService interface {
	GetLoginCode(username string) error
	Register(user *pb.User) error
	Login(user *pb.User) error
}
