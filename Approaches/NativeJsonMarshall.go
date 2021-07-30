package Approaches

import "encoding/json"

func ConvertWithNativeLib(input interface{}, output interface{}) error {

	body, err := json.Marshal(input)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, &output)
	if err != nil {
		return err
	}
	return nil
}
