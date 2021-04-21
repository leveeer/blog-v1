package service

import (
	"blog-go-gin/dao"
	"blog-go-gin/models"
	"blog-go-gin/models/vo"
	"sync"
)

type blogService struct {
	wg sync.WaitGroup
}

func GetBlogList(blogVO vo.BlogVO) []models.Blog{
	//return models.GetBlogList(blogVO)

	var blogList []models.Blog

	dao.Db.Find(&blogList)

	return blogList
}


