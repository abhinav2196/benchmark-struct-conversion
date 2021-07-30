package Approaches

import "BenchmarkStructConversion/Models"

func ConvertWithEasyJson(input Models.Response1, output *Models.Response1) error {
	body, err := input.MarshalJSON()
	if err != nil {
		return err
	}
	err = output.UnmarshalJSON(body)
	if err != nil {
		return err
	}
	return nil
}
