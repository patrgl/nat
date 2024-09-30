package auth

import (
	"errors"
	"os"
	"strings"
)

func ValidateAuthorizationHeader(ah string) error {
	headerParts := strings.Split(ah, " ")

	var token string

	if len(headerParts) == 2 && headerParts[0] == "Bearer" {
		token = headerParts[1]
	} else {
		return errors.New("Invalid authorization header")
	}

	if token == os.Getenv("token") {
		return nil
	} else {
		return errors.New("Incorrect token")
	}
}
