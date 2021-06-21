package model

import (
	"blog-go-gin/dao"
)

type Comment struct {
	ArticleID      int    `gorm:"column:article_id;not null" json:"article_id"`
	CommentContent string `gorm:"column:comment_content;not null" json:"comment_content"`
	CreateTime     int64  `gorm:"column:create_time;not null" json:"create_time"`
	ID             int    `gorm:"column:id;primaryKey;unique;not null;autoIncrement" json:"id"`
	IsDelete       int8   `gorm:"column:is_delete;not null" json:"is_delete"`
	ParentID       int    `gorm:"column:parent_id;not null" json:"parent_id"`
	ReplyID        int    `gorm:"column:reply_id;not null" json:"reply_id"`
	UserID         int    `gorm:"column:user_id;not null" json:"user_id"`
	UserNickname   string `gorm:"-"`
	ReplyNickname  string `gorm:"-"`
}

// TableName sets the insert table name for this struct type
func (model *Comment) TableName() string {
	return "tb_comment"
}

func AddComment(m *Comment) error {
	return dao.Db.Debug().Save(m).Error
}

func DeleteCommentByID(id int) (bool, error) {
	if err := dao.Db.Debug().Delete(&Comment{}, id).Error; err != nil {
		return false, err
	}
	return dao.Db.Debug().RowsAffected > 0, nil
}

func DeleteComment(condition string, args ...interface{}) (int64, error) {
	if err := dao.Db.Debug().Where(condition, args...).Delete(&Comment{}).Error; err != nil {
		return 0, err
	}
	return dao.Db.Debug().RowsAffected, nil
}

func UpdateComment(m *Comment) error {
	return dao.Db.Debug().Save(m).Error
}

func GetCommentByID(id int) (*Comment, error) {
	var m Comment
	if err := dao.Db.Debug().First(&m, id).Error; err != nil {
		return nil, err
	}
	return &m, nil
}

func GetComments(condition string, args ...interface{}) ([]*Comment, error) {
	res := make([]*Comment, 0)
	if err := dao.Db.Debug().Where(condition, args...).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}
