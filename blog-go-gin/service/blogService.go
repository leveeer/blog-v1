package service

import (
	"blog-go-gin/dao"
	"blog-go-gin/models"
	"blog-go-gin/models/vo"
	"sync"
)

var BlogService = &blogService{}

type blogService struct {
	wg sync.WaitGroup
}

func GetBlogList(blogVO vo.BlogVO) []models.Blog {
	//return models.GetBlogList(blogVO)

	var blogList []models.Blog

	dao.Db.Debug().Table("t_blog").Limit(blogVO.PageSize).Offset(blogVO.CurrentPage - 1).Find(&blogList)

	return blogList
}

func (b *blogService) GetArticleByUid(uid string) models.Blog {
	var blog models.Blog
	var blogTags []models.BlogTags
	var tags []models.Tag
	b.wg.Add(2)
	go func() {
		dao.Db.Debug().Table("t_blog").Where("uid = ?", uid).First(&blog)
		b.wg.Done()
	}()
	go func() {
		dao.Db.Debug().Table("t_blog_tags").Where("blog_id = ?", uid).Find(&blogTags)
		b.wg.Done()
	}()
	b.wg.Wait()
	for _, tagTag := range blogTags {
		var tag models.Tag
		dao.Db.Debug().Table("t_tag").Where("uid = ?", tagTag.TagUid).Find(&tag)
		tags = append(tags, tag)
	}
	blog.TagList = append(blog.TagList, tags...)
	return blog
}
