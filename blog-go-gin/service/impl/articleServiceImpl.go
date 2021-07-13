package impl

import (
	"blog-go-gin/common"
	"blog-go-gin/dao"
	pb "blog-go-gin/go_proto"
	"blog-go-gin/logging"
	"blog-go-gin/models/model"
	"blog-go-gin/models/page"
	"fmt"
	"gorm.io/gorm"
	"sync"
	"time"
)

type ArticleServiceImpl struct {
	wg sync.WaitGroup
}

func (b *ArticleServiceImpl) UpdateArticle(csArticle *pb.CsArticle) error {
	err := dao.SqlTransaction(dao.Db.Begin(), func(tx *gorm.DB) error {
		//更新文章表
		var isTop int8
		if csArticle.IsTop {
			isTop = 1
		}
		var isPublish int8
		if csArticle.IsPublish {
			isPublish = 1
		}
		article := &model.Article{
			ID:             int(csArticle.Id),
			ArticleTitle:   csArticle.ArticleTitle,
			ArticleContent: csArticle.ArticleContent,
			ArticleCover:   csArticle.ArticleCover,
			CategoryID:     int(csArticle.CategoryId),
			IsTop:          isTop,
			IsPublish:      isPublish,
			UpdateTime:     time.Now().Unix(),
		}
		err := model.UpdateArticle(tx, article)
		if err != nil {
			return err
		}
		//更新文章-标签表
		//先删除该文章所有标签，再添加新的标签
		_, err = model.DeleteArticleTags(tx, "article_id = ?", csArticle.GetId())
		if err != nil {
			return err
		}
		for _, tagId := range csArticle.TagIdList {
			err := model.AddArticleTags(tx, &model.ArticleTags{
				ArticleID:  int(csArticle.GetId()),
				TagID:      int(tagId),
				CreateTime: time.Now().Unix(),
			})
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func (b *ArticleServiceImpl) GetUpdateArticleInfoById(id int) (*pb.ScArticleInfo, error) {
	article, err := model.GetArticleByID(id)
	if err != nil {
		return nil, err
	}

	tags, err := model.GetArticleTagsList("article_id = ?", article.ID)
	if err != nil {
		return nil, err
	}
	var tagIdList []int64
	for _, tag := range tags {
		tagIdList = append(tagIdList, int64(tag.TagID))
	}

	logging.Logger.Debug(tagIdList)
	return &pb.ScArticleInfo{
		Id:             int64(article.ID),
		ArticleTitle:   article.ArticleTitle,
		ArticleContent: article.ArticleContent,
		ArticleCover:   article.ArticleCover,
		CategoryId:     int64(article.CategoryID),
		IsTop:          int32(article.IsTop),
		IsPublish:      int32(article.IsPublish),
		TagIdList:      tagIdList,
	}, nil
}

func (b *ArticleServiceImpl) GetAdminArticle(csAdminArticle *pb.CsAdminArticles) (*pb.ScAdminArticle, error) {
	condition := "is_delete = ? AND is_publish = ? "
	if csAdminArticle.GetKeywords() != "" {
		condition = fmt.Sprintf(condition+"%s", "AND article_title LIKE ?")
	}
	articles, err := model.GetArticlesByConditionWithPage(condition,
		&page.IPage{Current: int(csAdminArticle.GetCurrent()), Size: int(csAdminArticle.GetSize())},
		csAdminArticle.GetIsDelete(),
		csAdminArticle.GetIsPublish(),
		"%"+csAdminArticle.GetKeywords()+"%")
	if err != nil {
		return nil, err
	}
	var articleSlice []*pb.ScAdminArticleList
	for _, article := range articles {
		var tagSlice []*pb.Tag
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
		a := &pb.ScAdminArticleList{
			Id:           int64(article.ID),
			ArticleTitle: article.ArticleTitle,
			CreateTime:   article.CreateTime,
			UpdateTime:   article.UpdateTime,
			IsTop:        int32(article.IsTop),
			IsPublish:    int32(article.IsPublish),
			IsDelete:     int32(article.IsDelete),
			ViewsCount:   int64(article.ClickCount),
			LikeCount:    int64(article.CollectCount),
			TagList:      tagSlice,
			CategoryName: article.CategoryName,
		}
		articleSlice = append(articleSlice, a)
	}

	articleCount, err := model.GetArticlesCountByCondition(condition, csAdminArticle.GetIsDelete(), csAdminArticle.GetIsPublish(), "%"+csAdminArticle.GetKeywords()+"%")
	if err != nil {
		return nil, err
	}

	return &pb.ScAdminArticle{
		ArticleList: articleSlice,
		Count:       int32(articleCount),
	}, nil
}

func (b *ArticleServiceImpl) AddArticle(csArticle *pb.CsArticle) error {
	err := dao.SqlTransaction(dao.Db.Begin(), func(tx *gorm.DB) error {
		//文章表
		var isTop int8
		if csArticle.IsTop {
			isTop = 1
		}
		var isPublish int8
		if csArticle.IsPublish {
			isPublish = 1
		}
		articleID, err := model.AddArticle(tx, &model.Article{
			ArticleTitle:   csArticle.ArticleTitle,
			ArticleContent: csArticle.ArticleContent,
			ArticleCover:   csArticle.ArticleCover,
			CategoryID:     int(csArticle.CategoryId),
			CreateTime:     time.Now().Unix(),
			UserID:         common.BloggerId,
			IsTop:          isTop,
			IsPublish:      isPublish,
			IsOriginal:     common.True,
		})
		if err != nil {
			return err
		}
		//文章-标签表
		for _, tagId := range csArticle.TagIdList {
			err := model.AddArticleTags(tx, &model.ArticleTags{
				ArticleID:  articleID,
				TagID:      int(tagId),
				CreateTime: time.Now().Unix(),
			})
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func (b *ArticleServiceImpl) UploadImage(filepath string) (string, error) {
	key, err := common.GetQiNiuUtil().UploadQiNiu(filepath)
	if err != nil {
		return "", err
	}
	return key, nil
}

func (b *ArticleServiceImpl) GetArticleOptions() (*pb.ScArticleOptions, error) {
	//TODO 可并行查询
	categories, err := model.GetCategories("1 = 1")
	if err != nil {
		return nil, err
	}
	var categoryList []*pb.Category
	for _, category := range categories {
		categoryList = append(categoryList, &pb.Category{
			Id:           int32(category.ID),
			CategoryName: category.CategoryName,
		})
	}
	tags, err := model.GetTags("status = ?", true)
	if err != nil {
		return nil, err
	}
	var tagList []*pb.Tag
	for _, tag := range tags {
		tagList = append(tagList, &pb.Tag{
			Id:      int32(tag.ID),
			TagName: tag.TagName,
		})
	}
	return &pb.ScArticleOptions{
		CategoryList: categoryList,
		TagList:      tagList,
	}, nil
}

func NewArticleServiceImpl() *ArticleServiceImpl {
	return &ArticleServiceImpl{}
}

func (b *ArticleServiceImpl) GetArticleList(page page.IPage) ([]*pb.Article, error) {
	articles, err := model.GetArticlesOnHome(page)
	if err != nil {
		return nil, err
	}
	var articleSlice []*pb.Article
	for _, article := range articles {
		var tagSlice []*pb.Tag
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
	return articleSlice, err
}

func (b *ArticleServiceImpl) GetArticleById(id int) (*pb.ArticleInfo, error) {
	//获取当前文章
	article, err := model.GetArticleByID(id)
	if err != nil {
		return nil, err
	}
	currentArticle := &pb.Article{
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
		CategoryName:   article.CategoryName,
	}
	//获取前一篇文章
	article, err = model.GetLastOrNextArticle(id, "is_delete = ? and is_publish = ? and id < ?", "id DESC")
	if err != nil {
		return nil, err
	}
	lastArticle := &pb.Article{
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
		CategoryName:   article.CategoryName,
	}
	//获取后一篇文章
	article, err = model.GetLastOrNextArticle(id, "is_delete = ? and is_publish = ? and id > ?", "id ASC")
	if err != nil {
		return nil, err
	}
	nextArticle := &pb.Article{
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
		CategoryName:   article.CategoryName,
	}

	//获取推荐文章
	recommendArticles, err := model.GetRecommendArticles(id)
	if err != nil {
		return nil, err
	}
	var recommendArticleSlice []*pb.Article
	for _, article := range recommendArticles {
		recommendArticle := &pb.Article{
			Id:           int32(article.ID),
			ArticleCover: article.ArticleCover,
			ArticleTitle: article.ArticleTitle,
			CreateTime:   article.CreateTime,
		}
		recommendArticleSlice = append(recommendArticleSlice, recommendArticle)
	}

	//获取最新文章
	latestArticles, err := model.GetLatestArticles()
	if err != nil {
		return nil, err
	}
	var latestArticleSlice []*pb.Article
	for _, article := range latestArticles {
		latestArticle := &pb.Article{
			Id:           int32(article.ID),
			ArticleCover: article.ArticleCover,
			ArticleTitle: article.ArticleTitle,
			CreateTime:   article.CreateTime,
		}
		latestArticleSlice = append(latestArticleSlice, latestArticle)
	}

	return &pb.ArticleInfo{
		Article:              currentArticle,
		LastArticle:          lastArticle,
		NextArticle:          nextArticle,
		RecommendArticleList: recommendArticleSlice,
		ArticleLatestList:    latestArticleSlice,
	}, nil
}

func (b *ArticleServiceImpl) GetArchiveList(ipage *page.IPage) (*pb.Archives, error) {
	articles, err := model.GetArchives(ipage)
	if err != nil {
		return nil, err
	}
	var archiveList []*pb.ArchiveArticleInfo
	for _, article := range articles {
		a := &pb.ArchiveArticleInfo{
			Id:           int32(article.ID),
			ArticleTitle: article.ArticleTitle,
			CreateTime:   article.CreateTime,
		}
		archiveList = append(archiveList, a)
	}

	condition := "is_delete = ? and is_publish = ?"
	articleCount, err := model.GetArticlesCountByCondition(condition, common.False, common.True)
	if err != nil {
		return nil, err
	}

	return &pb.Archives{
		ArchiveList: archiveList,
		Count:       int32(articleCount),
	}, nil
}

func (b *ArticleServiceImpl) GetArticleByCategoryID(categoryId int, iPage *page.IPage) (*pb.ArticlesByCategoryOrTag, error) {
	var articleSlice []*pb.Article
	articles, err := model.GetArticlesByConditionWithPage("category_id = ? AND is_delete = 0 AND is_publish = 1", iPage, categoryId)
	if err != nil {
		return nil, err
	}
	for _, article := range articles {
		var tagSlice []*pb.Tag
		tags, err := model.GetTagNameByArticleId(article.ID)
		if err != nil {
			return nil, err
		}
		for _, tag := range tags {
			t := &pb.Tag{
				Id:         int32(tag.ID),
				TagName:    tag.TagName,
				ClickCount: int64(tag.ClickCount),
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

	category, err := model.GetCategoryByID(categoryId)
	if err != nil {
		return nil, err
	}
	return &pb.ArticlesByCategoryOrTag{
		ArticleList: articleSlice,
		Name:        category.CategoryName,
	}, nil
}

func (b *ArticleServiceImpl) GetArticleByTagID(tagId int, iPage *page.IPage) (*pb.ArticlesByCategoryOrTag, error) {
	var articleSlice []*pb.Article
	articles, err := model.GetArticlesByTagIdWithPage(tagId, iPage)
	if err != nil {
		return nil, err
	}
	for _, article := range articles {
		logging.Logger.Debug(article)
		var tagSlice []*pb.Tag
		tags, err := model.GetTagNameByArticleId(article.ID)
		if err != nil {
			return nil, err
		}
		for _, tag := range tags {
			t := &pb.Tag{
				Id:         int32(tag.ID),
				TagName:    tag.TagName,
				ClickCount: int64(tag.ClickCount),
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

	tag, err := model.GetTagByID(tagId)
	if err != nil {
		return nil, err
	}
	return &pb.ArticlesByCategoryOrTag{
		ArticleList: articleSlice,
		Name:        tag.TagName,
	}, nil
}
