package handlers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCardFilter_SentinelValues(t *testing.T) {
	name := "mage"
	manaCost := int64(99)
	health := int64(99)
	attack := int64(99)
	class := int64(2)
	rarity := int64(4)
	set := int64(10)
	types := []int64{3, 4}
	keywords := []int64{1, 5}

	filter := newCardFilter(&name, &manaCost, &health, &attack, &class, &rarity, types, &set, keywords)

	assert.NotNil(t, filter.NameContains)
	assert.Equal(t, "mage", *filter.NameContains)
	assert.Nil(t, filter.ManaCost)
	assert.Nil(t, filter.Health)
	assert.Nil(t, filter.Attack)
	assert.NotNil(t, filter.ManaCostGte)
	assert.NotNil(t, filter.HealthGte)
	assert.NotNil(t, filter.AttackGte)
	assert.Equal(t, 7, *filter.ManaCostGte)
	assert.Equal(t, 7, *filter.HealthGte)
	assert.Equal(t, 7, *filter.AttackGte)
	assert.Equal(t, []int{3, 4}, filter.TypeIDs)
	assert.Equal(t, []int{1, 5}, filter.KeywordIDsAll)
	assert.Equal(t, 2, *filter.ClassID)
	assert.Equal(t, 4, *filter.RarityID)
	assert.Equal(t, 10, *filter.SetID)
}

func TestNewCardFilter_ExactValues(t *testing.T) {
	manaCost := int64(4)
	health := int64(6)
	attack := int64(2)

	filter := newCardFilter(nil, &manaCost, &health, &attack, nil, nil, nil, nil, nil)

	assert.NotNil(t, filter.ManaCost)
	assert.NotNil(t, filter.Health)
	assert.NotNil(t, filter.Attack)
	assert.Equal(t, 4, *filter.ManaCost)
	assert.Equal(t, 6, *filter.Health)
	assert.Equal(t, 2, *filter.Attack)
	assert.Nil(t, filter.ManaCostGte)
	assert.Nil(t, filter.HealthGte)
	assert.Nil(t, filter.AttackGte)
}
