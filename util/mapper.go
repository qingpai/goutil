package util

import (
	"fmt"
	"github.com/petersunbag/coven"
	"reflect"
	"sync"
)

var (
	mutex sync.Mutex
	cMap  = make(map[string]*coven.Converter)
)

func Map(src, dst any) (err error) {
	key := fmt.Sprintf("%v_%v", reflect.TypeOf(src).String(), reflect.TypeOf(dst).String())
	if _, ok := cMap[key]; !ok {
		mutex.Lock()
		defer mutex.Unlock()
		if cMap[key], err = coven.NewConverter(dst, src); err != nil {
			return
		}
	}
	if err = cMap[key].Convert(dst, src); err != nil {
		return
	}
	return
}

// 转换map key为数据库字段, 方便更新数据库
func ToUpdateMap(data map[string]interface{}) map[string]interface{} {
	m := make(map[string]interface{})
	for k, v := range data {
		if v != nil {
			m[ToSnakeCase(k)] = v
		}
	}
	return m
}
