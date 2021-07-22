package model

import (
	"blog-go-gin/dao"
	"blog-go-gin/models/page"
	"gorm.io/gorm"
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
	Nickname       string `gorm:"->"`
	Avatar         string `gorm:"->"`
	WebSite        string `gorm:"->"`
	ReplyCount     int    `gorm:"->"`
	ReplyNickname  string `gorm:"->"`
	ReplyWebSite   string `gorm:"->"`
	ArticleTitle   string `gorm:"->"`
}

// TableName sets the insert table name for this struct type
func (model *Comment) TableName() string {
	return "tb_comment"
}

func AddComment(tx *gorm.DB, m *Comment) error {
	return tx.Debug().Save(m).Error
}

func DeleteCommentByID(id int) (bool, error) {
	if err := dao.Db.Debug().Delete(&Comment{}, id).Error; err != nil {
		return false, err
	}
	return dao.Db.Debug().RowsAffected > 0, nil
}

func DeleteComment(tx *gorm.DB, condition string, args ...interface{}) (int64, error) {
	if err := tx.Debug().Where(condition, args...).Delete(&Comment{}).Error; err != nil {
		return 0, err
	}
	return tx.Debug().RowsAffected, nil
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

func GetCommentsByConditionWithPage(condition string, iPage *page.IPage, args ...interface{}) ([]*Comment, error) {
	res := make([]*Comment, 0)
	if err := dao.Db.Debug().Table("tb_comment c").
		Select("c.id,u.avatar,u.nickname,r.nickname AS reply_nickname,a.article_title,c.comment_content,c.create_time,c.is_delete").
		Where(condition, args...).Joins("LEFT JOIN tb_article a ON c.article_id = a.id").
		Joins("LEFT JOIN tb_user_info u ON c.user_id = u.id").Joins("LEFT JOIN tb_user_info r ON c.reply_id = r.id").
		Scopes(page.Paginate(iPage)).Order("create_time DESC").Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

func GetCommentsCountByCondition(condition string, args ...interface{}) (int64, error) {
	var count int64
	if err := dao.Db.Debug().Table("tb_comment c").Joins("LEFT JOIN tb_user_info u ON c.user_id = u.id").Where(condition, args...).Count(&count).Error; err != nil {
		return int64(0), err
	}
	return count, nil
}

func GetCommentsAndUserInfo(iPage *page.IPage, condition string, args ...interface{}) ([]*Comment, error) {
	res := make([]*Comment, 0)
	iPage.Size = 10
	if err := dao.Db.Debug().Table("tb_comment").
		Select(" tb_user_info.nickname,tb_user_info.avatar,tb_user_info.web_site,tb_comment.user_id,tb_comment.id,tb_comment.comment_content,tb_comment.create_time").
		Where(condition, args...).Joins("JOIN tb_user_info ON tb_comment.user_id = tb_user_info.id").Order("create_time DESC").
		Scopes(page.Paginate(iPage)).
		Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

func GetReplies(commentIds []int64) ([]*Comment, error) {
	res := make([]*Comment, 0)
	subQuery := dao.Db.Debug().Table("tb_comment as c").
		Select("c.user_id,u.nickname,u.avatar,u.web_site,c.reply_id, c.id,c.parent_id,c.comment_content,c.create_time,ui.nickname as reply_nickname,ui.web_site as reply_web_site,row_number () over ( PARTITION BY parent_id ORDER BY c.create_time ) row_num").
		Joins("JOIN tb_user_info u ON c.user_id = u.id").
		Joins("JOIN tb_user_info ui ON c.reply_id = ui.id").
		Where("c.is_delete = 0 AND parent_id IN (?)", commentIds)
	if err := dao.Db.Debug().Table("(?) as t", subQuery).Where("4 > row_num").
		Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

func GetReplyCountByCommentId(commentIds []int64) ([]*Comment, error) {
	res := make([]*Comment, 0)
	if err := dao.Db.Debug().Table("tb_comment").Select("parent_id,count(1) AS reply_count").
		Where("is_delete = 0 AND parent_id IN (?)", commentIds).Group("parent_id").
		Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

func GetRepliesByCommentId(iPage *page.IPage, commentIds []int64) ([]*Comment, error) {
	res := make([]*Comment, 0)
	iPage.Size = 5
	if err := dao.Db.Debug().Table("tb_comment as c").
		Select("c.user_id,u.nickname,u.avatar,u.web_site,c.reply_id, c.id,c.parent_id,c.comment_content,c.create_time,ui.nickname as reply_nickname,ui.web_site as reply_web_site").
		Joins("JOIN tb_user_info u ON c.user_id = u.id").
		Joins("JOIN tb_user_info ui ON c.reply_id = ui.id").
		Where("c.is_delete = 0 AND parent_id IN (?)", commentIds).
		Order("create_time ASC").
		Scopes(page.Paginate(iPage)).
		Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

func UpdateCommentStatus(tx *gorm.DB, condition string, status int8, args ...interface{}) error {
	return tx.Debug().Table("tb_comment").Where(condition, args...).Select("is_delete").Update("is_delete", status).Error
}
