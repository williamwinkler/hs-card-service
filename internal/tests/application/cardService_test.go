package application_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/williamwinkler/hs-card-service/internal/application"
	"github.com/williamwinkler/hs-card-service/internal/domain"
	"github.com/williamwinkler/hs-card-service/internal/tests/mocks" // Update this import path based on your project structure
)

func Test_GetCards(t *testing.T) {
	// Mocked card data
	mockCards := []domain.Card{
		{ID: 1, Name: "Card 1"},
		{ID: 2, Name: "Card 2"},
	}

	// Create the manually created card repository
	cardRepoMock := &mocks.CardRepository{
		Cards: make(map[int]domain.Card),
	}

	// Insert mockCards into the repository
	for _, card := range mockCards {
		cardRepoMock.Cards[card.ID] = card
	}

	// Create the CardService with the manually created card repository
	cardService := application.NewCardService(nil, cardRepoMock, nil)

	// Test the GetCards method
	cards, _, err := cardService.GetCards(nil, 1, 100)
	assert.NoError(t, err)
	assert.Equal(t, len(mockCards), len(cards))
	for i, card := range cards {
		assert.Equal(t, mockCards[i].ID, card.ID)
		assert.Equal(t, mockCards[i].Name, card.Name)
	}

	// Assert that the repository was not modified
	actualCards, err := cardRepoMock.FindAll()

	assert.NoError(t, err)
	assert.ElementsMatch(t, mockCards, actualCards)
}
