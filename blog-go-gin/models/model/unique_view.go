package model

import (
	"blog-go-gin/dao"
)

type UniqueView struct {
	CreateTime int64 `gorm:"column:create_time;not null"`
	ID         int   `gorm:"column:id;primaryKey;unique;not null;autoIncrement"`
	ViewsCount int   `gorm:"column:views_count;not null"`
}

// TableName sets the insert table name for this struct type
func (model *UniqueView) TableName() string {
	return "tb_unique_view"
}

func AddUniqueView(m *UniqueView) error {
	return dao.Db.Debug().Save(m).Error
}

func DeleteUniqueViewByID(id int) (bool, error) {
	if err := dao.Db.Debug().Delete(&UniqueView{}, id).Error; err != nil {
		return false, err
	}
	return dao.Db.Debug().RowsAffected > 0, nil
}

func DeleteUniqueView(condition string, args ...interface{}) (int64, error) {
	if err := dao.Db.Debug().Where(condition, args...).Delete(&UniqueView{}).Error; err != nil {
		return 0, err
	}
	return dao.Db.Debug().RowsAffected, nil
}

func UpdateUniqueView(m *UniqueView) error {
	return dao.Db.Debug().Save(m).Error
}

func GetUniqueViewByID(id int) (*UniqueView, error) {
	var m UniqueView
	if err := dao.Db.Debug().First(&m, id).Error; err != nil {
		return nil, err
	}
	return &m, nil
}

func GetUniqueViews(condition string, args ...interface{}) ([]*UniqueView, error) {
	res := make([]*UniqueView, 0)
	if err := dao.Db.Debug().Where(condition, args...).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

func GetWeekUniqueViews() ([]*UniqueView, error) {
	res := make([]*UniqueView, 0)
	if err := dao.Db.Debug().Order("create_time DESC").Limit(7).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}
