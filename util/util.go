package util

import (
	"code.qingpai365.com/erp/goutil/validator"
	"fmt"
	"net/http"
	"reflect"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func Split(input string, char rune) []string {
	return strings.FieldsFunc(input, func(r rune) bool {
		return r == char
	})
}

func IfThenElse(condition bool, a interface{}, b interface{}) interface{} {
	if condition {
		return a
	}
	return b
}

func IsNil(i interface{}) bool {
	return i == nil || reflect.ValueOf(i).IsNil()
}

func GetClientIp(c *gin.Context) string {
	return c.ClientIP()
}

func CurrentDateTimeString() string {
	t := time.Now()
	return t.Format("20060102150405")
}

func ResponseMsg(c *gin.Context, msg string) {
	c.JSON(200, gin.H{
		"code":    0,
		"message": msg,
	})
}

func ResponseSuccess(c *gin.Context, data any) {
	c.JSON(200, gin.H{
		"code":    0,
		"message": "",
		"data":    data,
	})
}

func ResponseEmptyList(c *gin.Context) {
	c.JSON(200, gin.H{
		"code":    0,
		"message": "",
		"data":    make([]interface{}, 0),
	})
}

func ResponseErr(c *gin.Context, err error) {
	c.JSON(400, gin.H{
		"code":    1,
		"message": validator.Translate(err),
	})
}

func ResponseErrString(c *gin.Context, message string) {
	c.JSON(400, gin.H{
		"code":    1,
		"message": message,
	})
}

func ResponseErrf(c *gin.Context, template string, args ...interface{}) {
	c.JSON(400, gin.H{
		"code":    1,
		"message": fmt.Sprintf(template, args...),
	})
}

func ResponseErrData[T any](c *gin.Context, message string, data T) {
	c.JSON(400, gin.H{
		"code":    1,
		"message": message,
		"data":    data,
	})
}

func ResponseStatus(c *gin.Context, code int) {
	c.AbortWithStatus(code)
}

func ShowSuccess(c *gin.Context, message string) {
	c.Redirect(http.StatusMovedPermanently, fmt.Sprintf("/api/static/message.html?type=success&message=%s&timestamp=%d", message, time.Now().Unix()))
}

func ShowError(c *gin.Context, message string) {
	c.Redirect(http.StatusMovedPermanently, fmt.Sprintf("/api/static/message.html?type=error&message=%s&timestamp=%d", message, time.Now().Unix()))
}
