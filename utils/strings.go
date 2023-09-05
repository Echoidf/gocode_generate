package utils

import (
	"regexp"
)

// RemoveTypeLength 去除mysql字段后括号内的长度 如int(11)-->int
func RemoveTypeLength(mysqlType string) string {
	re := regexp.MustCompile(`\(\d+\)`)
	return re.ReplaceAllString(mysqlType, "")
}
