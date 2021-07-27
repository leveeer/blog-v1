package impl

import (
	pb "blog-go-gin/go_proto"
	"blog-go-gin/models/model"
	"blog-go-gin/models/page"
	"sync"
)

type RoleServiceImpl struct {
	wg sync.WaitGroup
}

func (r *RoleServiceImpl) GetRoles(c *pb.CsCondition) (*pb.ScAdminRoles, error) {
	roles, err := model.GetRolesByConditionWithPage(c.GetKeywords(), &page.IPage{Current: int(c.GetCurrent()), Size: int(c.GetSize())}, "%"+c.GetKeywords()+"%")
	if err != nil {
		return nil, err
	}
	var roleSlice []*pb.Role
	for _, role := range roles {
		roleSlice = append(roleSlice, &pb.Role{
			Id:             int64(role.ID),
			RoleName:       role.RoleName,
			RoleLabel:      role.RoleLabel,
			CreateTime:     role.CreateTime,
			IsDisable:      int32(role.IsDisable),
			ResourceIdList: nil,
			MenuIdList:     nil,
		})
	}

	rolesCount, err := model.GetRolesCountByCondition(c.GetKeywords(), "%"+c.GetKeywords()+"%")
	if err != nil {
		return nil, err
	}
	return &pb.ScAdminRoles{
		RoleList: roleSlice,
		Count:    rolesCount,
	}, nil
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
