package service

import pb "blog-go-gin/go_proto"

type ICategoryService interface {
	GetCategories() ([]*pb.Category, error)
	GetAdminCategory(c *pb.CsCondition) (*pb.ScAdminCategories, error)
	AddOrUpdateCategory(category *pb.CsCategory) error
	DeleteCategory(ids *pb.CsDeleteCategory) error
}
