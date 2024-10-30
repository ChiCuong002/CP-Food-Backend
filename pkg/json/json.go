package json

import "encoding/json"

func JSONMarshal(v interface{}) ([]byte, error) {
	jsonData, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}
	return jsonData, nil
}