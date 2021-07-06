package impl

import (
	pb "blog-go-gin/go_proto"
	"sync"
)

type UserRoleServiceImpl struct {
	wg sync.WaitGroup
}

func (u *UserRoleServiceImpl) GetUserRoleAndUsername(userId int) (*pb.UserRole, error) {
	panic("implement me")
}

func NewUserRoleServiceImpl() *UserRoleServiceImpl {
	return &UserRoleServiceImpl{}
}
