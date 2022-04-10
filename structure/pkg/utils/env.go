package utils

import "strings"

func ToEnvKey(str string) string {
	str = strings.ToUpper(str)
	str = strings.ReplaceAll(str, ".", "_")
	return str
}
