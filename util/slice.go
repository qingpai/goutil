package util

import (
	"fmt"
	"strings"
)

func SliceToString[T any](input []T, delim string) string {
	return strings.Trim(strings.Replace(fmt.Sprint(input), " ", delim, -1), "[]")
}
