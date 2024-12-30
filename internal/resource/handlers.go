package resource

type ResourceHandler struct {
	service ResourceServiceIface
}

func NewResourceHandler(service ResourceServiceIface) ResourceHandlerIface {
	return &ResourceHandler{service: service}
}

type ResourceHandlerIface interface{}

var _ ResourceHandlerIface = &ResourceHandler{}
