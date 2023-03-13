package api

import (
	"github.com/go-playground/validator/v10"
	"github.com/sachin-sharma-IN/personalbank/util"
)

// Once validator is done, we'll have to register it with gin.

var validCurrency validator.Func = func(fieldLevel validator.FieldLevel) bool {
	// Since fieldLevel.Field() type is reflection, we'll need to call Interface to get it's value
	// as empty interface. .(string) will convert this value to string.
	if currency, ok := fieldLevel.Field().Interface().(string); ok {
		// Check currency is supported or not.
		return util.IsSupportedCurrency(currency)
	}
	return false
}
