package impl

import (
	"blog-go-gin/dao"
	pb "blog-go-gin/go_proto"
	"blog-go-gin/models/model"
	"gorm.io/gorm"
	"sync"
)

type UserInfoServiceImpl struct {
	wg sync.WaitGroup
}

func (u *UserInfoServiceImpl) UpdateUserStatus(userStatus *pb.CsUserStatus) error {
	err := dao.SqlTransaction(dao.Db.Begin(), func(tx *gorm.DB) error {
		err := model.UpdateUserStatus(tx, "id = ?", userStatus.IsDisable, userStatus.UserId)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func NewUserInfoServiceImpl() *UserInfoServiceImpl {
	return &UserInfoServiceImpl{}
}
