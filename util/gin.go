package util

import (
	"fmt"
	"github.com/araddon/dateparse"
	"github.com/gin-gonic/gin"
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

func DefaultDate(c *gin.Context, key string) (time.Time, error) {
	input := c.DefaultQuery(key, "")
	if input == "" {
		return time.Time{}, fmt.Errorf("没有日期参数: %s", key)
	}
	loc, _ := time.LoadLocation("Asia/Shanghai")

	return dateparse.ParseIn(input, loc)
}
