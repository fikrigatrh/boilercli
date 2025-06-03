package utils

import "regexp"

func ValidateReq(typeErr, field string) (err error) {
	if typeErr == "" {
		return
	}
	if typeErr != "required" {
		err = MakeError("Invalid Format", field)
		return
	}
	return MakeError("Invalid Mandatory", field)

}

func IsNumeric(s string) bool {
	re := regexp.MustCompile(`^[0-9]+$`)
	return re.MatchString(s)
}
