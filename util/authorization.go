package util

import (
	"github.com/gin-gonic/gin"
)

type AuthorizationType string

const (
	AuthorizationPlatform AuthorizationType = "platform"
	AuthorizationClient   AuthorizationType = "client"
)

func Valid(ctx *gin.Context, requiredAutyType AuthorizationType, id string) bool {
	subject, exists := ctx.Get("subject")
	if !exists {
		return false
	}

	if id != "" {
		jti, exists := ctx.Get("jti")
		if !exists {
			return false
		}

		return string(requiredAutyType) == subject && jti == id
	}

	return string(requiredAutyType) == subject
}
