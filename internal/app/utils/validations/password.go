package validations

import (
	"errors"
	"unicode"
)

func Password(password string) error {
	var (
		hasMinLen    = false
		hasMaxLen    = false
		hasUppercase = false
		hasLowercase = false
		hasNumber    = false
		hasSpecial   = false
	)

	if len(password) >= 8 {
		hasMinLen = true
	}

	if len(password) <= 32 {
		hasMaxLen = true
	}

	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			hasUppercase = true
		case unicode.IsLower(char):
			hasLowercase = true
		case unicode.IsNumber(char):
			hasNumber = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		}
	}

	if hasMinLen && hasMaxLen && hasUppercase && hasLowercase && hasNumber && hasSpecial {
		return nil
	}

	return errors.New("password does not meet the requirements")
}
