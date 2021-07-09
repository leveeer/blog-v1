package service

import (
	pb "blog-go-gin/go_proto"
	"blog-go-gin/models/page"
)

type IArticleService interface {
	GetArticleList(page page.IPage) ([]*pb.Article, error)
	GetArticleById(id int) (*pb.ArticleInfo, error)
	GetArchiveList(ipage *page.IPage) (*pb.Archives, error)
	GetArticleByCategoryID(categoryId int, iPage *page.IPage) (*pb.ArticlesByCategoryOrTag, error)
	GetArticleByTagID(tagId int, iPage *page.IPage) (*pb.ArticlesByCategoryOrTag, error)
	GetArticleOptions() (*pb.ScArticleOptions, error)
	UploadImage(filepath string) (string, error)
	AddArticle(article *pb.CsArticle) error
}
