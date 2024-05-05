package util

import (
	"bytes"
	"encoding/json"
	"regexp"
)

// MarshalModel 序列化模型，生成下划线命名的key json格式，以便gorm update语句
func MarshalModel[T any](input T) ([]byte, error) {
	// Regexp definitions
	var keyMatchRegex = regexp.MustCompile(`\"(\w+)\":`)
	var wordBarrierRegex = regexp.MustCompile(`(\w)([A-Z])`)
	marshalled, err := json.Marshal(input) //jettison.Marshal(input)
	converted := keyMatchRegex.ReplaceAllFunc(
		marshalled,
		func(match []byte) []byte {
			return bytes.ToLower(wordBarrierRegex.ReplaceAll(
				match,
				[]byte(`${1}_${2}`),
			))
		},
	)

	return converted, err
}
