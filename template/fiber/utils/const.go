package utils

import (
	"crypto/rand"
	"math/big"
)

const (
	DEV                    string = "dev"
	PROD                   string = "prod"
	DEV_TEST               string = "dev_test"
	EnvFile                       = "env/env.yml"
	EnvDevFile                    = "env/env_dev.yml"
	EnvProdFile                   = "env/env_prod.yml"
	UserId                        = "userid"
	ServiceCode                   = "01"
	ClientKeySignature            = "X-CLIENT-KEY"
	SignatureHeader               = "X-SIGNATURE"
	TimestampHeader               = "X-TIMESTAMP"
	HeaderRequestID               = "X-REQUEST-ID"
	ContextRequestIDKey           = "X-Context-Request-ID"
	XTimestampLayoutFormat        = "2006-01-02T15:04:05+07:00"
	XPrivateKey                   = "X-PRIVATE_KEY"
	TraceHeaderKey                = "trace_header"
	EndpointTarget                = "endpoint_target"
)

const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func GenerateRandomString(length int) (string, error) {
	result := make([]byte, length)
	for i := range result {
		n, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			return "", err
		}
		result[i] = charset[n.Int64()]
	}
	return string(result), nil
}

const (
	MessageWhatsapp = `Your login OTP for is: *%s*. This OTP is valid for *%v minutes*. Please do not share it with anyone. If you did not request this OTP, please contact our support team immediately.`
)
