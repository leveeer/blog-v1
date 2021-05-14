package models

import (
	_ "gorm.io/gorm"
	"time"
)

type ArticleTags struct {
	Id         int       `gorm:"primaryKey" json:"id"`
	ArticleId  int       `json:"article_id"`
	TagId      int       `json:"tag_id"`
	CreateTime time.Time `json:"create_time"`
	UpdateTime time.Time `json:"update_time"`
}

func (ArticleTags) TableName() string {
	return "tb_article_tags"
}
