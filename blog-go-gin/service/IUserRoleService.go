package service

import pb "blog-go-gin/go_proto"

type IUserRoleService interface {
	GetUserRoleAndUsername(userId int) (*pb.UserRole, error)
}
