package impl

import (
	"blog-go-gin/common"
	pb "blog-go-gin/go_proto"
	"blog-go-gin/logging"
	"blog-go-gin/models/model"
	"errors"
	"github.com/go-redis/redis/v8"
	"github.com/tal-tech/go-zero/core/mr"
	"strconv"
	"sync"
)

type BlogInfoServiceImpl struct {
	wg sync.WaitGroup
}

func NewBlogInfoServiceImpl() *BlogInfoServiceImpl {
	return &BlogInfoServiceImpl{}
}

func (b *BlogInfoServiceImpl) GetAbout() (*pb.About, error) {
	about, err := common.RedisUtil.Get(common.ABOUT)
	logging.Logger.Debug(about)

	if err != nil && errors.Is(err, redis.Nil) {
		about = "博客Go语言版即将上线，敬请期待！"
		if err = common.RedisUtil.Set(common.ABOUT, about); err != nil {
			logging.Logger.Error(err)
			return nil, err
		}
	} else if err != nil && !errors.Is(err, redis.Nil) {
		logging.Logger.Error(err)
		return nil, err
	}
	return &pb.About{
		Content: about,
	}, nil
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

/*func (b *BlogInfoServiceImpl) GetBlogInfo() (*vo.BlogHomeInfoVo, error) {
	result := make(map[string]interface{}, 6)
	done := make(chan struct{})
	// 新增阻塞chan
	errChan := make(chan error)
	pool := common.NewCoroutinePool(6)
	pool.AddGoroutine(func() {
		userInfo, err := model.GetUserInfoByID(common.BloggerId)
		if err != nil {
			errChan <- err
		}
		result["userInfo"] = userInfo
	})
	pool.AddGoroutine(func() {
		//查询分类数量
		categoryCount, err := model.GetCategoryCount()
		if err != nil {
			errChan <- err
		}
		result["categoryCount"] = categoryCount
	})

	pool.AddGoroutine(func() {
		//查询文章数量
		condition := "is_delete = ? AND is_publish = ?"
		articleCount, err := model.GetArticlesCountByCondition(condition, common.False, common.True)
		if err != nil {
			errChan <- err
		}
		result["articleCount"] = articleCount
	})

	pool.AddGoroutine(func() {
		//查询标签数量
		tagCount, err := model.GetTagCount()
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
				UserInfo:      result["userInfo"].(*model.UserInfo),
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
}*/

func (b *BlogInfoServiceImpl) GetBlogInfo() (*pb.BlogHomeInfo, error) {

	var userInfo *model.UserInfo
	var categoryCount int64
	var articleCount int64
	var tagCount int64
	var notice string
	var viewsCount int

	err := mr.Finish(
		func() (err error) {
			userInfo, err = model.GetUserInfoByID(common.BloggerId)
			if err != nil {
				return nil
			}
			return nil
		},
		func() (err error) {
			categoryCount, err = model.GetCategoryCount()
			if err != nil {
				return err
			}
			return nil
		},

		func() (err error) {
			condition := "is_delete = ? and is_publish = ?"
			articleCount, err = model.GetArticlesCountByCondition(condition, common.False, common.True)
			if err != nil {
				return err
			}
			return nil
		},
		func() (err error) {
			tagCount, err = model.GetTagCount()
			if err != nil {
				return err
			}
			return nil
		},
		func() (err error) {
			notice, err = common.RedisUtil.Get(common.NOTICE)
			if err != nil && errors.Is(err, redis.Nil) {
				notice = "博客Go语言版即将上线，敬请期待！"
				if err = common.RedisUtil.Set(common.NOTICE, notice); err != nil {
					logging.Logger.Error(err)
					return err
				}
				return nil
			} else if err != nil && !errors.Is(err, redis.Nil) {
				logging.Logger.Error(err)
				return err
			}
			return nil
		},

		func() (err error) {
			viewsCountStr, err := common.RedisUtil.Get(common.BlogViewsCount)
			if err != nil && errors.Is(err, redis.Nil) {
				if err := common.RedisUtil.Set(common.BlogViewsCount, strconv.Itoa(0)); err != nil {
					return err
				}
				viewsCountStr = "0"
			} else if err != nil && !errors.Is(err, redis.Nil) {
				logging.Logger.Error(err)
				return err
			}
			viewsCount, err = strconv.Atoi(viewsCountStr)
			if err != nil {
				return err
			}
			return nil
		},
	)
	if err != nil {
		return nil, err
	}

	user := &pb.UserInfo{
		Id:         int32(userInfo.ID),
		Email:      userInfo.Email,
		NickName:   userInfo.Nickname,
		Avatar:     userInfo.Avatar,
		Intro:      userInfo.Intro,
		Website:    userInfo.WebSite,
		CreateTime: userInfo.CreateTime,
		UpdateTime: userInfo.UpdateTime,
		IsDisable:  userInfo.IsDisable == 1,
	}
	return &pb.BlogHomeInfo{
		UserInfo:      user,
		ArticleCount:  articleCount,
		CategoryCount: categoryCount,
		TagCount:      tagCount,
		Notice:        notice,
		ViewCount:     int64(viewsCount),
	}, nil
}
