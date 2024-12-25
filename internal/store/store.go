package store

type AllStores struct {
	*UserStore
}

var Store *AllStores

func InitStore() error {
	db, err := ConnectDB()
	if err != nil {
		return err
	}
	Store.UserStore = NewUserStore(db, 25, 10, 30)
	return nil
}
