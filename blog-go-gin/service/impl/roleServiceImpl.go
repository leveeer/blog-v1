package impl

import (
	pb "blog-go-gin/go_proto"
	"blog-go-gin/models/model"
	"sync"
)

type RoleServiceImpl struct {
	wg sync.WaitGroup
}

func (r *RoleServiceImpl) GetAdminUsersRole() ([]*pb.ScUserRole, error) {
	roles, err := model.GetRoles("1 = 1")
	if err != nil {
		return nil, err
	}
	var roleSlice []*pb.ScUserRole
	for _, role := range roles {
		roleSlice = append(roleSlice, &pb.ScUserRole{
			Id:       int64(role.ID),
			RoleName: role.RoleName,
		})
	}
	return roleSlice, nil
}

func NewRoleServiceImpl() *RoleServiceImpl {
	return &RoleServiceImpl{}
}
