package util

import (
	"fmt"
	"strings"
)

// JoinToString 将任意类型的数组转换成字符串
func JoinToString[T any](input []T, delim string) string {
	return strings.Trim(strings.Join(strings.Split(fmt.Sprint(input), " "), delim), "[]")
}
