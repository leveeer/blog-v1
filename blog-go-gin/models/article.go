package models

import "time"

type Article struct {
	Id             int       `gorm:"primaryKey" json:"id"`
	UserId         int       `json:"user_id"`
	CategoryId     int       `json:"category_id"`
	ArticleCover   string    `json:"article_cover"`
	ArticleTitle   string    `json:"article_title"`
	ArticleContent string    `json:"article_content"`
	CreateTime     time.Time `json:"create_time"`
	UpdateTime     time.Time `json:"update_time"`
	IsTop          int8      `json:"is_top"`
	IsPublish      int8      `json:"is_publish"`
	IsDelete       int8      `json:"is_delete"`
	IsOriginal     int8      `json:"is_original"`
	ClickCount     int       `json:"click_count"`
	CollectCount   int       `json:"collect_count"`
}

func (*Article) TableName() string {
	return "tb_article"
}
