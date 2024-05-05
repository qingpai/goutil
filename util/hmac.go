package util

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
)

func HmacSha256(data string, secret string) string {
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}

func CheckMAC(data string, dataMac string, secret string) bool {
	mac := hmac.New(sha256.New, []byte(secret))

	originMac, err := hex.DecodeString(dataMac)
	if err != nil {
		return false
	}

	mac.Write([]byte(data))
	expectedMAC := mac.Sum(nil)
	return hmac.Equal(expectedMAC, originMac)
}
