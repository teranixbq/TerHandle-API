package entity

type UserRequestRepositoryInterface interface {
	Insert(data Core) error
	Update(userid int, data Core) error
	UpdateField(userid int, field string, value string) error
	SelectByIdAndRole(userid, teknisiid int, role_user, role_teknisi string) error
	SelectById(id int) ([]Core, error)

	// Delete(id int, userid int) (row int, err error)
}

type UserRequestServiceInterface interface {
	Create(data Core) error
	CreateUpdateDetail(userid int, data Core) error
	GetById(userid int) ([]Core, error)
	//GetAll() ([]Core, error)
	// GetAll(userid int) ([]Core, error)
	// GetById(id int, userid int) (Core, error)
	// Update(id int, userid int, input Core) (data Core, err error)
	// Delete(id int, userid int) (err error)
}
