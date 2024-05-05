package util

import "github.com/golang-jwt/jwt/v5"

type ApiClaims struct {
	ClientId int64 `json:"clientId"`
	UserId   int64 `json:"userId"`
	LoginId  int64 `json:"loginId"`
	jwt.RegisteredClaims
}

// 眼部数据
type EyeData struct {
	Od string `json:"od"`
	Os string `json:"os"`
}
