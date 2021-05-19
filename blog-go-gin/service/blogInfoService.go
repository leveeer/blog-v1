package service

import (
	"blog-go-gin/common"
	"blog-go-gin/dao"
	"blog-go-gin/models"
	"blog-go-gin/models/vo"
	"log"
	"strconv"
	"sync"
)

var BlogInfoService = &blogInfoService{}

type blogInfoService struct {
	wg sync.WaitGroup
}

func (b *blogInfoService) GetBlogInfo() vo.BlogHomeInfoVo {
	//查询博主信息
	var userInfo models.UserInfo
	b.wg.Add(6)
	go func() {
		dao.Db.Debug().Table("tb_user_info").Where("id = ?", common.BloggerId).First(&userInfo)
		b.wg.Done()
	}()
	//查询文章数量
	var article models.Article
	var articleCount int64
	go func() {
		dao.Db.Debug().Table("tb_article").Where("is_delete = ? AND is_publish = ?", common.False, common.True).Find(&article).Count(&articleCount)
		b.wg.Done()
	}()
	//查询分类数量
	var categoryCount int64
	go func() {
		dao.Db.Debug().Table("tb_category").Count(&categoryCount)
		b.wg.Done()
	}()
	//查询标签数量
	var tagCount int64
	go func() {
		dao.Db.Debug().Table("tb_tag").Count(&tagCount)
		b.wg.Done()
	}()
	//查询公告
	var notice string
	go func() {
		notice = common.RedisUtil.Get(common.NOTICE)
		if notice == "" {
			notice = "博客Go语言版即将上线，敬请期待！"
			common.RedisUtil.Set(common.NOTICE, notice)
		}
		b.wg.Done()
	}()

	//查询访问量
	var viewsCount int
	var err error
	go func() {
		viewsCountStr := common.RedisUtil.Get(common.BlogViewsCount)
		if viewsCountStr == "" {
			common.RedisUtil.Set(common.BlogViewsCount, strconv.Itoa(0))
			viewsCountStr = "0"
		}
		viewsCount, err = strconv.Atoi(viewsCountStr)
		if err != nil {
			log.Println("strconv.Atoi err:", err)
		}
		b.wg.Done()
	}()

	b.wg.Wait()
	//数据封装
	return vo.BlogHomeInfoVo{
		UserInfo:      userInfo,
		ArticleCount:  articleCount,
		CategoryCount: categoryCount,
		TagCount:      tagCount,
		Notice:        notice,
		ViewsCount:    viewsCount,
	}

}
