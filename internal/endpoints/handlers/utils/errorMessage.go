package utils

import "github.com/williamwinkler/hs-card-service/codegen/models"

// Returns a *models.Error with the error code and optional message
// If no message is given the default message will be: "Something went wrong. Try again later"
func CreateErrorMessage(code int, message ...string) *models.Error {
	m := "Something went wrong. Try again later"
	if len(message) != 0 {
		m = message[0]
	}

	return &models.Error{
		Code:    int64(code),
		Message: &m,
	}
}
