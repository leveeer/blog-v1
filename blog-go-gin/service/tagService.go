package service

import (
	"blog-go-gin/dao"
	"blog-go-gin/models"
	"blog-go-gin/models/page"
)

func GetTagList(page page.IPage) []models.Tag {
	var tagList []models.Tag

	dao.Db.Debug().Table("t_tag").Limit(page.Size).Offset(page.Current - 1).Find(&tagList)

	return tagList
}