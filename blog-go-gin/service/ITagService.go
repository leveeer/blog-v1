package service

import pb "blog-go-gin/go_proto"

type ITagService interface {
	GetTags() ([]*pb.Tag, error)
	GetAdminTags(c *pb.CsCondition) (*pb.ScAdminTags, error)
	AddOrUpdateTag(tag *pb.CsTag) error
	DeleteTag(ids *pb.CsDeleteTag) error
}
