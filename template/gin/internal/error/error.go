package error

import "net/http"

const (
	DefaultErrorCaseCode = "00"
)

const (
	BadRequest          = "Bad Request"
	InternalServerError = "internal server error"
	AccountNumberFailed = "Account Number Failed"
	InvalidFormat       = "Invalid Format"
	InvalidMandatory    = "Invalid Mandatory"
	RequestIdDuplicated = "RequestId Duplicated"
	BankNotValid        = "Bank Not Valid"
	ExternalSvcError    = "Error External Service"
	ErrUnauthorized     = "Unauthorized"
)

var (
	ErrorMapCaseCode = map[string]string{
		BadRequest:          "00",
		InternalServerError: "00",
		AccountNumberFailed: "00",
		InvalidFormat:       "01",
		InvalidMandatory:    "02",
		RequestIdDuplicated: "01",
		BankNotValid:        "01",
		ExternalSvcError:    "01",
		ErrUnauthorized:     "01",
	}

	ErrorMapMessage = map[string]string{
		BadRequest:          "Bad Request",
		InternalServerError: "internal server error",
		InvalidFormat:       "Format %v tidak sesuai",
		InvalidMandatory:    "Field %v tidak boleh kosong",
		AccountNumberFailed: "Terjadi Kesalahan. Periksa Kembali Nomor atau Nama Akun Tujuan",
		RequestIdDuplicated: "Request Id Duplicated",
		BankNotValid:        "Bank Not Valid",
		ExternalSvcError:    "%s",
		ErrUnauthorized:     "Unauthorized",
	}

	ErrorMapHttpCode = map[string]int{
		BadRequest:          http.StatusBadRequest,
		InternalServerError: http.StatusInternalServerError,
		InvalidFormat:       http.StatusBadRequest,
		InvalidMandatory:    http.StatusBadRequest,
		AccountNumberFailed: http.StatusForbidden,
		RequestIdDuplicated: http.StatusForbidden,
		BankNotValid:        http.StatusNotFound,
		ExternalSvcError:    http.StatusServiceUnavailable,
		ErrUnauthorized:     http.StatusUnauthorized,
	}
)
