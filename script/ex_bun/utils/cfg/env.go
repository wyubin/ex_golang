package cfg

import "strings"

func ToEnvKey(str string) string {
	str = strings.ToUpper(str)
	return str
}
