package utils

import (
	"net/mail"
	"regexp"
)

func ValidateEmailAddressFormat(emailAddr string) bool {
	_, err := mail.ParseAddress(emailAddr)
	return err == nil
}

func ValidatePhoneNumberFormat(phoneNumber string) bool {
	for _, c := range phoneNumber {
		if c < '0' || c > '9' {
			return false
		}
	}
	return true
}

func IsValidID(e string) bool {
	regex := regexp.MustCompile("^[a-zA-Z0-9_-]+$")
	return regex.MatchString(e)
}
