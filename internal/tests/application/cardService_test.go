package application

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/williamwinkler/hs-card-service/internal/application"
	"github.com/williamwinkler/hs-card-service/internal/domain"
	"github.com/williamwinkler/hs-card-service/internal/tests/mocks"
)

// TODO: enable this test once UpdateCards is handling updates to cards appropriately
func _Test_UpdateCards(t *testing.T) {
	// Arrange
	newCards := []domain.Card{
		{
			ID:   1,
			Name: "card_a",
		},
		{
			ID:   2,
			Name: "card_b_UPDATED",
		},
		{
			ID:   4,
			Name: "card_d",
		},
		{
			ID:   5,
			Name: "card_e",
		},
	}
	newCardsMap := make(map[int]domain.Card)
	for _, card := range newCards {
		newCardsMap[card.ID] = card
	}

	oldCards := []domain.Card{
		{
			ID:   1,
			Name: "card_a",
		},
		{
			ID:   2,
			Name: "card_b",
		},
		{
			ID:   3,
			Name: "card_c",
		},
	}

	assert := assert.New(t)
	cardRepo := mocks.NewCardRepository()
	hsClient := mocks.HsClient{}
	cardMetaRepo := mocks.UpdateMetaRepository{}
	cardService := application.NewCardService(&hsClient, &cardRepo, &cardMetaRepo)
	cardRepo.InsertMany(oldCards)
	hsClient.On("GetAllCards").Return(newCards, nil)
	cardMetaRepo.On("InsertOne", mock.AnythingOfType("domain.CardMeta")).Return(nil)

	// Act
	err := cardService.Update()
	assert.Nil(err)

	// Assert
	cards, _ := cardRepo.FindAll()
	for _, card := range cards {
		assert.Equal(newCardsMap[card.ID], card)
	}
}
