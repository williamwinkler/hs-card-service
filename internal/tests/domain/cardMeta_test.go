package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/williamwinkler/hs-card-service/internal/domain"
)

func Test_UpdateCard(t *testing.T) {
	assert := assert.New(t)
	cardMeta := domain.CardMeta{}

	prevCard := domain.Card{
		ID:     1,
		Attack: 6,
	}
	newCard := domain.Card{
		ID:     1,
		Attack: 5,
	}

	cardMeta.UpdateCard(prevCard, newCard)

	assert.Equal(1, len(cardMeta.Changes))
	assert.Equal(domain.CardChangeTypeUpdated, cardMeta.Changes[0].Change)
	assert.Equal(prevCard, cardMeta.Changes[0].PrevCard)
	assert.Equal(newCard, cardMeta.Changes[0].NewCard)
}

func Test_AddCard(t *testing.T) {
	assert := assert.New(t)
	cardMeta := domain.CardMeta{}

	newCard := domain.Card{
		ID:     1,
		Attack: 5,
	}

	cardMeta.AddCard(newCard)

	assert.Equal(1, len(cardMeta.Changes))
	assert.Equal(domain.CardChangeTypeAdded, cardMeta.Changes[0].Change)
	assert.Equal(newCard, cardMeta.Changes[0].NewCard)
	assert.Equal(domain.Card{}, cardMeta.Changes[0].PrevCard)
}

func Test_DeleteCard(t *testing.T) {
	assert := assert.New(t)
	cardMeta := domain.CardMeta{}

	prevCard := domain.Card{
		ID:     1,
		Attack: 5,
	}

	cardMeta.DeleteCard(prevCard)

	assert.Equal(1, len(cardMeta.Changes))
	assert.Equal(domain.CardChangeTypeDeleted, cardMeta.Changes[0].Change)
	assert.Equal(prevCard, cardMeta.Changes[0].PrevCard)
	assert.Equal(domain.Card{}, cardMeta.Changes[0].NewCard)
}
