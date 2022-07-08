package domain

type Service interface {
	InsertCategory(domain Category) (categoryObj Category, err error)
	GetAllCategory() (categoryObj []Category, err error)
	GetCategoryByID(id int) (categoryObj Category, err error)
	UpdateCategory(id int, domain Category) (categoryObj Category, err error)
	DeleteCategory(id int) (err error)
	InsertOnline(domain Online) (onlineClassObj Online, err error)
	GetAllOnline() (onlineClassObj []Online, err error)
	GetOnlineByID(id int) (onlineClassObj Online, err error)
	UpdateOnline(id int, domain Online) (onlineClassObj Online, err error)
	DeleteOnline(id int) (err error)
	InsertOffline(domain Offline) (offlineClassObj Offline, err error)
	GetAllOffline() (offlineClassObj []Offline, err error)
	GetOfflineByID(id int) (offlineClassObj Offline, err error)
	UpdateOffline(id int, domain Offline) (offlineClassObj Offline, err error)
	DeleteOffline(id int) (err error)
}

type Repository interface {
	CreateCategory(domain Category) (categoryObj Category, err error)
	GetCategory() (categoryObj []Category, err error)
	GetCategoryByID(id int) (categoryObj Category, err error)
	UpdateCategory(id int, domain Category) (categoryObj Category, err error)
	DeleteCategory(id int) (err error)
	CreateOnline(domain Online) (onlineClassObj Online, err error)
	GetOnline() (onlineClassObj []Online, err error)
	GetOnlineByID(id int) (onlineClassObj Online, err error)
	UpdateOnline(id int, domain Online) (onlineClassObj Online, err error)
	DeleteOnline(id int) (err error)
	CreateOffline(domain Offline) (offlineClassObj Offline, err error)
	GetOffline() (offlineClassObj []Offline, err error)
	GetOfflineByID(id int) (offlineClassObj Offline, err error)
	UpdateOffline(id int, domain Offline) (offlineClassObj Offline, err error)
	DeleteOffline(id int) (err error)
}
