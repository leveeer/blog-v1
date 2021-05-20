package service

import (
	"blog-go-gin/common"
	"blog-go-gin/logging"
	"blog-go-gin/models"
	"blog-go-gin/models/vo"
	"strconv"
	"sync"
)

var BlogInfoService = &blogInfoService{}

type blogInfoService struct {
	wg sync.WaitGroup
}

func (b *blogInfoService) GetBlogInfo() (*vo.BlogHomeInfoVo, error) {
	//查询博主信息
	userInfo, err := models.GetUserInfoByID(common.BloggerId)
	if err != nil {
		return nil, err
	}
	//查询文章数量
	articles, err := models.GetArticles("is_delete = ? AND is_publish = ?", common.False, common.True)
	if err != nil {
		return nil, err
	}
	articleCount := int64(len(articles))
	//查询分类数量
	categoryCount, err := models.GetCategoryCount()
	if err != nil {
		return nil, err
	}
	//查询标签数量
	tagCount, err := models.GetTagCount()
	if err != nil {
		return nil, err
	}
	//查询公告
	var notice string
	notice = common.RedisUtil.Get(common.NOTICE)
	if notice == "" {
		notice = "博客Go语言版即将上线，敬请期待！"
		common.RedisUtil.Set(common.NOTICE, notice)
	}
	//查询访问量
	viewsCountStr := common.RedisUtil.Get(common.BlogViewsCount)
	if viewsCountStr == "" {
		common.RedisUtil.Set(common.BlogViewsCount, strconv.Itoa(0))
		viewsCountStr = "0"
	}
	viewsCount, err := strconv.Atoi(viewsCountStr)
	if err != nil {
		logging.Logger.Errorf("strconv Atoi err: %s", err)
	}

	//数据封装
	return &vo.BlogHomeInfoVo{
		UserInfo:      userInfo,
		ArticleCount:  articleCount,
		CategoryCount: categoryCount,
		TagCount:      tagCount,
		Notice:        notice,
		ViewsCount:    viewsCount,
	}, nil

}
