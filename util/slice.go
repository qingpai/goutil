package util

import (
	"fmt"
	"strings"
)

func SliceToString[T any](input []T, delim string) string {
	return strings.Trim(strings.Replace(fmt.Sprint(input), " ", delim, -1), "[]")
}

func SliceRemove[T comparable](slice []T, target T) []T {
	result := make([]T, 0)
	for _, v := range slice {
		if v != target {
			result = append(result, v)
		}
	}
	return result
}
