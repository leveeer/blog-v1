package vo

import "blog-go-gin/models"

type BlogHomeInfoVo struct {
	UserInfo      models.UserInfo `json:"user_info"`
	ArticleCount  int64           `json:"article_count"`
	CategoryCount int64           `json:"category_count"`
	TagCount      int64           `json:"tag_count"`
	Notice        string          `json:"notice"`
	ViewsCount    int             `json:"views_count"`
}
