package model

import (
	"blog-go-gin/dao"
	"gorm.io/gorm"
)

type UserInfo struct {
	ID         int    `gorm:"column:id;primaryKey;unique;not null;autoIncrement" json:"id"`
	Email      string `gorm:"column:email;not null" json:"email"`
	Nickname   string `gorm:"column:nickname;not null" json:"nickname"`
	Avatar     string `gorm:"column:avatar;not null" json:"avatar"`
	Intro      string `gorm:"column:intro;not null" json:"intro"`
	WebSite    string `gorm:"column:web_site;not null" json:"web_site"`
	CreateTime int64  `gorm:"column:create_time;not null" json:"create_time"`
	UpdateTime int64  `gorm:"column:update_time;not null" json:"update_time"`
	IsDisable  int8   `gorm:"column:is_disable;not null" json:"is_disable"`
}

// TableName sets the insert table name for this struct type
func (model *UserInfo) TableName() string {
	return "tb_user_info"
}

func AddUserInfo(tx *gorm.DB, m *UserInfo) (int, error) {
	err := tx.Debug().Save(m).Error
	if err != nil {
		return 0, err
	}
	return m.ID, nil
}

func DeleteUserInfoByID(id int) (bool, error) {
	if err := dao.Db.Debug().Delete(&UserInfo{}, id).Error; err != nil {
		return false, err
	}
	return dao.Db.Debug().RowsAffected > 0, nil
}

func DeleteUserInfo(condition string, args ...interface{}) (int64, error) {
	if err := dao.Db.Debug().Where(condition, args...).Delete(&UserInfo{}).Error; err != nil {
		return 0, err
	}
	return dao.Db.Debug().RowsAffected, nil
}

func UpdateUserInfo(tx *gorm.DB, m *UserInfo) error {
	return tx.Debug().Save(m).Error
}

func UpdateNicknameByCondition(tx *gorm.DB, condition, value string, args ...interface{}) error {
	return tx.Debug().Table("tb_user_info").Where(condition, args...).Update("nickname", value).Error
}

func UpdateUserStatus(tx *gorm.DB, condition string, value bool, args ...interface{}) error {
	return tx.Debug().Table("tb_user_info").Where(condition, args...).Select("is_disable").Update("is_disable", value).Error
}

func GetUserInfoByID(id int) (*UserInfo, error) {
	var m UserInfo
	if err := dao.Db.Debug().First(&m, id).Error; err != nil {
		return nil, err
	}
	return &m, nil
}

func GetUserInfos(condition string, args ...interface{}) ([]*UserInfo, error) {
	res := make([]*UserInfo, 0)
	if err := dao.Db.Debug().Where(condition, args...).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

func GetUserInfoCount() (count int64, err error) {
	var userCount int64
	if err := dao.Db.Debug().Table("tb_user_info").Count(&userCount).Error; err != nil {
		return 0, err
	}
	return userCount, nil
}
