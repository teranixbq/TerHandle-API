package entity

type UserRepositoryInterface interface {
	Login(email, password string) (Core, string, error)
	Insert(data Core) error
	Update(userid int, data Core) error
	UpdateField(userid int, field string, value string) error
	SelectAll() ([]Core, error)
	// SelectById(id int, userid int) (Core, error)
	
	// Delete(id int, userid int) (row int, err error)
}

type UserServiceInterface interface {
	Create(data Core) error
	Login(email, password string) (Core, string, error)
	CreateUpdateDetail(userid int, data Core) error
	RequestTeknisiRole(userid int) error
	GetAll() ([]Core, error)
	// GetAll(userid int) ([]Core, error)
	// GetById(id int, userid int) (Core, error)
	// Update(id int, userid int, input Core) (data Core, err error)
	// Delete(id int, userid int) (err error)
}
