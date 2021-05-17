package models

import "time"

type UserInfo struct {
	Id         int       `gorm:"primaryKey" json:"id"`
	Email      string    `json:"email"`
	NickName   string    `gorm:"column:nickname" json:"nickname"`
	Avatar     string    `json:"avatar"`
	Intro      string    `json:"intro"`
	WebSite    string    `gorm:"column:web_site" json:"web_site"`
	CreateTime time.Time `json:"create_time"`
	UpdateTime time.Time `json:"update_time"`
	IsDisable  int8      `json:"is_disable"`
}

func (UserInfo) TableName() string {
	return "tb_user_info"
}
