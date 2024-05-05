package util

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetParamInt64(c *gin.Context, key string) (int64, error) {
	k := c.Param(key)
	if k == "" {
		return 0, errors.New(fmt.Sprintf("%s empty", key))
	}
	id, err := strconv.ParseInt(k, 10, 64)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func GetParamInt64Arr(c *gin.Context, key string) ([]int64, error) {
	k := c.Param(key)
	if k == "" {
		return nil, errors.New(fmt.Sprintf("%s empty", key))
	}

	return SplitToIntArr(k)
}

func GetQueryint64(c *gin.Context, key string) (int64, error) {
	k := c.DefaultQuery(key, "")
	if k == "" {
		return 0, errors.New(fmt.Sprintf("%s empty", key))
	}
	id, err := strconv.ParseInt(k, 10, 64)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func GetQueryMap(c *gin.Context, key []string) map[string]string {
	params := make(map[string]string)
	for _, v := range key {
		if param := c.DefaultQuery(v, ""); param != "" {
			params[ToSnakeCase(v)] = param
		}
	}

	return params
}

func GetQuerySlice(c *gin.Context, key string) (int64, error) {
	k := c.DefaultQuery(key, "")
	if k == "" {
		return 0, errors.New(fmt.Sprintf("%s empty", key))
	}
	id, err := strconv.ParseInt(k, 10, 64)
	if err != nil {
		return 0, err
	}

	return id, nil
}
