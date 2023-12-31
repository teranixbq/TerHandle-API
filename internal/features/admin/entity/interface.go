package entity

type AdminRepositoryInterface interface {
	Insert(data CoreProfesi) error
	SelectAll() ([]CoreProfesi, error)
	Update(id_profesi uint, data CoreProfesi) error
	Delete(id_profesi uint) error
	InsertBiaya(data CoreTransport) error
	UpdateBiaya(id uint,data CoreTransport) error
}

type AdminServiceInterface interface {
	Create(data CoreProfesi) error
	GetAll() ([]CoreProfesi, error)
	Update(id_profesi uint, data CoreProfesi) error
	Delete(id_profesi uint) error
	InsertBiaya(data CoreTransport) error
	UpdateBiaya(id uint, data CoreTransport) error 
}
