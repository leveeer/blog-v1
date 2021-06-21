package service

import pb "blog-go-gin/go_proto"

type ITagService interface {
	GetTags() ([]*pb.Tag, error)
}
