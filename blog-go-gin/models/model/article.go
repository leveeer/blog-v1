package model

import (
	"blog-go-gin/common"
	"blog-go-gin/dao"
	"blog-go-gin/models/page"
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
	CategoryName   string `json:"category_name"`
}

// TableName sets the insert table name for this struct type
func (model *Article) TableName() string {
	return "tb_article"
}

func AddArticle(m *Article) error {
	return dao.Db.Table("tb_article").Save(m).Error
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

func UpdateArticle(m *Article) error {
	return dao.Db.Table("tb_article").Save(m).Error
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

func GetArticlesByPage(iPage page.IPage) ([]*Article, error) {
	res := make([]*Article, 0)
	if err := dao.Db.Debug().Table("tb_article").Scopes(page.Paginate(&iPage)).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

func GetArticlesOnHome(iPage page.IPage) ([]*Article, error) {
	res := make([]*Article, 0)
	if err := dao.Db.Debug().Table("tb_article").
		Select("tb_article.*,tb_category.category_name").
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
