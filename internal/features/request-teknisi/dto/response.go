package dto

import ( 
	"terhandle/internal/features/request-teknisi/entity"
	// _feedback"terhandle/internal/features/feedback/entity"
	_feedbackdto"terhandle/internal/features/feedback/dto"
)

type ResponHistoryUser struct {
	Id        uint
	TeknisiID uint
	Foto      []ResponFotoUser
	Deskripsi string
	Jarak     float64
	Biaya     float64
	Status    string
	Feedback  []_feedbackdto.ResponseFeedback
}

type ResponFotoUser struct {
	Id   int
	Foto string
}


func CoreToResponseHistory(dataCore entity.Core) ResponHistoryUser {
	return ResponHistoryUser{
		Id:        dataCore.Id,
		TeknisiID: dataCore.TeknisiID,
		Foto:      ListFotoCoreToResponFotoUser(dataCore.Foto),
		Deskripsi: dataCore.Deskripsi,
		Jarak:     dataCore.Jarak,
		Biaya:     dataCore.Biaya,
		Status:    dataCore.Status,
		Feedback: _feedbackdto.CoreToResponseFeedbackList(dataCore.Feedback),
	}
}

func ListFotoCoreToResponFotoUser(dataCore []entity.FotoCore) []ResponFotoUser {
	var result []ResponFotoUser
	for _, value := range dataCore {
		var UserRequestFoto = ResponFotoUser{
			Id:   value.Id,
			Foto: value.Foto,
		}
		result = append(result, UserRequestFoto)
	}
	return result
}

func CoreToResponseHistoryUserList(dataCore []entity.Core) []ResponHistoryUser {
	var result []ResponHistoryUser
	for _, v := range dataCore {
		result = append(result, CoreToResponseHistory(v))
	}
	return result
}

type ResponHistoryTeknisi struct {
	Id        uint
	UsersID   uint
	Foto      []ResponFotoTeknisi
	Deskripsi string
	Jarak     float64
	Biaya     float64
	Status    string
	Feedback  []_feedbackdto.ResponseFeedback
}

type ResponFotoTeknisi struct {
	Id   int
	Foto string
}

func CoreToResponseHistoryTeknisi(dataCore entity.Core) ResponHistoryTeknisi {
	return ResponHistoryTeknisi{
		Id:        dataCore.Id,
		UsersID:   dataCore.UsersID,
		Foto:      ListFotoCoreToResponFotoTeknisi(dataCore.Foto),
		Deskripsi: dataCore.Deskripsi,
		Jarak:     dataCore.Jarak,
		Biaya:     dataCore.Biaya,
		Status:    dataCore.Status,
		Feedback: _feedbackdto.CoreToResponseFeedbackList(dataCore.Feedback),
	}
}

func ListFotoCoreToResponFotoTeknisi(dataCore []entity.FotoCore) []ResponFotoTeknisi {
	var result []ResponFotoTeknisi
	for _, value := range dataCore {
		var TeknisiRequestFoto = ResponFotoTeknisi{
			Id:   value.Id,
			Foto: value.Foto,
		}
		result = append(result, TeknisiRequestFoto)
	}
	return result
}

func CoreToResponseHistoryTeknisiList(dataCore []entity.Core) []ResponHistoryTeknisi {
	var result []ResponHistoryTeknisi
	for _, v := range dataCore {
		result = append(result, CoreToResponseHistoryTeknisi(v))
	}
	return result
}

func ResponsesHistoryList(data []entity.Core, role string) interface{} {
	if role == "user" {
		return CoreToResponseHistoryUserList(data)
	}
	if role == "teknisi" {
		return CoreToResponseHistoryTeknisiList(data)
	}
	return nil
}
