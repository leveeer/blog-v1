package impl

import (
	pb "blog-go-gin/go_proto"
	"blog-go-gin/models/model"
	"sync"
)

type UserRoleServiceImpl struct {
	wg sync.WaitGroup
}

func NewUserRoleServiceImpl() *UserRoleServiceImpl {
	return &UserRoleServiceImpl{}
}

func (u *UserRoleServiceImpl) GetUserRoleAndUsername(userId int) (*pb.UserRole, error) {
	role, err := model.GetUserRoleAndUserName(userId)
	if err != nil {
		return nil, err
	}
	return &pb.UserRole{
		Id:       int32(role.RoleID),
		RoleId:   int32(role.ID),
		UserId:   int32(role.UserID),
		Username: role.Username,
	}, err
}
