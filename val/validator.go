package val

import (
	"fmt"
	"net/mail"
	"regexp"
)

var (
	isValidUsername = regexp.MustCompile(`^[a-z0-9_]+$`).MatchString
	isValidFullName = regexp.MustCompile(`^[a-zA-Z\s]+$`).MatchString
)

func ValidateString(val string, minLen int, maxLen int) error {
	n := len(val)
	if n < minLen || n > maxLen {
		return fmt.Errorf("must contain between %d and %d characters", minLen, maxLen)
	}
	return nil
}
func ValidateUsername(val string) error {
	if err := ValidateString(val, 3, 100); err != nil {
		return err
	}
	if !isValidUsername(val) {
		return fmt.Errorf("must contain only lowercase letters, digits, or underscore")
	}
	return nil
}
func ValidateFullName(val string) error {
	if err := ValidateString(val, 3, 100); err != nil {
		return err
	}
	if !isValidFullName(val) {
		return fmt.Errorf("must contain only letters or spaces")
	}
	return nil
}
func ValidatePassword(val string) error {
	return ValidateString(val, 6, 100)
}
func ValidateEmail(val string) error {
	if err := ValidateString(val, 3, 200); err != nil {
		return err
	}
	if _, err := mail.ParseAddress(val); err != nil {
		return fmt.Errorf("is not a valid email address")
	}
	return nil
}
