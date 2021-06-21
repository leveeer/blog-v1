package service

import pb "blog-go-gin/go_proto"

type IBlogInfoService interface {
	GetBlogInfo() (*pb.BlogHomeInfo, error)
	GetAbout() (*pb.About, error)
}
