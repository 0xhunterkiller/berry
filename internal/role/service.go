package role

type roleService struct {
	store RoleStoreIface
}

func NewRoleService(store RoleStoreIface) RoleServiceIface {
	return &roleService{store: store}
}

type RoleServiceIface interface{}

var _ RoleServiceIface = &roleService{}
