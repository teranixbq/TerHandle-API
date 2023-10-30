package entity

import (
	"terhandle/internal/app/model"
)

func CoreProfesiToModelProfesi(dataCore CoreProfesi) model.Profesi {
	return model.Profesi{
		Id:          dataCore.Id,
		NamaProfesi: dataCore.NamaProfesi,
	}
}

func ProfesiModelToCoreProfesi(dataModel model.Profesi) CoreProfesi {
	return CoreProfesi{
		Id:          dataModel.Id,
		NamaProfesi: dataModel.NamaProfesi,
	}
}

func ProfesiModelToCoreProfesiList(dataModel []model.Profesi) []CoreProfesi {
	var ListProfesiModel []CoreProfesi
	for _, v := range dataModel {
		ListProfesiModel = append(ListProfesiModel, ProfesiModelToCoreProfesi(v))
	}
	return ListProfesiModel
}

func ProfesiCoreToCoreProfesiList(dataModel []CoreProfesi) []model.Profesi {
	var ListProfesiModel []model.Profesi
	for _, v := range dataModel {
		ListProfesiModel = append(ListProfesiModel, CoreProfesiToModelProfesi(v))
	}
	return ListProfesiModel
}

func CoreTransportToModelTransport(dataCore CoreTransport) model.Transport {
	return model.Transport{
		Id:    dataCore.Id,
		Biaya: dataCore.Biaya,
	}
}

func TransportModelToCoreTransport(dataModel model.Transport) CoreTransport {
	return CoreTransport{
		Id:    dataModel.Id,
		Biaya: dataModel.Biaya,
	}
}
