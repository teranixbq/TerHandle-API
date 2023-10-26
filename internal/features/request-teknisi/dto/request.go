package dto

import "terhandle/internal/features/request-teknisi/entity"

type RequestCreate struct {
	UsersID   uint          `json:"user_id" form:"user_id"`
	TeknisiID uint          `json:"teknisi_id" form:"teknisi_id"`
	
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

// type RequestCreateDetail struct {
// 	Profesi      string
// 	NIK          int
// 	Alamat       string
// 	Longitude    string
// 	Latitude     string
// 	No_telp      int
// 	Email        string
// 	Achievements []entity.AchievementCore
// }

// func RequestCreateDetailToCore(dataRequest RequestCreateDetail) entity.Core {
// 	return entity.Core{
// 		Profesi:      dataRequest.Profesi,
// 		NIK:          dataRequest.NIK,
// 		Alamat:       dataRequest.Alamat,
// 		Longitude:    dataRequest.Longitude,
// 		Latitude:     dataRequest.Latitude,
// 		No_telp:      dataRequest.No_telp,
// 		Email:        dataRequest.Email,
// 		Achievements: dataRequest.Achievements,
// 	}
// }
