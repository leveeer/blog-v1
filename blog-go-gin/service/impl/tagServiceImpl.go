package impl

import (
	pb "blog-go-gin/go_proto"
	"blog-go-gin/models/model"
	"sync"
)

type TagServiceImpl struct {
	wg sync.WaitGroup
}

func NewTagServiceImpl() *TagServiceImpl {
	return &TagServiceImpl{}
}

func (t *TagServiceImpl) GetTags() ([]*pb.Tag, error) {
	var tagList []*pb.Tag
	tags, err := model.GetTags("1 =1")
	if err != nil {
		return nil, err
	}
	for _, tag := range tags {
		tagList = append(tagList,
			&pb.Tag{
				Id:         int32(tag.ID),
				TagName:    tag.TagName,
				CreateTime: tag.CreateTime,
				UpdateTime: tag.UpdateTime,
				Status:     tag.Status == 1,
				ClickCount: int64(tag.ClickCount),
			},
		)
	}
	return tagList, nil
}
