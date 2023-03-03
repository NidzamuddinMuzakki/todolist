package exception

type BadRequestError struct {
	Desc      string `json:"errDesc"`
	DescGlob  string `json:"errDescGlob"`
	FieldName string `json:"fieldName"`
}
type BadRequestErrors struct {
	Error []BadRequestError
}

func NewBadRequestError(error []BadRequestError) BadRequestErrors {
	var d = BadRequestErrors{Error: error}
	return d
}
