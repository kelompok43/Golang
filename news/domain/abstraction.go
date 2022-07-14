package domain

type Service interface {
	InsertData(domain News) (newsObj News, err error)
	GetAllData() (newsObj []News, err error)
	GetByID(id int) (newsObj News, err error)
	UpdateData(id int, domain News) (newsObj News, err error)
	DeleteData(id int) (err error)
	InsertCategory(domain Category) (newsCategoryObj Category, err error)
	GetAllCategory() (newsCategoryObj []Category, err error)
	GetCategoryByID(id int) (newsCategoryObj Category, err error)
	UpdateCategory(id int, domain Category) (newsCategoryObj Category, err error)
	DeleteCategory(id int) (err error)
}

type Repository interface {
	Create(domain News) (newsObj News, err error)
	Update(id int, domain News) (newsObj News, err error)
	Get() (newsObj []News, err error)
	GetByID(id int) (newsObj News, err error)
	Delete(id int) (err error)
	CreateCategory(domain Category) (newsCategoryObj Category, err error)
	UpdateCategory(id int, domain Category) (newsCategoryObj Category, err error)
	GetCategory() (newsCategoryObj []Category, err error)
	GetCategoryByID(id int) (newsCategoryObj Category, err error)
	DeleteCategory(id int) (err error)
}
