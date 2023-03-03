package entity

type WebResponse struct {
	Status  string      `json:"status"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}
type WebResponseListAndDetail struct {
	Status  string `json:"status"`
	Message string `json:"message"`

	Data interface{} `json:"data"`
}
type WebResponseListAndDetailNotFound struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}
type InfoDetail struct {
	NextRowId int `json:"nextRowId"`
	PrevRowId int `json:"prevRowId"`
}
type InfoList struct {
	Allrec  int `json:"allrec"`
	Sentrec int `json:"sentrec"`
}
type WebResponseError struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Error  interface{} `json:"error"`
}
type ReqList struct {
	Page            int    `query:"page" myvalidator:"type:stringnumber;minLength:5;maxLength:5"`
	Perpage         int    `query:"perpage"`
	Filter          string `query:"filter"`
	Order           string `query:"order"`
	Header          string `query:"header"`
	ActivityGroupId int    `query:"activity_group_id"`
}

type ReqListById struct {
	Id int `query:"id"`
}
type ReqListPilihan struct {
	Filter    string `query:"filter"`
	Order     string `query:"order"`
	Type      string `query:"type"`
	Condition string `query:"condition"`
	Header    string `query:"header"`
}
