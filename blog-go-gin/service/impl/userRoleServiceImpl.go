package impl

import (
	"blog-go-gin/dao"
	pb "blog-go-gin/go_proto"
	"blog-go-gin/models/model"
	"gorm.io/gorm"
	"sync"
)

type UserRoleServiceImpl struct {
	wg sync.WaitGroup
}

func (u *UserRoleServiceImpl) UpdateUserRole(userRole *pb.CsUpdateUserRole) error {

	err := dao.SqlTransaction(dao.Db.Begin(), func(tx *gorm.DB) error {
		//修改用户昵称
		err := model.UpdateNicknameByCondition(tx, "id = ?", userRole.Nickname, userRole.UserInfoId)
		if err != nil {
			return err
		}
		//用户角色先删除再重新添加
		_, err = model.DeleteUserRole(tx, "user_id IN (?)", userRole.UserInfoId)
		if err != nil {
			return err
		}
		for _, roleId := range userRole.RoleIdList {
			err := model.AddUserRole(tx, &model.UserRole{
				RoleID: int(roleId),
				UserID: int(userRole.UserInfoId),
			})
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func (u *UserRoleServiceImpl) GetUserRoleAndUsername(userId int) (*pb.UserRole, error) {
	panic("implement me")
}

func NewUserRoleServiceImpl() *UserRoleServiceImpl {
	return &UserRoleServiceImpl{}
}
