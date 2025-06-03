package utils

import "strings"

const MaskedValue = "******"

func MaskSensitive(data map[string]interface{}) map[string]interface{} {
	for key, value := range data {
		loweredKey := strings.ToLower(key)

		if isSensitive(loweredKey) {
			data[key] = MaskedValue
			continue
		}

		switch v := value.(type) {
		case map[string]interface{}:
			data[key] = MaskSensitive(v)
		case []interface{}:
			for i, item := range v {
				if itemMap, ok := item.(map[string]interface{}); ok {
					v[i] = MaskSensitive(itemMap)
				}
			}
			data[key] = v
		}
	}
	return data
}

var sensitiveKeys = []string{
	"password", "pass", "pwd", "secret",
	"token", "access_token", "refresh_token",
	"username", // optional, depending on your use case
}

func isSensitive(key string) bool {
	for _, s := range sensitiveKeys {
		if key == s {
			return true
		}
	}
	return false
}
