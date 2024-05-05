package util

import (
	"encoding/json"
	"gorm.io/datatypes"
)

func TypeToDatatypeJson(input interface{}) datatypes.JSON {
	str, _ := json.Marshal(input)

	return str
}

func DatatypeJsonToMap(data datatypes.JSON) (map[string]interface{}, error) {
	var obj map[string]interface{}

	if err := json.Unmarshal(data, &obj); err != nil {
		return nil, err
	}

	return obj, nil
}
