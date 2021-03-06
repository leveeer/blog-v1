package model

import (
	"blog-go-gin/dao"
	"gorm.io/gorm"
)

type UserRole struct {
	ID       int    `gorm:"column:id;primaryKey;unique;not null;autoIncrement"`
	RoleID   int    `gorm:"column:role_id;not null"`
	UserID   int    `gorm:"column:user_id;not null"`
	Username string `gorm:"->"`
}

// TableName sets the insert table name for this struct type
func (model *UserRole) TableName() string {
	return "tb_user_role"
}

func AddUserRole(tx *gorm.DB, m *UserRole) error {
	return tx.Debug().Save(m).Error
}

func DeleteUserRoleByID(id int) (bool, error) {
	if err := dao.Db.Debug().Delete(&UserRole{}, id).Error; err != nil {
		return false, err
	}
	return dao.Db.Debug().RowsAffected > 0, nil
}

func DeleteUserRole(tx *gorm.DB, condition string, args ...interface{}) (int64, error) {
	if err := tx.Debug().Where(condition, args...).Delete(&UserRole{}).Error; err != nil {
		return 0, err
	}
	return tx.Debug().RowsAffected, nil
}

func UpdateUserRole(m *UserRole) error {
	return dao.Db.Debug().Save(m).Error
}

func GetUserRoleByID(id int) (*UserRole, error) {
	var m UserRole
	if err := dao.Db.Debug().First(&m, id).Error; err != nil {
		return nil, err
	}
	return &m, nil
}

func GetUserRoles(condition string, args ...interface{}) ([]*UserRole, error) {
	res := make([]*UserRole, 0)
	if err := dao.Db.Debug().Where(condition, args...).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

func GetUserRoleAndUserName(userID int) (*UserRole, error) {
	var m UserRole
	if err := dao.Db.Debug().Select("tb_user_role.id,user_id,role_id,username").Where("user_id = ?", userID).
		Joins("JOIN tb_user_auth as tua on tua.user_info_id = tb_user_role.user_id").
		Find(&m).Error; err != nil {
		return nil, err
	}
	return &m, nil
}
