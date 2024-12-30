package action

type actionService struct {
	store ActionStoreIface
}

func NewActionService(store ActionStoreIface) ActionServiceIface {
	return &actionService{store: store}
}

type ActionServiceIface interface{}

var _ ActionServiceIface = &actionService{}
