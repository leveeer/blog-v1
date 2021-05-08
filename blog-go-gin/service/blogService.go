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

func GetBlogList(blogVO vo.BlogVO) []models.Blog {
	//return models.GetBlogList(blogVO)

	var blogList []models.Blog

	dao.Db.Debug().Table("t_blog").Limit(blogVO.PageSize).Offset(blogVO.CurrentPage - 1).Find(&blogList)

	return blogList
}

func GetArticleByUid(uid string) models.Blog {
	var blog models.Blog

	dao.Db.Debug().Table("t_blog").Where("uid = ?", uid).First(&blog)

	return blog
}
