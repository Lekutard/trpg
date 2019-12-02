package api

import (
	"net/http"

	"github.com/trpg/usecases"
)

//ErrorResponseObject error時に返却するオブジェクト
type ErrorResponseObject struct {
	Message string `json:"message"`
}

//GetErrorResponse ErrorCodeとErrorResponseObjectを返却する
func GetErrorResponse(err *usecases.UError) (int, ErrorResponseObject) {
	return getErrorHTTPStatus(err.Code), ErrorResponseObject{
		Message: err.Msg,
	}
}

func getErrorHTTPStatus(errCode int) int {
	switch errCode {
	case usecases.BadRequest:
		return http.StatusBadRequest
	case usecases.Unauthorized:
		return http.StatusUnauthorized
	case usecases.NotFound:
		return http.StatusNotFound
	case usecases.Conflict:
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}
