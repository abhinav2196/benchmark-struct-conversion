package Approaches

import (
	"BenchmarkStructConversion/Models"
	"github.com/jinzhu/copier"
)

func ConvertWithCopier(input Models.Response1, output *Models.Response1) error {
	err := copier.Copy(&output, input)

	if err != nil {
		return err
	}
	return nil
}
