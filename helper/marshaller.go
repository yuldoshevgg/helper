package helper

import "encoding/json"

func JsonToJson(data interface{}, newData interface{}) error {
	body, err := json.Marshal(data)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, newData)
	return err
}
