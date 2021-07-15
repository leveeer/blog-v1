package model

import (
	"blog-go-gin/dao"
	"blog-go-gin/models/page"
	"gorm.io/gorm"
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

func AddTag(tx *gorm.DB, m *Tag) error {
	return tx.Debug().Save(m).Error
}

func DeleteTagByID(tx *gorm.DB, id int) (bool, error) {
	if err := tx.Debug().Delete(&Tag{}, id).Error; err != nil {
		return false, err
	}
	return tx.RowsAffected > 0, nil
}

func DeleteTag(tx *gorm.DB, condition string, args ...interface{}) (int64, error) {
	if err := tx.Debug().Where(condition, args...).Delete(&Tag{}).Error; err != nil {
		return 0, err
	}
	return tx.RowsAffected, nil
}

func UpdateTag(tx *gorm.DB, m *Tag) error {
	return tx.Debug().Select("*").Omit("createTime", "status").Updates(m).Error
}

func GetTagByID(id int) (*Tag, error) {
	var m Tag
	if err := dao.Db.Debug().First(&m, id).Error; err != nil {
		return nil, err
	}
	return &m, nil
}

func GetTags(condition string, args ...interface{}) ([]*Tag, error) {
	res := make([]*Tag, 0)
	if err := dao.Db.Debug().Where(condition, args...).Find(&res).Error; err != nil {
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

func GetTagsByConditionWithPage(condition string, iPage *page.IPage, args ...interface{}) ([]*Tag, error) {
	res := make([]*Tag, 0)
	db := dao.Db
	if condition != "" {
		db = db.Where("tag_name LIKE ?", args...)
	}
	if err := db.Debug().Scopes(page.Paginate(iPage)).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

func GetTagsCountByCondition(condition string, args ...interface{}) (int64, error) {
	var count int64
	db := dao.Db
	if condition != "" {
		db = db.Where("tag_name LIKE ?", args...)
	}
	if err := db.Debug().Table("tb_tag").Count(&count).Error; err != nil {
		return int64(0), err
	}
	return count, nil
}
