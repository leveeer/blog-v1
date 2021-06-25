package service

import pb "blog-go-gin/go_proto"

type IUserInfoService interface {
	GetUserInfoByUid(userId int) pb.UserInfo
}
