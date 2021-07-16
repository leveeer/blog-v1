package impl

import (
	"blog-go-gin/dao"
	pb "blog-go-gin/go_proto"
	"blog-go-gin/models/model"
	"blog-go-gin/models/page"
	"gorm.io/gorm"
	"sync"
	"time"
)

type TagServiceImpl struct {
	wg sync.WaitGroup
}

func (t *TagServiceImpl) DeleteTag(ids *pb.CsDeleteTag) error {
	err := dao.SqlTransaction(dao.Db.Begin(), func(tx *gorm.DB) error {
		_, err := model.DeleteTag(tx, "id in (?)", ids.TagIdList)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func (t *TagServiceImpl) AddOrUpdateTag(tag *pb.CsTag) error {
	if tag.Id == 0 {
		//新增
		err := dao.SqlTransaction(dao.Db.Begin(), func(tx *gorm.DB) error {
			err := model.AddTag(tx, &model.Tag{
				TagName:    tag.GetTagName(),
				CreateTime: time.Now().Unix(),
				Status:     1,
			})
			if err != nil {
				return err
			}
			return nil
		})
		if err != nil {
			return err
		}
		return nil
	}
	err := dao.SqlTransaction(dao.Db.Begin(), func(tx *gorm.DB) error {
		err := model.UpdateTag(tx, &model.Tag{
			ID:         int(tag.Id),
			TagName:    tag.TagName,
			UpdateTime: time.Now().Unix(),
		})
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func (t *TagServiceImpl) GetAdminTags(c *pb.CsCondition) (*pb.ScAdminTags, error) {
	tags, err := model.GetTagsByConditionWithPage(c.GetKeywords(), &page.IPage{Current: int(c.GetCurrent()), Size: int(c.GetSize())}, "%"+c.GetKeywords()+"%")
	if err != nil {
		return nil, err
	}
	var tagSlice []*pb.Tag
	for _, tag := range tags {
		tagSlice = append(tagSlice, &pb.Tag{
			Id:         int32(tag.ID),
			TagName:    tag.TagName,
			CreateTime: tag.CreateTime,
		})
	}
	tagCount, err := model.GetTagsCountByCondition(c.GetKeywords(), "%"+c.GetKeywords()+"%")
	if err != nil {
		return nil, err
	}
	return &pb.ScAdminTags{
		TagList: tagSlice,
		Count:   tagCount,
	}, nil
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
