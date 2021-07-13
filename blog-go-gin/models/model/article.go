package model

import (
	"blog-go-gin/common"
	"blog-go-gin/dao"
	"blog-go-gin/models/page"
	"gorm.io/gorm"
)

type Article struct {
	ID             int    `gorm:"column:id;primaryKey;unique;not null;autoIncrement" json:"id"`
	UserID         int    `gorm:"column:user_id;not null" json:"user_id"`
	CategoryID     int    `gorm:"column:category_id;not null" json:"category_id"`
	ArticleCover   string `gorm:"column:article_cover;not null" json:"article_cover"`
	ArticleTitle   string `gorm:"column:article_title;not null" json:"article_title"`
	ArticleContent string `gorm:"column:article_content;not null" json:"article_content"`
	CreateTime     int64  `gorm:"column:create_time;not null" json:"create_time"`
	UpdateTime     int64  `gorm:"column:update_time;not null" json:"update_time"`
	IsTop          int8   `gorm:"column:is_top;not null" json:"is_top"`
	IsPublish      int8   `gorm:"column:is_publish;not null" json:"is_publish"`
	IsDelete       int8   `gorm:"column:is_delete;not null" json:"is_delete"`
	IsOriginal     int8   `gorm:"column:is_original;not null" json:"is_original"`
	ClickCount     int    `gorm:"column:click_count;not null" json:"click_count"`
	CollectCount   int    `gorm:"column:collect_count;not null" json:"collect_count"`
	Tags           []*Tag `gorm:"-" json:"tags"`
	CategoryName   string `gorm:"->"`
}

// TableName sets the insert table name for this struct type
func (model *Article) TableName() string {
	return "tb_article"
}

func AddArticle(tx *gorm.DB, m *Article) (int, error) {
	if err := tx.Debug().Table("tb_article").Create(m).Error; err != nil {
		return 0, err
	}
	return m.ID, nil
}

func DeleteArticleByID(id int) (bool, error) {
	if err := dao.Db.Table("tb_article").Delete(&Article{}, id).Error; err != nil {
		return false, err
	}
	return dao.Db.RowsAffected > 0, nil
}

func DeleteArticle(condition string, args ...interface{}) (int64, error) {
	if err := dao.Db.Table("tb_article").Where(condition, args...).Delete(&Article{}).Error; err != nil {
		return 0, err
	}
	return dao.Db.RowsAffected, nil
}

func UpdateArticle(tx *gorm.DB, m *Article) error {
	return tx.Debug().Table("tb_article").Save(m).Error
}

func UpdateArticleClickCount(articleId int) error {
	return dao.Db.Debug().Table("tb_article").Where("id = ?", articleId).Select("click_count").Update("click_count", gorm.Expr("click_count + ?", 1)).Error
}

func GetArticleByID(id int) (*Article, error) {
	var m Article
	if err := dao.Db.Debug().Table("tb_article").First(&m, id).Error; err != nil {
		return nil, err
	}
	return &m, nil
}

func GetArticles(condition string, args ...interface{}) ([]*Article, error) {
	res := make([]*Article, 0)
	if err := dao.Db.Debug().Table("tb_article").Where(condition, args...).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

func GetArticlesCountByCondition(condition string, args ...interface{}) (int64, error) {
	var m Article
	var count int64
	if err := dao.Db.Debug().Table("tb_article").Where(condition, args...).Find(&m).Count(&count).Error; err != nil {
		return int64(0), err
	}
	return count, nil
}

func GetArticlesByConditionWithPage(condition string, iPage *page.IPage, args ...interface{}) ([]*Article, error) {
	res := make([]*Article, 0)
	if err := dao.Db.Debug().Table("tb_article").Select("tb_article.*,tb_category.category_name").
		Where(condition, args...).Joins("left join tb_category on tb_category.id = tb_article.category_id").
		Scopes(page.Paginate(iPage)).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

func GetArticlesByTagIdWithPage(tagId int, iPage *page.IPage) ([]*Article, error) {
	res := make([]*Article, 0)
	if err := dao.Db.Debug().Table("tb_article").Select("tb_article.id, article_cover, article_title, tb_article.create_time, category_id,category_name").
		Joins("JOIN tb_category c ON tb_article.category_id = c.id").
		Where("tb_article.id IN (SELECT article_id FROM tb_article_tags WHERE tag_id = ?) AND is_delete = 0 AND is_publish = 1", tagId).
		Scopes(page.Paginate(iPage)).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

func GetArticlesOnHome(iPage page.IPage) ([]*Article, error) {
	res := make([]*Article, 0)
	if err := dao.Db.Debug().Table("tb_article").
		Select("tb_article.*,tb_category.category_name").
		Where("is_delete = 0 AND is_publish = 1").
		Joins("left join tb_category on tb_category.id = tb_article.category_id").
		Scopes(page.Paginate(&iPage)).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

func GetLastOrNextArticle(id int, condition string, orderValue string) (*Article, error) {
	var m Article
	if err := dao.Db.Debug().Table("tb_article").Where(condition, common.False, common.True, id).
		Order(orderValue).Limit(1).Find(&m).Error; err != nil {
		return nil, err
	}
	return &m, nil
}

func GetRecommendArticles(id int) ([]*Article, error) {
	res := make([]*Article, 0)
	subQuery := dao.Db.Debug().Table("tb_article_tags").Select("tag_id").Where("tb_article_tags.article_id = ?", id)
	subQuery2 := dao.Db.Debug().Distinct("article_id").Table("(?) as t", subQuery).Where("article_id <> ?", id).Joins("left join tb_article_tags as t1 on t.tag_id = t1.tag_id")
	if err := dao.Db.Debug().Table("(?) as t2", subQuery2).Select("id,article_title,article_cover,create_time").
		Joins("left join tb_article as a on t2.article_id = a.id").Order("is_top DESC, id DESC").Limit(6).
		Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

func GetLatestArticles() ([]*Article, error) {
	res := make([]*Article, 0)
	if err := dao.Db.Debug().Table("tb_article").Select("id,article_title,article_cover,create_time").Where("is_delete = ? and is_publish = ?", false, true).Order("id DESC").Limit(5).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

func GetArchives(iPage *page.IPage) ([]*Article, error) {
	res := make([]*Article, 0)
	if err := dao.Db.Debug().Table("tb_article").Select("id,article_title,create_time").Scopes(page.Paginate(iPage)).
		Where("is_delete = ? and is_publish = ?", false, true).Order("create_time DESC").
		Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

func GetViewCountRank(rankNo int) ([]*Article, error) {
	res := make([]*Article, 0)
	if err := dao.Db.Debug().Where("is_delete = ? and is_publish = ?", false, true).
		Order("click_count DESC").Limit(rankNo).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}
