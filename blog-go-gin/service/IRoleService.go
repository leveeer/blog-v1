package service

import pb "blog-go-gin/go_proto"

type IRoleService interface {
	GetAdminUsersRole() ([]*pb.ScUserRole, error)
	GetRoles(c *pb.CsCondition) (*pb.ScAdminRoles, error)
}
