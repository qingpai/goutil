package util

import (
	"github.com/google/uuid"
	"strings"
)

func NewUuid() string {
	return strings.ReplaceAll(uuid.New().String(), "-", "")
}
