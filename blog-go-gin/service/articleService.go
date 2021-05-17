package service

import (
	"blog-go-gin/dao"
	"blog-go-gin/models"
	"blog-go-gin/models/page"
	"sync"
)

var ArticleService = &articleService{}

type articleService struct {
	wg sync.WaitGroup
}

func (b *articleService) GetArticleList(page page.IPage) []models.Article {
	var articles []models.Article
	dao.Db.Debug().Scopes(page.Paginate(&page)).Find(&articles)
	return articles
}

func (b *articleService) GetArticleByUid(uid string) models.Article {
	return models.Article{}
}
