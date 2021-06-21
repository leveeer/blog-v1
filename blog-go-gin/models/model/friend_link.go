package model

import "blog-go-gin/dao"

type FriendLink struct {
	CreateTime  int64  `gorm:"column:create_time"`
	ID          int    `gorm:"column:id;primaryKey;unique;not null;autoIncrement"`
	LinkAddress string `gorm:"column:link_address;not null"`
	LinkAvatar  string `gorm:"column:link_avatar;not null"`
	LinkIntro   string `gorm:"column:link_intro;not null"`
	LinkName    string `gorm:"column:link_name;not null"`
}

// TableName sets the insert table name for this struct type
func (model *FriendLink) TableName() string {
	return "tb_friend_link"
}

func AddFriendLink(m *FriendLink) error {
	return dao.Db.Debug().Save(m).Error
}

func DeleteFriendLinkByID(id int) (bool, error) {
	if err := dao.Db.Debug().Delete(&FriendLink{}, id).Error; err != nil {
		return false, err
	}
	return dao.Db.Debug().RowsAffected > 0, nil
}

func DeleteFriendLink(condition string, args ...interface{}) (int64, error) {
	if err := dao.Db.Debug().Where(condition, args...).Delete(&FriendLink{}).Error; err != nil {
		return 0, err
	}
	return dao.Db.Debug().RowsAffected, nil
}

func UpdateFriendLink(m *FriendLink) error {
	return dao.Db.Debug().Save(m).Error
}

func GetFriendLinkByID(id int) (*FriendLink, error) {
	var m FriendLink
	if err := dao.Db.Debug().First(&m, id).Error; err != nil {
		return nil, err
	}
	return &m, nil
}

func GetFriendLinks(condition string, args ...interface{}) ([]*FriendLink, error) {
	res := make([]*FriendLink, 0)
	if err := dao.Db.Debug().Where(condition, args...).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}
