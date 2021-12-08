package api

import (
	"github.com/AntonioMorales97/ecom-go/pkg/util"
	"github.com/go-playground/validator/v10"
)

var validProductTypeID validator.Func = func(fieldLevel validator.FieldLevel) bool {
	if productTypeID, ok := fieldLevel.Field().Interface().(int64); ok {
		return util.IsSupportedProductType(productTypeID)
	}
	return false
}
