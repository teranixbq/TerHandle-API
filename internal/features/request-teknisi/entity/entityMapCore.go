package entity

import (
	"terhandle/internal/app/model"
	_feedbackuser "terhandle/internal/features/user/entity"
)

func UserCoreToUserModel(dataCore Core) model.RequestTeknisi {
	return model.RequestTeknisi{
		UsersID:             dataCore.UsersID,
		TeknisiID:           dataCore.TeknisiID,
		Foto:                ListFotoCoreToModelFoto(dataCore.Foto),
		Deskripsi:           dataCore.Deskripsi,
		Jarak:               dataCore.Jarak,
		Biaya:               dataCore.Biaya,
		Diproses:            dataCore.Diproses,
		Konfirmasi_biaya:    dataCore.Konfirmasi_biaya,
		Menunggu_konfirmasi: dataCore.Menunggu_konfirmasi,
		Dibatalkan:          dataCore.Dibatalkan,
		Selesai:             dataCore.Selesai,
		Status:              dataCore.Status,
		Feedback:            _feedbackuser.ListCoreFeedbackToModelFeedback(dataCore.Feedback),
		CreatedAt:           dataCore.CreatedAt,
		UpdatedAt:           dataCore.UpdatedAt,
	}
}

func ListFotoCoreToModelFoto(dataCore []FotoCore) []model.Foto {
	var result []model.Foto
	for _, value := range dataCore {
		var UserFoto = model.Foto{
			RequestTeknisiID: value.RequestTeknisiID,
			Foto:             value.Foto,
		}
		result = append(result, UserFoto)
	}
	return result
}

func UserModelToUserCore(dataModel model.RequestTeknisi) Core {
	return Core{
		Id:                  dataModel.ID,
		UsersID:             dataModel.UsersID,
		TeknisiID:           dataModel.TeknisiID,
		Foto:                ListModelFotoToFotoCore(dataModel.Foto),
		Deskripsi:           dataModel.Deskripsi,
		Jarak:               dataModel.Jarak,
		Biaya:               dataModel.Biaya,
		Diproses:            dataModel.Diproses,
		Konfirmasi_biaya:    dataModel.Konfirmasi_biaya,
		Menunggu_konfirmasi: dataModel.Menunggu_konfirmasi,
		Dibatalkan:          dataModel.Dibatalkan,
		Selesai:             dataModel.Selesai,
		Status:              dataModel.Status,
		Feedback:            _feedbackuser.ListModelFeedbackToCoreFeedback(dataModel.Feedback),
		CreatedAt:           dataModel.CreatedAt,
		UpdatedAt:           dataModel.UpdatedAt,
	}
}

func ListModelFotoToFotoCore(dataModel []model.Foto) []FotoCore {
	var result []FotoCore
	for _, value := range dataModel {
		var UserFoto = FotoCore{
			Id:               int(value.ID),
			RequestTeknisiID: value.RequestTeknisiID,
			Foto:             value.Foto,
		}
		result = append(result, UserFoto)
	}
	return result
}

func UserModelToUserCoreList(dataModel []model.RequestTeknisi) []Core {
	var ListUserRequestModel []Core
	for _, v := range dataModel {
		ListUserRequestModel = append(ListUserRequestModel, UserModelToUserCore(v))
	}
	return ListUserRequestModel
}
