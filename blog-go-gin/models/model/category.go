package model

import (
	"blog-go-gin/dao"
)

type Category struct {
	ID           int    `gorm:"column:id;primaryKey;unique;not null;autoIncrement" json:"id"`
	CategoryName string `gorm:"column:category_name;not null" json:"category_name"`
	CreateTime   int64  `gorm:"column:create_time" json:"create_time"`
	UpdateTime   int64  `gorm:"column:update_time" json:"update_time"`
}

// TableName sets the insert table name for this struct type
func (model *Category) TableName() string {
	return "tb_category"
}

func AddCategory(m *Category) error {
	return dao.Db.Save(m).Error
}

func DeleteCategoryByID(id int) (bool, error) {
	if err := dao.Db.Delete(&Category{}, id).Error; err != nil {
		return false, err
	}
	return dao.Db.RowsAffected > 0, nil
}

func DeleteCategory(condition string, args ...interface{}) (int64, error) {
	if err := dao.Db.Where(condition, args...).Delete(&Category{}).Error; err != nil {
		return 0, err
	}
	return dao.Db.RowsAffected, nil
}

func UpdateCategory(m *Category) error {
	return dao.Db.Save(m).Error
}

func GetCategoryByID(id int) (*Category, error) {
	var m Category
	if err := dao.Db.First(&m, id).Error; err != nil {
		return nil, err
	}
	return &m, nil
}

func GetCategories(condition string, args ...interface{}) ([]*Category, error) {
	res := make([]*Category, 0)
	if err := dao.Db.Where(condition, args...).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

func GetCategoryCount() (count int64, err error) {
	var categoryCount int64
	if err := dao.Db.Debug().Table("tb_category").Count(&categoryCount).Error; err != nil {
		return 0, err
	}
	return categoryCount, nil
}
