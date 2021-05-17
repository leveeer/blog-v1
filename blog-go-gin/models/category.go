package models

import "time"

type Category struct {
	Id int `gorm:"primaryKey" json:"id"`
	CategoryName string `json:"category_name"`
	CreateTime time.Time `json:"create_time"`
	UpdateTime time.Time `json:"update_time"`
}

func (Category) TableName() string {
	return "tb_category"
}