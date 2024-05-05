package util

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"strings"
)

func EncodeToBase64(v interface{}) (string, error) {
	var buf bytes.Buffer
	encoder := base64.NewEncoder(base64.RawURLEncoding, &buf)
	err := json.NewEncoder(encoder).Encode(v)
	if err != nil {
		return "", err
	}
	encoder.Close()
	return buf.String(), nil
}

func DecodeFromBase64(v interface{}, enc string) error {
	base64Decoder := base64.NewDecoder(base64.RawURLEncoding, strings.NewReader(enc))
	return json.NewDecoder(base64Decoder).Decode(v)
}
