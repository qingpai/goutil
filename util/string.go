package util

import (
	"encoding/json"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func init() {
	rand.NewSource(time.Now().UnixNano())
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ~!@#$%^&*_+")

func SplitToIntArr(input string) ([]int64, error) {
	strs := Split(input, ',')
	ary := make([]int64, len(strs))
	var err error
	for i := range ary {
		ary[i], err = strconv.ParseInt(strs[i], 10, 64)
	}

	if err != nil {
		return nil, err
	}

	return ary, nil
}

func RandomString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func ArrayStartsWith(haystack, needle []string) bool {
	for i := 0; i < len(needle); i++ {
		if haystack[i] != needle[i] {
			return false
		}
	}

	return true
}

func ArrayEndsWith(haystack, needle []string) bool {
	haystackLen := len(haystack)
	needleLen := len(needle)
	for i := 0; i < len(needle); i++ {
		if haystack[haystackLen-needleLen+i] != needle[i] {
			return false
		}
	}

	return true
}

func Atoi(input string) int {
	if result, err := strconv.Atoi(input); err == nil {
		return result
	}

	return 0
}

func Atoi64(input string) int64 {
	if result, err := strconv.Atoi(input); err == nil {
		return int64(result)
	}

	return 0
}

func Atof(input string) float64 {
	if result, err := strconv.ParseFloat(input, 32); err == nil {
		return result
	}

	return 0
}

func ToString(v interface{}) string {
	if result, err := json.Marshal(v); err == nil {
		return string(result)
	}

	return ""
}

var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

func ToSnakeCase(str string) string {
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}

func Isalphanumeric(input string) bool {
	regex := `^[a-zA-Z0-9+-_/\.\\s+]{1,10}$`
	reg := regexp.MustCompile(regex)

	return reg.MatchString(input)
}
