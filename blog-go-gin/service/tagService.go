package service

import (
	"blog-go-gin/dao"
	"blog-go-gin/models/model"
	"blog-go-gin/models/page"
)

func GetTagList(page page.IPage) []model.Tag {
	var tagList []model.Tag

	dao.Db.Debug().Table("t_tag").Limit(page.Size).Offset(page.Current - 1).Find(&tagList)

	return tagList
}
