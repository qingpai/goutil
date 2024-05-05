package util

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slices"
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
