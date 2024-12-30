package role

type RoleHandler struct {
	service RoleServiceIface
}

func NewRoleHandler(service RoleServiceIface) RoleHandlerIface {
	return &RoleHandler{service: service}
}

type RoleHandlerIface interface{}

var _ RoleHandlerIface = &RoleHandler{}
