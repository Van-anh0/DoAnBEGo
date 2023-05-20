package response

import (
	"github.com/praslar/cloud0/ginext"
)

type apiErr struct {
	code    int
	message string
}

func NewError(r *ginext.Request, code int, message string) error {
	r.GinCtx.AbortWithStatus(code)
	return &apiErr{code: code, message: message}
}

func (e *apiErr) Error() string {
	return e.message
}

func (e *apiErr) StatusCode() int {
	return e.code
}
