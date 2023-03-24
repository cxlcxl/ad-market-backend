package error_catch

import (
	"fmt"
	"market/app/validator"
)

// ValidateErr Key: 'LoginData.Email' Error:Field validation for 'Email' failed on the 'email' tag"
func ValidateErr(err error, prefix string) string {
	return fmt.Sprintf("%s：%s", prefix, validator.Translate(err))
}
