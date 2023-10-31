package entity

type UserRepositoryInterface interface {
	Login(email, password string) (Core, string, error)
	Insert(data Core) error
	Update(id_user uint, data Core) error
	UpdateField(id_user uint, field string, value string) error
	SelectAll() ([]Core, error)
	SelectUserById(id_user uint, role string) (Core, error)
	SelectByIdWithFeedback(id_user uint) ([]Core, error)
	DeleteById(id_user uint) error
}

type UserServiceInterface interface {
	Create(data Core) error
	Login(email, password string) (Core, string, error)
	CreateUpdateDetail(id_user uint, data Core) error
	RequestTeknisiRole(id_user uint) error
	GetAll() ([]Core, error)
	GetUserById(id_user uint, role string) (Core, error)
	GetById(id_user uint) ([]Core, error)
	DeleteById(id_user uint) error
}
