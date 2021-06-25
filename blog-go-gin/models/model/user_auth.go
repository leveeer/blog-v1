package model

import (
	"blog-go-gin/dao"
	"gorm.io/gorm"
)

type UserAuth struct {
	CreateTime    int64  `gorm:"column:create_time;not null"`
	ID            int    `gorm:"column:id;primaryKey;unique;not null;autoIncrement"`
	IPAddr        string `gorm:"column:ip_addr;not null"`
	IPSource      string `gorm:"column:ip_source;not null"`
	LastLoginTime int64  `gorm:"column:last_login_time;not null"`
	LoginType     int8   `gorm:"column:login_type;not null"`
	Password      string `gorm:"column:password;not null"`
	UserInfoID    int    `gorm:"column:user_info_id;not null"`
	Username      string `gorm:"column:username;unique;not null"`
	NickName      string `gorm:"->"`
	Avatar        string `gorm:"->"`
	Intro         string `gorm:"->"`
	WebSite       string `gorm:"->"`
	IsDisable     bool   `gorm:"->"`
}

// TableName sets the insert table name for this struct type
func (model *UserAuth) TableName() string {
	return "tb_user_auth"
}

func AddUserAuth(tx *gorm.DB, m *UserAuth) error {
	return tx.Debug().Save(m).Error
}

func DeleteUserAuthByID(id int) (bool, error) {
	if err := dao.Db.Debug().Delete(&UserAuth{}, id).Error; err != nil {
		return false, err
	}
	return dao.Db.Debug().RowsAffected > 0, nil
}

func DeleteUserAuth(condition string, args ...interface{}) (int64, error) {
	if err := dao.Db.Debug().Where(condition, args...).Delete(&UserAuth{}).Error; err != nil {
		return 0, err
	}
	return dao.Db.Debug().RowsAffected, nil
}

func UpdateUserAuth(m *UserAuth) error {
	return dao.Db.Debug().Save(m).Error
}

func GetUserAuthByID(id int) (*UserAuth, error) {
	var m UserAuth
	if err := dao.Db.Debug().First(&m, id).Error; err != nil {
		return nil, err
	}
	return &m, nil
}

func GetUserAuthByUsername(username string) (*UserAuth, error) {
	var m UserAuth
	if err := dao.Db.Debug().Where("username = ?", username).First(&m).Error; err != nil {
		return nil, err
	}
	return &m, nil
}

func GetUserAuths(condition string, args ...interface{}) ([]*UserAuth, error) {
	res := make([]*UserAuth, 0)
	if err := dao.Db.Debug().Where(condition, args...).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

func GetLoginResponse(username string) (*UserAuth, error) {
	var m UserAuth
	if err := dao.Db.Debug().Select("tb_user_auth.*,tui.nickname,tui.avatar,tui.intro,tui.web_site,tui.is_disable").
		Where("username = ?", username).
		Joins("JOIN tb_user_info as tui on tui.id = tb_user_auth.user_info_id").
		Find(&m).Error; err != nil {
		return nil, err
	}
	return &m, nil

}
