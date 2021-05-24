package vo

import (
	"blog-go-gin/models/model"
)

type BlogHomeInfoVo struct {
	UserInfo      *model.UserInfo `json:"user_info"`
	ArticleCount  int64           `json:"article_count"`
	CategoryCount int64           `json:"category_count"`
	TagCount      int64           `json:"tag_count"`
	Notice        string          `json:"notice"`
	ViewsCount    int             `json:"views_count"`
}
