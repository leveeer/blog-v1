package service

import (
	"blog-go-gin/models/model"
	"blog-go-gin/models/page"
	"sync"
)

var ArticleService = &articleService{}

type articleService struct {
	wg sync.WaitGroup
}

func (b *articleService) GetArticleList(page page.IPage) ([]*model.Article, error) {
	articles, err := model.GetArticlesOnHome(page)
	for _, article := range articles {
		tags, err := model.GetTagNameByArticleId(article.ID)
		if err != nil {
			return nil, err
		}
		article.Tags = append(article.Tags, tags...)
	}
	if err != nil {
		return nil, err
	}
	return articles, err
}

func (b *articleService) GetArticleByUid(uid string) model.Article {
	return model.Article{}
}
