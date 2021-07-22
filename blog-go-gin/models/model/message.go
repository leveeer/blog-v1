package model

import (
	"blog-go-gin/dao"
	"blog-go-gin/models/page"
	"gorm.io/gorm"
)

type Message struct {
	Avatar         string `gorm:"column:avatar;not null"`
	CreateTime     int64  `gorm:"column:create_time;not null"`
	ID             int    `gorm:"column:id;primaryKey;unique;not null;autoIncrement"`
	IPAddress      string `gorm:"column:ip_address;not null"`
	IPSource       string `gorm:"column:ip_source;not null"`
	MessageContent string `gorm:"column:message_content;not null"`
	Nickname       string `gorm:"column:nickname;not null"`
	Time           int8   `gorm:"column:time;not null"`
}

// TableName sets the insert table name for this struct type
func (model *Message) TableName() string {
	return "tb_message"
}

func AddMessage(tx *gorm.DB, m *Message) error {
	return tx.Debug().Save(m).Error
}

func DeleteMessageByID(id int) (bool, error) {
	if err := dao.Db.Debug().Delete(&Message{}, id).Error; err != nil {
		return false, err
	}
	return dao.Db.Debug().RowsAffected > 0, nil
}

func DeleteMessage(tx *gorm.DB, condition string, args ...interface{}) (int64, error) {
	if err := tx.Debug().Where(condition, args...).Delete(&Message{}).Error; err != nil {
		return 0, err
	}
	return tx.Debug().RowsAffected, nil
}

func UpdateMessage(m *Message) error {
	return dao.Db.Debug().Save(m).Error
}

func GetMessageByID(id int) (*Message, error) {
	var m Message
	if err := dao.Db.Debug().First(&m, id).Error; err != nil {
		return nil, err
	}
	return &m, nil
}

func GetMessages(condition string, args ...interface{}) ([]*Message, error) {
	res := make([]*Message, 0)
	if err := dao.Db.Debug().Where(condition, args...).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

func GetMessagesCount() (count int64, err error) {
	var messageCount int64
	if err := dao.Db.Debug().Table("tb_message").Count(&messageCount).Error; err != nil {
		return 0, err
	}
	return messageCount, nil
}

func GetMessagesByConditionWithPage(condition string, iPage *page.IPage, args ...interface{}) ([]*Message, error) {
	res := make([]*Message, 0)
	db := dao.Db
	if condition != "" {
		db = db.Where("nickname LIKE ?", args...)
	}
	if err := db.Debug().Scopes(page.Paginate(iPage)).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

func GetMessagesCountByCondition(condition string, args ...interface{}) (int64, error) {
	var count int64
	db := dao.Db
	if condition != "" {
		db = db.Where("nickname LIKE ?", args...)
	}
	if err := db.Debug().Table("tb_message").Count(&count).Error; err != nil {
		return int64(0), err
	}
	return count, nil
}
