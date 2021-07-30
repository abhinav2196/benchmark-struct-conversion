package Approaches

import (
	jsoniter "github.com/json-iterator/go"
)

func ConvertWithJsonIter(input interface{}, output interface{}) error {

	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	//var json = jsoniter.ConfigFastest

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
