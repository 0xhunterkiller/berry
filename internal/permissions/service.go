package permissions

type permissionService struct {
	store PermissionStoreIface
}

func NewPermissionService(store PermissionStoreIface) PermissionServiceIface {
	return &permissionService{store: store}
}

type PermissionServiceIface interface{}

var _ PermissionServiceIface = &permissionService{}
