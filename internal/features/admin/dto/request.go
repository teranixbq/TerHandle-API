package dto

import "terhandle/internal/features/admin/entity"

type RequestCreate struct {
	NamaProfesi string `json:"nama_profesi" form:"nama_profesi"`
}

func RequestCreateToCore(dataRequest RequestCreate) entity.CoreProfesi {
	return entity.CoreProfesi{
		NamaProfesi: dataRequest.NamaProfesi,
	}
}

type RequestCreateBiaya struct {
	Biaya float64 `json:"biaya" form:"biaya"`
}

func RequestCreateBiayaToCore(dataRequest RequestCreateBiaya) entity.CoreTransport {
	return entity.CoreTransport{
		Biaya: dataRequest.Biaya,
	}
}
