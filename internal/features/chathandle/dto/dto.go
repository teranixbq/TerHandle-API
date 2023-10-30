package dto

type RequestChat struct {
	Pertanyaan string `json:"pertanyaan" form:"pertanyaan"`
}

type ResponseRecomended struct {
	Status     string
	ChatHandle string
}
