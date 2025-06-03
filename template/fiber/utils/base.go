package utils

import (
	"net/http"
	"strconv"
)

type ResponseTemp struct {
	ResponseCode    string `json:"responseCode"`
	ResponseMessage string `json:"responseMessage"`
	Data            any    `json:"data"`
}

type Metadata struct {
	PageSize  int `json:"pageSize"`
	Page      int `json:"page"`
	TotalPage int `json:"totalPage"`
	TotalData int `json:"totalData"`
}

type Error interface {
	GetError() error
	GetHTTPCode() int
	GetMessage() string
	GetCaseCode() string
}

type Base struct {
	Error           string `json:"-"`
	StatusCode      int    `json:"-"`
	ResponseCode    string `json:"responseCode"`
	ResponseMessage string `json:"responseMessage"`
}

func Failure() *Base {
	return &Base{
		Error:           http.StatusText(http.StatusBadRequest),
		StatusCode:      http.StatusBadRequest,
		ResponseCode:    http.StatusText(http.StatusBadRequest),
		ResponseMessage: http.StatusText(http.StatusBadRequest),
	}
}

func CustomError(e Error) func(b *Base) {
	return func(b *Base) {
		b.StatusCode = e.GetHTTPCode()
		httpCode := strconv.Itoa(e.GetHTTPCode())
		b.Error = e.GetError().Error()
		b.ResponseCode = httpCode + ServiceCode + e.GetCaseCode()
		b.ResponseMessage = e.GetMessage()
	}
}
