package enum

type RoleEnum int

const (
	Admin RoleEnum = iota + 1
	User
	Test
)

var (
	RoleCh = map[RoleEnum]string{
		Admin: "管理员",
		User:  "用户",
		Test:  "测试",
	}
	RoleZh = map[RoleEnum]string{
		Admin: "admin",
		User:  "user",
		Test:  "test",
	}
	RoleID = map[RoleEnum]int{
		Admin: 1,
		User:  2,
		Test:  3,
	}
)

func (r RoleEnum) GetRoleCh() string {
	return RoleCh[r]
}

func (r RoleEnum) GetRoleZh() string {
	return RoleZh[r]
}

func (r RoleEnum) GetRoleId() int {
	return RoleID[r]
}
