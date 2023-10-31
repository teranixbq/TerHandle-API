package entity


type UserRequestRepositoryInterface interface {
	Insert(data Core) error
	Update(userid int, data Core) error
	UpdateField(id_request int, field string, value string) error
	SelectByIdAndRole(userid, teknisiid int, role_user, role_teknisi string) error
	SelectAllById(id int) ([]Core, error)
	SelectById(userid, id int) ([]Core, error)
	UpdateClaims(id int, data Core) error
	UpdateStatusKonfirmBiaya(id_user, id_request int, data Core) error
	UpdateStatusBatal(id_user, id_request int, data Core) error
	UpdateStatusSelesai(id_user, id_request int, data Core) error
	
}

type UserRequestServiceInterface interface {
	Create(data Core) error
	CreateUpdateDetail(userid int, data Core) error
	GetAllById(userid int) ([]Core, error)
	GetById(userid, id int) ([]Core, error)
	ClaimRequest(id_request int, data Core) error
	KonfirmasiBiaya(id_user, id_request int, data Core) error
	BatalkanRequest(id_user, id_request int, data Core) error
	SelesaikanRequest(id_user, id_request int, data Core) error
}
