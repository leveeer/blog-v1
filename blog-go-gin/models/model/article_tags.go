package model

import (
	"blog-go-gin/dao"
	"gorm.io/gorm"
	_ "gorm.io/gorm"
)

type ArticleTags struct {
	ID         int   `gorm:"column:id;primaryKey;unique;not null;autoIncrement" json:"id"`
	ArticleID  int   `gorm:"column:article_id;not null" json:"article_id"`
	TagID      int   `gorm:"column:tag_id;not null" json:"tag_id"`
	CreateTime int64 `gorm:"column:create_time" json:"create_time"`
	UpdateTime int64 `gorm:"column:update_time" json:"update_time"`
}

// TableName sets the insert table name for this struct type
func (model *ArticleTags) TableName() string {
	return "tb_article_tags"
}

func AddArticleTags(tx *gorm.DB, m *ArticleTags) error {
	return tx.Debug().Save(m).Error
}

func DeleteArticleTagsByID(id int) (bool, error) {
	if err := dao.Db.Delete(&ArticleTags{}, id).Error; err != nil {
		return false, err
	}
	return dao.Db.RowsAffected > 0, nil
}

func DeleteArticleTags(tx *gorm.DB, condition string, args ...interface{}) (int64, error) {
	if err := tx.Debug().Where(condition, args...).Delete(&ArticleTags{}).Error; err != nil {
		return 0, err
	}
	return tx.RowsAffected, nil
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
