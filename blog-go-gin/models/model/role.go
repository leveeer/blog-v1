package model

import (
	"blog-go-gin/dao"
	"blog-go-gin/models/page"
)

type Role struct {
	CreateTime int64  `gorm:"column:create_time" json:"create_time"`
	ID         int    `gorm:"column:id;primaryKey;unique;not null;autoIncrement" json:"id"`
	IsDisable  int8   `gorm:"column:is_disable;not null" json:"is_disable"`
	RoleLabel  string `gorm:"column:role_label;not null" json:"role_label"`
	RoleName   string `gorm:"column:role_name;not null" json:"role_name"`
	UpdateTime int64  `gorm:"column:update_time" json:"update_time"`
	ResourceId int    `gorm:"->"`
	MenuId     int    `gorm:"->"`
}

// TableName sets the insert table name for this struct type
func (model *Role) TableName() string {
	return "tb_role"
}

func AddRole(m *Role) error {
	return dao.Db.Debug().Debug().Save(m).Error
}

func DeleteRoleByID(id int) (bool, error) {
	if err := dao.Db.Debug().Delete(&Role{}, id).Error; err != nil {
		return false, err
	}
	return dao.Db.Debug().RowsAffected > 0, nil
}

func DeleteRole(condition string, args ...interface{}) (int64, error) {
	if err := dao.Db.Debug().Where(condition, args...).Delete(&Role{}).Error; err != nil {
		return 0, err
	}
	return dao.Db.Debug().RowsAffected, nil
}

func UpdateRole(m *Role) error {
	return dao.Db.Debug().Save(m).Error
}

func GetRoleByID(id int) (*Role, error) {
	var m Role
	if err := dao.Db.Debug().First(&m, id).Error; err != nil {
		return nil, err
	}
	return &m, nil
}

func GetRoles(condition string, args ...interface{}) ([]*Role, error) {
	res := make([]*Role, 0)
	if err := dao.Db.Debug().Where(condition, args...).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

func GetRolesByConditionWithPage(condition string, iPage *page.IPage, args ...interface{}) ([]*Role, error) {
	res := make([]*Role, 0)
	db := dao.Db
	if condition != "" {
		db = db.Where("role_name LIKE ?", args...)
	}
	subQuery := db.Debug().Table("tb_role").Scopes(page.Paginate(iPage))
	if err := dao.Db.Debug().Table("(?) as r", subQuery).Select("r.id,role_name,role_label,r.create_time,r.is_disable,rr.resource_id,rm.menu_id").
		Joins("LEFT JOIN tb_role_resource rr ON r.id = rr.role_id").Joins("LEFT JOIN tb_role_menu rm on r.id = rm.role_id").
		Order("r.id DESC").Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

func GetRolesCountByCondition(condition string, args ...interface{}) (int64, error) {
	var count int64
	db := dao.Db
	if condition != "" {
		db = db.Where("role_name LIKE ?", args...)
	}
	if err := db.Debug().Table("tb_role ").Count(&count).Error; err != nil {
		return int64(0), err
	}
	return count, nil
}
