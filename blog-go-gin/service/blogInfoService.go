package service

import (
	"blog-go-gin/common"
	"blog-go-gin/logging"
	"blog-go-gin/models"
	"blog-go-gin/models/vo"
	"errors"
	"strconv"
	"sync"
	"time"
)

var BlogInfoService = &blogInfoService{}

type blogInfoService struct {
	wg sync.WaitGroup
}

/*func (b *blogInfoService) GetBlogInfo() (*vo.BlogHomeInfoVo, error) {
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
	notice, _ = common.RedisUtil.Get(common.NOTICE)
	if notice == "" {
		notice = "博客Go语言版即将上线，敬请期待！"
		_ = common.RedisUtil.Set(common.NOTICE, notice)
	}
	//查询访问量
	viewsCountStr, _ := common.RedisUtil.Get(common.BlogViewsCount)
	if viewsCountStr == "" {
		_ = common.RedisUtil.Set(common.BlogViewsCount, strconv.Itoa(0))
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
}*/

func (b *blogInfoService) GetBlogInfo() (*vo.BlogHomeInfoVo, error) {
	result := make(map[string]interface{}, 6)
	done := make(chan struct{})
	// 新增阻塞chan
	errChan := make(chan error)
	pool := common.NewCoroutinePool(6)
	pool.AddGoroutine(func() {
		userInfo, err := models.GetUserInfoByID(common.BloggerId)
		if err != nil {
			errChan <- err
		}
		result["userInfo"] = userInfo
	})
	pool.AddGoroutine(func() {
		//查询分类数量
		categoryCount, err := models.GetCategoryCount()
		if err != nil {
			errChan <- err
		}
		result["categoryCount"] = categoryCount
	})

	pool.AddGoroutine(func() {
		//查询文章数量
		condition := "is_delete = ? AND is_publish = ?"
		articleCount, err := models.GetArticlesCountByCondition(condition, common.False, common.True)
		if err != nil {
			errChan <- err
		}
		result["articleCount"] = articleCount
	})

	pool.AddGoroutine(func() {
		//查询标签数量
		tagCount, err := models.GetTagCount()
		if err != nil {
			errChan <- err
		}
		result["tagCount"] = tagCount
	})
	pool.AddGoroutine(func() {
		notice, err := common.RedisUtil.Get(common.NOTICE)
		if err != nil {
			errChan <- err
		}
		if notice == "" {
			notice = "博客Go语言版即将上线，敬请期待！"
			if err := common.RedisUtil.Set(common.NOTICE, notice); err != nil {
				errChan <- err
			}
		}
		result["notice"] = notice
	})

	pool.AddGoroutine(func() {
		//查询访问量
		viewsCountStr, err := common.RedisUtil.Get(common.BlogViewsCount)
		if err != nil {
			errChan <- err
		}
		if viewsCountStr == "" {
			if err := common.RedisUtil.Set(common.BlogViewsCount, strconv.Itoa(0)); err != nil {
				errChan <- err
			}
			viewsCountStr = "0"
		}
		viewsCount, err := strconv.Atoi(viewsCountStr)
		if err != nil {
			errChan <- err
		}
		result["viewsCount"] = viewsCount
	})
	pool.Wait(done)
	for {
		select {
		// 错误快返回
		case err := <-errChan:
			logging.Logger.Error(err)
			return nil, err
		case <-done:
			return &vo.BlogHomeInfoVo{
				UserInfo:      result["userInfo"].(*models.UserInfo),
				ArticleCount:  result["articleCount"].(int64),
				CategoryCount: result["categoryCount"].(int64),
				TagCount:      result["tagCount"].(int64),
				Notice:        result["notice"].(string),
				ViewsCount:    result["viewsCount"].(int),
			}, nil
		//超时处理
		case <-time.After(500 * time.Millisecond):
			logging.Logger.Error(errors.New(common.GetMsg(common.ApiCallTimeout)))
			return nil, errors.New(common.GetMsg(common.ApiCallTimeout))
		}
	}
}
