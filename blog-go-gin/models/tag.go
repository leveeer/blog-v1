package models

import (
	_ "gorm.io/gorm"
	"time"
)

type Tag struct {
	Id          int       `gorm:"primaryKey" json:"id"`
	TagName     string    `json:"tag_name"`
	Status      int8      `gorm:"default:1" json:"status"`
	ClickCount  int       `json:"click_count"`
	CreatedTime time.Time ` json:"create_time"`
	UpdatedTime time.Time `json:"update_time"`
}

func (Tag) TableName() string {
	return "tb_tag"
}
