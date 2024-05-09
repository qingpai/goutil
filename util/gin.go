package util

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/now"
	"golang.org/x/exp/slices"
	"strconv"
	"time"
)

func ShouldBindJSONWithFields(c *gin.Context, fields []string) (map[string]any, error) {
	result := map[string]interface{}{}

	err := c.ShouldBindJSON(&result)
	if err != nil {
		return nil, err
	}

	for k := range result {
		if !slices.Contains(fields, k) {
			delete(result, k)
		}
	}

	return result, nil
}

func ShouldBindJSONWithOutFields(c *gin.Context, fields []string) (map[string]any, error) {
	result := map[string]interface{}{}

	err := c.ShouldBindJSON(&result)
	if err != nil {
		return nil, err
	}

	for k := range result {
		if slices.Contains(fields, k) {
			delete(result, k)
		}
	}

	return result, nil
}

func DefaultQueryInt(c *gin.Context, key string) *int {
	v := c.DefaultQuery(key, "")
	if v == "" {
		return nil
	}

	if i, err := strconv.Atoi(v); err != nil {
		return nil
	} else {
		return &i
	}
}

func DefaultQueryInt64(c *gin.Context, key string) *int64 {
	v := c.DefaultQuery(key, "")
	if v == "" {
		return nil
	}

	if i, err := strconv.ParseInt(v, 10, 64); err != nil {
		return nil
	} else {
		return &i
	}
}

func DefaultDate(c *gin.Context, key string) *time.Time {
	input := c.DefaultQuery(key, "")
	if input == "" {
		return nil
	}
	loc, _ := time.LoadLocation("Asia/Shanghai")
	input = input + " 00:00:00"

	if v, err := now.ParseInLocation(loc, input); err != nil {
		return nil
	} else {
		return &v
	}
}
