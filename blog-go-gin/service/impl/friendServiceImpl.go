package impl

import (
	pb "blog-go-gin/go_proto"
	"blog-go-gin/models/model"
	"sync"
)

type FriendLinkServiceImpl struct {
	wg sync.WaitGroup
}

func (f *FriendLinkServiceImpl) GetFriendLinks() ([]*pb.FriendLink, error) {
	friendLinks, err := model.GetFriendLinks("1 = 1")
	if err != nil {
		return nil, err
	}
	var links []*pb.FriendLink
	for _, link := range friendLinks {
		links = append(links, &pb.FriendLink{
			Id:          int32(link.ID),
			LinkName:    link.LinkName,
			LinkAvatar:  link.LinkAvatar,
			LinkAddress: link.LinkAddress,
			LinkIntro:   link.LinkIntro,
			CreateTime:  link.CreateTime,
		})
	}
	return links, err
}
