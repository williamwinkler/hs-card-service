package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/williamwinkler/hs-card-service/internal/domain"
)

func Test_Equals(t *testing.T) {
	tests := []struct {
		card1    domain.Card
		card2    domain.Card
		expected bool
	}{
		{
			card1:    domain.Card{ID: 1},
			card2:    domain.Card{ID: 1},
			expected: true,
		},
		{
			card1:    domain.Card{ID: 1, KeywordIds: []int{1, 2, 3}},
			card2:    domain.Card{ID: 1, KeywordIds: []int{1, 2, 3}},
			expected: true,
		},
		{
			card1:    domain.Card{Duels: domain.Duels{Relevant: true, Constructed: false}},
			card2:    domain.Card{Duels: domain.Duels{Relevant: true, Constructed: false}},
			expected: true,
		},
		{
			card1:    domain.Card{Duels: domain.Duels{Relevant: true, Constructed: false}, Name: "test", KeywordIds: []int{1, 2, 3}},
			card2:    domain.Card{Name: "test", KeywordIds: []int{1, 2, 3}, Duels: domain.Duels{Relevant: true, Constructed: false}},
			expected: true,
		},
		{
			card1:    domain.Card{Duels: domain.Duels{Relevant: true, Constructed: false}, Name: "test", KeywordIds: []int{1, 2, 3}},
			card2:    domain.Card{Duels: domain.Duels{Relevant: true, Constructed: false}, Name: "test", KeywordIds: []int{1, 2, 3, 4}},
			expected: false,
		},
	}
	for _, test := range tests {
		assert := assert.New(t)

		actual := test.card1.Equals(test.card2)

		assert.Equal(test.expected, actual)
	}
}
