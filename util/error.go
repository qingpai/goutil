package util

import (
	"errors"
	"fmt"
)

func Errorf(text string, args ...interface{}) error {
	return errors.New(fmt.Sprintf(text, args...))
}
