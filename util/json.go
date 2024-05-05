package util

import (
	"encoding/json"
	"golang.org/x/exp/maps"
	"gorm.io/datatypes"
	"strings"
)

type JsonParse struct {
}

func (m *JsonParse) ParseObject(input string, keyToFind []string) map[string]EyeData {
	eyeDataMap := make(map[string]EyeData)

	if strings.HasPrefix(input, "{") {
		var data map[string]any
		if err := json.Unmarshal([]byte(input), &data); err != nil {
			return nil
		}
		maps.Copy(eyeDataMap, m.parseMap(data, keyToFind))
	} else if strings.HasPrefix(input, "[") {
		var data []any
		if err := json.Unmarshal([]byte(input), &data); err != nil {
			return nil
		}
		maps.Copy(eyeDataMap, m.parseArray(data, keyToFind))
	} else {

	}

	return eyeDataMap
}

func (m *JsonParse) parseMap(aMap map[string]any, keys []string) map[string]EyeData {
	eyeDataMap := make(map[string]EyeData, 0)

	for key, val := range aMap {
		for _, keyToFind := range keys {
			if strings.Contains(key, keyToFind) {
				if val != nil {
					valString, err := json.Marshal(val)
					if err != nil {
						continue
					}

					var eyeData EyeData
					if err := json.Unmarshal(valString, &eyeData); err != nil {
						continue
					}
					if eyeData.Od != "" || eyeData.Os != "" {
						eyeDataMap[key] = eyeData
					}
				}
			}
		}

		switch val.(type) {
		case map[string]any:
			maps.Copy(eyeDataMap, m.parseMap(val.(map[string]any), keys))
		case []any:
			maps.Copy(eyeDataMap, m.parseArray(val.([]any), keys))
		}
	}

	return eyeDataMap
}

func (m *JsonParse) parseArray(anArray []any, keyToFind []string) map[string]EyeData {
	eyeDataMap := make(map[string]EyeData, 0)
	for _, val := range anArray {
		switch val.(type) {
		case map[string]any:
			maps.Copy(eyeDataMap, m.parseMap(val.(map[string]any), keyToFind))
		case []any:
			maps.Copy(eyeDataMap, m.parseArray(val.([]any), keyToFind))
		}
	}

	return eyeDataMap
}

func ConvertToTarget[T any](input any, target *T) error {
	bytes, err := json.Marshal(input)
	if err != nil {
		return err
	}

	if err = json.Unmarshal(bytes, &target); err != nil {
		return err
	}

	return nil
}

func MarshalToDatabaseJson[T any](input T) datatypes.JSON {
	v, _ := json.Marshal(input)

	return v
}

func UnMarshalFromDatabaseJson(input []byte) (map[string]any, error) {
	var result map[string]any
	if err := json.Unmarshal(input, &result); err != nil {
		return nil, err
	}

	return result, nil
}

func UnMarshalType[T any](data datatypes.JSON) (*T, error) {
	var t *T
	input, err := data.MarshalJSON()
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(input, &t); err != nil {
		return nil, err
	}

	return t, nil
}
