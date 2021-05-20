package service

import (
	"blog-go-gin/models"
	"blog-go-gin/models/page"
	"sync"
)

var ArticleService = &articleService{}

type articleService struct {
	wg sync.WaitGroup
}

func (b *articleService) GetArticleList(page page.IPage) ([]*models.Article, error) {
	articles, err := models.GetArticlesByPage(page)
	if err != nil {
		return nil, err
	}
	return articles, err
}

func (b *articleService) GetArticleByUid(uid string) models.Article {
	return models.Article{}
}
