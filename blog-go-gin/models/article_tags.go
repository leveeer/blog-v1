package models

import (
	"blog-go-gin/dao"
	_ "gorm.io/gorm"
	"time"
)

type ArticleTags struct {
	ID         int       `gorm:"column:id;primaryKey;unique;not null;autoIncrement" json:"id"`
	ArticleID  int       `gorm:"column:article_id;not null" json:"article_id"`
	TagID      int       `gorm:"column:tag_id;not null" json:"tag_id"`
	CreateTime time.Time `gorm:"column:create_time" json:"create_time"`
	UpdateTime time.Time `gorm:"column:update_time" json:"update_time"`
}

// TableName sets the insert table name for this struct type
func (model *ArticleTags) TableName() string {
	return "tb_article_tags"
}

func AddArticleTags(m *ArticleTags) error {
	return dao.Db.Save(m).Error
}

func DeleteArticleTagsByID(id int) (bool, error) {
	if err := dao.Db.Delete(&ArticleTags{}, id).Error; err != nil {
		return false, err
	}
	return dao.Db.RowsAffected > 0, nil
}

func DeleteArticleTags(condition string, args ...interface{}) (int64, error) {
	if err := dao.Db.Where(condition, args...).Delete(&ArticleTags{}).Error; err != nil {
		return 0, err
	}
	return dao.Db.RowsAffected, nil
}

func UpdateArticleTags(m *ArticleTags) error {
	return dao.Db.Save(m).Error
}

func GetArticleTagsByID(id int) (*ArticleTags, error) {
	var m ArticleTags
	if err := dao.Db.First(&m, id).Error; err != nil {
		return nil, err
	}
	return &m, nil
}

func GetArticleTagsList(condition string, args ...interface{}) ([]*ArticleTags, error) {
	res := make([]*ArticleTags, 0)
	if err := dao.Db.Where(condition, args...).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}
