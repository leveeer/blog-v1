package models

import (
	"blog-go-gin/dao"
	"blog-go-gin/models/page"
	"time"
)

type Article struct {
	ID             int       `gorm:"column:id;primaryKey;unique;not null;autoIncrement" json:"id"`
	UserID         int       `gorm:"column:user_id;not null" json:"user_id"`
	CategoryID     int       `gorm:"column:category_id;not null" json:"category_id"`
	ArticleCover   string    `gorm:"column:article_cover;not null" json:"article_cover"`
	ArticleTitle   string    `gorm:"column:article_title;not null" json:"article_title"`
	ArticleContent string    `gorm:"column:article_content;not null" json:"article_content"`
	CreateTime     time.Time `gorm:"column:create_time;not null" json:"create_time"`
	UpdateTime     time.Time `gorm:"column:update_time;not null" json:"update_time"`
	IsTop          int8      `gorm:"column:is_top;not null" json:"is_top"`
	IsPublish      int8      `gorm:"column:is_publish;not null" json:"is_publish"`
	IsDelete       int8      `gorm:"column:is_delete;not null" json:"is_delete"`
	IsOriginal     int8      `gorm:"column:is_original;not null" json:"is_original"`
	ClickCount     int       `gorm:"column:click_count;not null" json:"click_count"`
	CollectCount   int       `gorm:"column:collect_count;not null" json:"collect_count"`
}

// TableName sets the insert table name for this struct type
func (model *Article) TableName() string {
	return "tb_article"
}

func AddArticle(m *Article) error {
	return dao.Db.Save(m).Error
}

func DeleteArticleByID(id int) (bool, error) {
	if err := dao.Db.Delete(&Article{}, id).Error; err != nil {
		return false, err
	}
	return dao.Db.RowsAffected > 0, nil
}

func DeleteArticle(condition string, args ...interface{}) (int64, error) {
	if err := dao.Db.Where(condition, args...).Delete(&Article{}).Error; err != nil {
		return 0, err
	}
	return dao.Db.RowsAffected, nil
}

func UpdateArticle(m *Article) error {
	return dao.Db.Save(m).Error
}

func GetArticleByID(id int) (*Article, error) {
	var m Article
	if err := dao.Db.First(&m, id).Error; err != nil {
		return nil, err
	}
	return &m, nil
}

func GetArticles(condition string, args ...interface{}) ([]*Article, error) {
	res := make([]*Article, 0)
	if err := dao.Db.Debug().Where(condition, args...).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

func GetArticlesByPage(ipage page.IPage) ([]*Article, error) {
	res := make([]*Article, 0)
	if err := dao.Db.Debug().Scopes(page.Paginate(&ipage)).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}
