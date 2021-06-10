package model

import (
	"blog-go-gin/dao"
	_ "gorm.io/gorm"
)

type Tag struct {
	ID         int    `gorm:"column:id;primaryKey;unique;not null;autoIncrement" json:"id"`
	TagName    string `gorm:"column:tag_name;not null" json:"tag_name"`
	CreateTime int64  `gorm:"column:create_time" json:"create_time"`
	UpdateTime int64  `gorm:"column:update_time" json:"update_time"`
	Status     int8   `gorm:"column:status;not null" json:"status"`
	ClickCount int    `gorm:"column:click_count;not null" json:"click_count"`
}

// TableName sets the insert table name for this struct type
func (model *Tag) TableName() string {
	return "tb_tag"
}

func AddTag(m *Tag) error {
	return dao.Db.Save(m).Error
}

func DeleteTagByID(id int) (bool, error) {
	if err := dao.Db.Delete(&Tag{}, id).Error; err != nil {
		return false, err
	}
	return dao.Db.RowsAffected > 0, nil
}

func DeleteTag(condition string, args ...interface{}) (int64, error) {
	if err := dao.Db.Where(condition, args...).Delete(&Tag{}).Error; err != nil {
		return 0, err
	}
	return dao.Db.RowsAffected, nil
}

func UpdateTag(m *Tag) error {
	return dao.Db.Save(m).Error
}

func GetTagByID(id int) (*Tag, error) {
	var m Tag
	if err := dao.Db.First(&m, id).Error; err != nil {
		return nil, err
	}
	return &m, nil
}

func GetTags(condition string, args ...interface{}) ([]*Tag, error) {
	res := make([]*Tag, 0)
	if err := dao.Db.Where(condition, args...).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}
func GetTagCount() (count int64, err error) {
	var tagCount int64
	if err := dao.Db.Debug().Table("tb_tag").Count(&tagCount).Error; err != nil {
		return 0, err
	}
	return tagCount, nil
}

func GetTagNameByArticleId(articleId int) ([]*Tag, error) {
	res := make([]*Tag, 0)
	if err := dao.Db.Debug().Table("tb_tag").Select("tb_tag.*").
		Joins("left join tb_article_tags on tb_article_tags.tag_id = tb_tag.id").
		Where("tb_article_tags.article_id = ?", articleId).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}
