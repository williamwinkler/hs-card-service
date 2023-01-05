package application

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/williamwinkler/hs-card-service/internal/application"
	"github.com/williamwinkler/hs-card-service/internal/domain"
	"github.com/williamwinkler/hs-card-service/internal/tests/mocks"
)

func Test_UpdateCards(t *testing.T) {
	// Arrange
	assert := assert.New(t)
	cardRepo := mocks.CardRepository{}
	hsClient := mocks.HsClient{}
	cardService := application.NewCardService(&hsClient, &cardRepo)
	cardRepo.On("FindAll").Return([]domain.Card{}, nil)
	hsClient.On("GetAllCards").Return([]domain.Card{}, nil)

	// Act
	err := cardService.UpdateCards()

	// Assert
	assert.Nil(err)
}
