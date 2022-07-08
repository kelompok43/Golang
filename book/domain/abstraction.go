package domain

type Service interface {
	InsertOnlineClass(domain OnlineClass) (onlineClassObj OnlineClass, err error)
	GetAllOnlineClass() (onlineClassObj []OnlineClass, err error)
	GetOnlineClassByID(id int) (onlineClassObj OnlineClass, err error)
	UpdateOnlineClass(id int, domain OnlineClass) (onlineClassObj OnlineClass, err error)
	DeleteOnlineClass(id int) (err error)
	InsertOfflineClass(domain OfflineClass) (offlineClassObj OfflineClass, err error)
	GetAllOfflineClass() (offlineClassObj []OfflineClass, err error)
	GetOfflineClassByID(id int) (offlineClassObj OfflineClass, err error)
	UpdateOfflineClass(id int, domain OfflineClass) (offlineClassObj OfflineClass, err error)
	DeleteOfflineClass(id int) (err error)
}

type Repository interface {
	CreateOnlineClass(domain OnlineClass) (onlineClassObj OnlineClass, err error)
	GetOnlineClass() (onlineClassObj []OnlineClass, err error)
	GetOnlineClassByID(id int) (onlineClassObj OnlineClass, err error)
	UpdateOnlineClass(id int, domain OnlineClass) (onlineClassObj OnlineClass, err error)
	DeleteOnlineClass(id int) (err error)
	CreateOfflineClass(domain OfflineClass) (offlineClassObj OfflineClass, err error)
	GetOfflineClass() (offlineClassObj []OfflineClass, err error)
	GetOfflineClassByID(id int) (offlineClassObj OfflineClass, err error)
	UpdateOfflineClass(id int, domain OfflineClass) (offlineClassObj OfflineClass, err error)
	DeleteOfflineClass(id int) (err error)
}
