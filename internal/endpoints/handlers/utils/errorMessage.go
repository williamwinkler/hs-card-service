package utils

import "github.com/williamwinkler/hs-card-service/codegen/models"

func CreateErrorMessage(code int, message string) *models.Error {
	return &models.Error{
		Code:    int64(code),
		Message: &message,
	}
}
