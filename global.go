package lita

import (
	"strings"
)

const (
	Name = "lita"
	Port = "9933"
)

func FirstUpper(str string) string {
	if str == "" {
		return ""
	}
	return strings.ToUpper(str[:1]) + str[1:]
}
