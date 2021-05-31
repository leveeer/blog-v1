package service

import (
	"blog-go-gin/models/model"
	"blog-go-gin/models/page"
	"blog-go-gin/models/vo"
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

func (b *articleService) GetArticleById(id int) (*vo.ArticleVo, error) {
	//获取当前文章
	article, err := model.GetArticleByID(id)
	if err != nil {
		return nil, err
	}
	//获取前一篇文章
	lastArticle, err := model.GetLastOrNextArticle(id, "is_delete = ? AND is_publish = ? AND id < ?", "id DESC")
	if err != nil {
		return nil, err
	}
	//获取后一篇文章
	nextArticle, err := model.GetLastOrNextArticle(id, "is_delete = ? AND is_publish = ? AND id > ?", "id ASC")
	if err != nil {
		return nil, err
	}
	return &vo.ArticleVo{
		Article:              *article,
		LastArticle:          *lastArticle,
		NextArticle:          *nextArticle,
		RecommendArticleList: nil,
	}, nil
}
