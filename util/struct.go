package util

import "encoding/json"

func StructToMap(input any) map[string]any {
	b, err := json.Marshal(input)
	if err != nil {
		return nil
	}

	var result map[string]any
	err = json.Unmarshal(b, &result)
	if err != nil {
		return nil
	}

	return result
}
