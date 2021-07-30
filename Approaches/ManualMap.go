package Approaches

import "BenchmarkStructConversion/Models"

func ConvertWithManualMapping(input *Models.Response1, output *Models.Response1) error {
	output.Page = input.Page
	output.Fruits = input.Fruits
	newStruct := &Models.Response2{input.Z.X, input.Z.Y}
	output.Z = newStruct
	return nil
}
