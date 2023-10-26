package dto

import "terhandle/internal/features/request-teknisi/entity"

type RequestCreate struct {
	UsersID   uint `json:"user_id" form:"user_id"`
	TeknisiID uint `json:"teknisi_id" form:"teknisi_id"`

	Foto      []RequestFoto `json:"foto" form:"foto"`
	Deskripsi string        `json:"deskripsi" form:"deskripsi"`
	Jarak     int           `json:"jarak" form:"jarak"`
}

type RequestFoto struct {
	RequestTeknisiID uint
	Foto             string
}

func RequestCreateToCore(dataRequest RequestCreate) entity.Core {
	return entity.Core{
		UsersID:   dataRequest.UsersID,
		TeknisiID: dataRequest.TeknisiID,
		Foto:      ListRequestFotoToFotoCore(dataRequest.Foto),
		Deskripsi: dataRequest.Deskripsi,
		Jarak:     dataRequest.Jarak,
	}
}

func ListRequestFotoToFotoCore(dataCore []RequestFoto) []entity.FotoCore {
	var result []entity.FotoCore
	for _, value := range dataCore {
		var UserFoto = entity.FotoCore{
			RequestTeknisiID: value.RequestTeknisiID,
			Foto:             value.Foto,
		}
		result = append(result, UserFoto)
	}
	return result
}

type RequestClaims struct {
	Biaya               float64
	Diproses            bool

}

func RequestClaimsToCore(dataRequest RequestClaims) entity.Core {
	return entity.Core{

		Biaya:               dataRequest.Biaya,
		Diproses:            dataRequest.Diproses,

	}
}

type RequestUpdateStatus struct {
	Diproses            bool
	Konfirmasi_biaya    bool
	Menunggu_konfirmasi bool
	Dibatalkan          bool
	Selesai             bool
}

func RequestUpdateStatusToCore(dataRequest RequestUpdateStatus) entity.Core {
	return entity.Core{
		Diproses:            dataRequest.Diproses,
		Konfirmasi_biaya:    dataRequest.Konfirmasi_biaya,
		Menunggu_konfirmasi: dataRequest.Dibatalkan,
		Dibatalkan:          dataRequest.Dibatalkan,
		Selesai:             dataRequest.Selesai,
	}
}
