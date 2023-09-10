package rest_err

import "net/http"

type Rest_err struct {
	Message string   `json:"message"`
	Err     string   `json:"err"`
	Code    int      `json:"code"`
	Causes  []Causes `json:"causes"`
}

type Causes struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func (r *Rest_err) Error() string {
	return r.Message
}

func NewRestErr(message, err string, code int, causes []Causes) *Rest_err {
	return &Rest_err{
		Message: message,
		Err:     err,
		Code:    code,
		Causes:  causes,
	}
}

func NewBadRequestErr(message string) *Rest_err {
	return &Rest_err{
		Message: message,
		Err:     "bad_request",
		Code:    http.StatusBadRequest,
	}
}

func NewBadRequestValdiationErr(message string, causes []Causes) *Rest_err {
	return &Rest_err{
		Message: message,
		Err:     "bad_request",
		Code:    http.StatusBadRequest,
	}
}

func NewInternalServerError(message string) *Rest_err {
	return &Rest_err{
		Message: message,
		Err:     "internal_server_error",
		Code:    http.StatusInternalServerError,
	}
}

func NewNotFoundError(message string) *Rest_err {
	return &Rest_err{
		Message: message,
		Err:     "not_found",
		Code:    http.StatusNotFound,
	}
}

func NewForbiddenError(message string) *Rest_err {
	return &Rest_err{
		Message: message,
		Err:     "forbidden",
		Code:    http.StatusForbidden,
	}
}
