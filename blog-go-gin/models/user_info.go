package models

import (
	"blog-go-gin/dao"
	"time"
)

type UserInfo struct {
	ID         int       `gorm:"column:id;primaryKey;unique;not null;autoIncrement" json:"id"`
	Email      string    `gorm:"column:email;not null" json:"email"`
	Nickname   string    `gorm:"column:nickname;not null" json:"nickname"`
	Avatar     string    `gorm:"column:avatar;not null" json:"avatar"`
	Intro      string    `gorm:"column:intro;not null" json:"intro"`
	WebSite    string    `gorm:"column:web_site;not null" json:"web_site"`
	CreateTime time.Time `gorm:"column:create_time;not null" json:"create_time"`
	UpdateTime time.Time `gorm:"column:update_time;not null" json:"update_time"`
	IsDisable  int8      `gorm:"column:is_disable;not null" json:"is_disable"`
}

// TableName sets the insert table name for this struct type
func (model *UserInfo) TableName() string {
	return "tb_user_info"
}

func AddUserInfo(m *UserInfo) error {
	return dao.Db.Save(m).Error
}

func DeleteUserInfoByID(id int) (bool, error) {
	if err := dao.Db.Delete(&UserInfo{}, id).Error; err != nil {
		return false, err
	}
	return dao.Db.RowsAffected > 0, nil
}

func DeleteUserInfo(condition string, args ...interface{}) (int64, error) {
	if err := dao.Db.Where(condition, args...).Delete(&UserInfo{}).Error; err != nil {
		return 0, err
	}
	return dao.Db.RowsAffected, nil
}

func UpdateUserInfo(m *UserInfo) error {
	return dao.Db.Save(m).Error
}

func GetUserInfoByID(id int) (*UserInfo, error) {
	var m UserInfo
	if err := dao.Db.First(&m, id).Error; err != nil {
		return nil, err
	}
	return &m, nil
}

func GetUserInfos(condition string, args ...interface{}) ([]*UserInfo, error) {
	res := make([]*UserInfo, 0)
	if err := dao.Db.Where(condition, args...).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}
