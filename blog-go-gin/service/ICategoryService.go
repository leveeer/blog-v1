package service

import pb "blog-go-gin/go_proto"

type ICategoryService interface {
	GetCategories() ([]*pb.Category, error)
	GetAdminCategory(c *pb.CsCondition) (*pb.ScAdminCategories, error)
}
