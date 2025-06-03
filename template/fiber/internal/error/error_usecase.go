package error

import (
	"boilerplate/utils"
	"errors"
	"fmt"
	"net/http"
)

type UsecaseError struct {
	err  error
	args []any
}

func ErrorCustom(err error) *UsecaseError {
	msg, arguments := utils.GenerateError(err)
	return &UsecaseError{
		err:  errors.New(msg),
		args: arguments,
	}
}

func (p *UsecaseError) GetError() error {
	return p.err
}

func (p *UsecaseError) GetHTTPCode() int {
	val, ok := ErrorMapHttpCode[p.err.Error()]
	if !ok {
		return http.StatusInternalServerError
	}

	return val
}

func (p *UsecaseError) GetMessage() string {
	val, ok := ErrorMapMessage[p.err.Error()]
	if !ok {
		return ErrorMapMessage[InternalServerError]
	}

	return fmt.Sprintf(val, p.args...)
}

func (p *UsecaseError) GetCaseCode() string {
	val, ok := ErrorMapCaseCode[p.err.Error()]
	if !ok {
		return DefaultErrorCaseCode
	}

	return val
}
