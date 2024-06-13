package util

import (
	"crypto/sha1"
	"encoding/hex"
)

func Sha1(input string) string {
	h := sha1.New()
	h.Write([]byte(input))

	return hex.EncodeToString(h.Sum(nil))
}
