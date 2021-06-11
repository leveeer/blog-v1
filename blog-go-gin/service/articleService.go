package service

import (
	pb "blog-go-gin/go_proto"
	"blog-go-gin/models/model"
	"blog-go-gin/models/page"
	"blog-go-gin/models/vo"
	"sync"
)

type ArticleService struct {
	wg sync.WaitGroup
}

func (b *ArticleService) GetArticleList(page page.IPage) ([]*pb.Article, error) {
	articles, err := model.GetArticlesOnHome(page)
	var tagSlice []*pb.Tag
	var articleSlice []*pb.Article
	for _, article := range articles {
		tags, err := model.GetTagNameByArticleId(article.ID)
		if err != nil {
			return nil, err
		}
		for _, tag := range tags {
			t := &pb.Tag{
				Id:         int32(tag.ID),
				TagName:    tag.TagName,
				Status:     tag.Status == 1,
				ClickCount: int64(tag.ClickCount),
				CreateTime: tag.CreateTime,
				UpdateTime: tag.UpdateTime,
			}
			tagSlice = append(tagSlice, t)
		}
		a := &pb.Article{
			Id:             int32(article.ID),
			UserId:         int32(article.UserID),
			CategoryID:     int32(article.CategoryID),
			ArticleCover:   article.ArticleCover,
			ArticleTitle:   article.ArticleTitle,
			ArticleContent: article.ArticleContent,
			CreateTime:     article.CreateTime,
			UpdateTime:     article.UpdateTime,
			IsTop:          article.IsTop == 1,
			IsPublish:      article.IsPublish == 1,
			IsDelete:       article.IsDelete == 1,
			IsOriginal:     article.IsOriginal == 1,
			ClickCount:     int64(article.ClickCount),
			CollectCount:   int64(article.CollectCount),
			Tags:           tagSlice,
			CategoryName:   article.CategoryName,
		}
		articleSlice = append(articleSlice, a)
	}
	if err != nil {
		return nil, err
	}
	return articleSlice, err
}

func (b *ArticleService) GetArticleById(id int) (*vo.ArticleVo, error) {
	//获取当前文章
	article, err := model.GetArticleByID(id)
	if err != nil {
		return nil, err
	}
	//获取前一篇文章
	lastArticle, err := model.GetLastOrNextArticle(id, "is_delete = ? and is_publish = ? and id < ?", "id DESC")
	if err != nil {
		return nil, err
	}
	//获取后一篇文章
	nextArticle, err := model.GetLastOrNextArticle(id, "is_delete = ? and is_publish = ? and id > ?", "id ASC")
	if err != nil {
		return nil, err
	}
	return &vo.ArticleVo{
		Article:              article,
		LastArticle:          lastArticle,
		NextArticle:          nextArticle,
		RecommendArticleList: nil,
	}, nil
}
