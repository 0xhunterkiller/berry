package permissions

type PermissionHandler struct {
	service PermissionServiceIface
}

func NewPermissionHandler(service PermissionServiceIface) PermissionHandlerIface {
	return &PermissionHandler{service: service}
}

type PermissionHandlerIface interface{}

var _ PermissionHandlerIface = &PermissionHandler{}
