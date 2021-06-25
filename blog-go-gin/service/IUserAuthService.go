package service

import (
	pb "blog-go-gin/go_proto"
)

type IUserAuthService interface {
	GetLoginCode(username string) error
	Register(user *pb.User) error
	Login(user *pb.User) (bool, error)
	GetUserAuthByUsername(username string) (*pb.UserAuth, error)
	GetLoginResponse(username string) (*pb.LoginResponse, error)
}
