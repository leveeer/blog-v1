package service

import pb "blog-go-gin/go_proto"

type IFriendLinkService interface {
	GetFriendLinks() ([]*pb.FriendLink, error)
}
