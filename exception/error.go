package exception

import (
	"fmt"
	"net/http"

	"github.com/NidzamuddinMuzakki/todolist/entity"
	"github.com/NidzamuddinMuzakki/todolist/helper"

	"github.com/labstack/echo/v4"
)

func ErrorHandler(c echo.Context, err interface{}) {

	if badRequestErrorsss(c, err) {
		return
	}
	internalServerError(c, err)
}
func badRequestErrorsss(c echo.Context, err interface{}) bool {

	execption, ok := err.(BadRequestErrors)
	fmt.Println(ok)
	if ok {
		webResponse := entity.WebResponse{
			Message: "",
			Status:  "Bad Request",
			Data:    execption.Error,
		}
		helper.WriteToResponseBody(c, webResponse, 200)
		return true
	} else {
		return false
	}
}
func internalServerError(c echo.Context, err interface{}) {
	webResponse := entity.WebResponseError{}
	webResponse = entity.WebResponseError{
		Code:   http.StatusInternalServerError,
		Status: "INTERNAL SERVER ERROR",
		Error:  err,
	}
	fmt.Println(err, "cek error")

	helper.WriteToResponseBody(c, webResponse, webResponse.Code)
}
