package helper

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type Response struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Errors  interface{} `json:"errors"`
	Data    interface{} `json:"data"`
}
type Token struct {
	Role     string
	Username string
	Exp      string
}
type Message struct {
	Desc     string `json:"desc"`
	DescGlob string `json:"descGlob"`
}
type ResponseExec struct {
	Status  bool    `json:"status"`
	Code    int     `json:"code"`
	Message Message `json:"message"`
}

type EmptyObj struct{}

func ReadFromRequestBody(ctx echo.Context, result interface{}) {
	err := ctx.Bind(&result)
	fmt.Println(err)
	PanicIfError(err)
}

func WriteToResponseBody(c echo.Context, response interface{}, code int) {
	c.JSON(code, response)
}

func ReadResponseOutbond(res *http.Response) map[string]interface{} {
	body, _ := io.ReadAll(res.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)
	return responseBody
}
func BuildErrorResponse(message string, err string, data interface{}) Response {
	splittedError := strings.Split(err, "\n")
	res := Response{
		Status:  false,
		Message: message,
		Errors:  splittedError,
		Data:    data,
	}
	return res
}
func BuildErrorResponseExec(message Message, code1 int) ResponseExec {
	status := true
	if code1 != 200 {
		status = false
	}
	res := ResponseExec{
		Status:  status,
		Code:    code1,
		Message: message,
	}
	return res
}
func ReadDataToken(token string) Token {
	claims := jwt.MapClaims{}
	var tokenReturn Token
	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("<YOUR VERIFICATION KEY>"), nil
	})
	if err != nil {
		fmt.Println(err)
	}

	for key, val := range claims {
		if key == "exp" {
			tokenReturn.Exp = fmt.Sprintf("%v", val)
		}
		if key == "username" {
			tokenReturn.Username = fmt.Sprintf("%v", val)
		}
		if key == "role" {
			tokenReturn.Role = fmt.Sprintf("%v", val)
		}

	}
	return tokenReturn
}

func TimePlus7(now time.Time) string {

	// Note: without explicit zone, returns time in given location.

	return time.Now().Format("2006-01-02 15:04:05")
}
func DatePlus7(now time.Time) string {

	// Note: without explicit zone, returns time in given location.

	return time.Now().Format("2006-01-02")
}
