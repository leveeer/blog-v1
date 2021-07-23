package service

import (
	pb "blog-go-gin/go_proto"
)

type IUserInfoService interface {
	UpdateUserStatus(userStatus *pb.CsUserStatus) error
}
