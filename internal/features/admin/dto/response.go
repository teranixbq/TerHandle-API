package dto

import "terhandle/internal/features/admin/entity"

type ResponseProfesi struct {
	Id   uint   `json:"id"`
	NamaProfesi string `json:"nama_profesi" form:"nama_profesi"`
}

func CoreProfesiToResponseProfesi(dataCore entity.CoreProfesi) ResponseProfesi {
	return ResponseProfesi{
		Id:   dataCore.Id,
		NamaProfesi: dataCore.NamaProfesi,
	}
}

func CoreToResponseProfesiList(dataCore []entity.CoreProfesi) []ResponseProfesi {
	var result []ResponseProfesi
	for _, v := range dataCore {
		result = append(result, CoreProfesiToResponseProfesi(v))
	}
	return result
}
