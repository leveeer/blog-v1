package service

import pb "blog-go-gin/go_proto"

type IMenuService interface {
	GetUserMenus(roleid int) ([]*pb.ScUserMenuMessage, error)
}
