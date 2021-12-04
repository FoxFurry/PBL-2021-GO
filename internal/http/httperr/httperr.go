package httperr

import (
	"foxy/internal/http/httpresponse"
	"github.com/gin-gonic/gin"
)

//---------------------------------------------------------------
type baseError struct {
	Msg string `json:"message"`
}

func (b baseError) Error() string {
	return b.Msg
}

//---------------------------------------------------------------

//---------------------------------------------------------------
type errorNotFound struct {
	baseError
}

func NewErrorNotFound() error {
	return errorNotFound{
		baseError: baseError{"User not found"},
	}
}

//---------------------------------------------------------------

//---------------------------------------------------------------
type errorAlreadyExists struct {
	baseError
}

func NewErrorAlreadyExists() error {
	return errorAlreadyExists{
		baseError: baseError{"User already exists"},
	}
}

//---------------------------------------------------------------

//---------------------------------------------------------------
type errorUnauthorized struct {
	baseError
}

func NewErrorUnauthorized() error {
	return errorUnauthorized{
		baseError: baseError{"Authentication failed"},
	}
}

//---------------------------------------------------------------

//---------------------------------------------------------------
type errorInternal struct {
	baseError
}

func NewErrorInternal() error {
	return errorInternal{baseError{"Unexpected internal error"}}
}

//---------------------------------------------------------------

func Handle(c *gin.Context, err interface{}) {
	switch err.(type) {
	case errorNotFound:
		httpresponse.RespondNotFound(c, err)
	case errorAlreadyExists:
		httpresponse.RespondAlreadyExists(c, err)
	case errorUnauthorized:
		httpresponse.RespondUnauthorized(c, err)
	default:
		httpresponse.RespondInternalError(c, err)
	}
}
