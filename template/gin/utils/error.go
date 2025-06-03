package utils

import (
	"errors"
	"fmt"
	"strings"
)

func MakeError(msgError string, a ...any) error {
	args := make([]string, 0)
	for _, arg := range a {
		args = append(args, fmt.Sprintf("%v", arg))
	}
	args = append([]string{msgError}, args...)
	return errors.New(strings.Join(args, ","))
}

func GenerateError(err error) (string, []any) {
	var a []any
	msg := err.Error()

	args := strings.Split(err.Error(), ",")
	if len(args) > 1 {
		msg = args[0]

		for _, arg := range args[1:] {
			a = append(a, arg)
		}
	}

	return msg, a
}

func ParsingError(str ...string) (field, typeErr string) {
	for i, s := range str {
		str[i] = strings.ReplaceAll(s, "'", "")
	}
	field = str[0]
	typeErr = str[1]

	return
}
