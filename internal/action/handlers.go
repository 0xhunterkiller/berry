package action

type ActionHandler struct {
	service ActionServiceIface
}

func NewActionHandler(service ActionServiceIface) ActionHandlerIface {
	return &ActionHandler{service: service}
}

type ActionHandlerIface interface{}

var _ ActionHandlerIface = &ActionHandler{}
