package resource

type resourceService struct {
	store ResourceStoreIface
}

func NewResourceService(store ResourceStoreIface) ResourceServiceIface {
	return &resourceService{store: store}
}

type ResourceServiceIface interface{}

var _ ResourceServiceIface = &resourceService{}
