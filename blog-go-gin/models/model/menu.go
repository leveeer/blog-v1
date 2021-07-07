package model

import "blog-go-gin/dao"

type Menu struct {
	Component  string `gorm:"column:component;not null"`
	CreateTime int64  `gorm:"column:create_time;not null"`
	Icon       string `gorm:"column:icon;not null"`
	ID         int    `gorm:"column:id;primaryKey;unique;not null;autoIncrement"`
	IsDisable  int8   `gorm:"column:is_disable;not null"`
	IsHidden   int8   `gorm:"column:is_hidden;not null"`
	Name       string `gorm:"column:name;not null"`
	OrderNum   int8   `gorm:"column:order_num;not null"`
	ParentID   int    `gorm:"column:parent_id;not null"`
	Path       string `gorm:"column:path;not null"`
	UpdateTime int64  `gorm:"column:update_time;not null"`
}

// TableName sets the insert table name for this struct type
func (model *Menu) TableName() string {
	return "tb_menu"
}

func AddMenu(m *Menu) error {
	return dao.Db.Debug().Save(m).Error
}

func DeleteMenuByID(id int) (bool, error) {
	if err := dao.Db.Debug().Delete(&Menu{}, id).Error; err != nil {
		return false, err
	}
	return dao.Db.Debug().RowsAffected > 0, nil
}

func DeleteMenu(condition string, args ...interface{}) (int64, error) {
	if err := dao.Db.Debug().Where(condition, args...).Delete(&Menu{}).Error; err != nil {
		return 0, err
	}
	return dao.Db.Debug().RowsAffected, nil
}

func UpdateMenu(m *Menu) error {
	return dao.Db.Debug().Save(m).Error
}

func GetMenuByID(id int) (*Menu, error) {
	var m Menu
	if err := dao.Db.Debug().First(&m, id).Error; err != nil {
		return nil, err
	}
	return &m, nil
}

func GetMenus(condition string, args ...interface{}) ([]*Menu, error) {
	res := make([]*Menu, 0)
	if err := dao.Db.Debug().Where(condition, args...).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

func GetUserMenus(roleid int) ([]*Menu, error) {
	res := make([]*Menu, 0)
	if err := dao.Db.Debug().Table("tb_role_menu as rm").Distinct("m.id").Select("m.id, `name`, `path`,component,icon,is_hidden,parent_id,order_num").
		Where("rm.role_id = ?", roleid).
		Joins("JOIN `tb_menu` m ON rm.menu_id = m.id").
		Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}
