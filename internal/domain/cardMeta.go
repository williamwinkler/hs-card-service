package domain

import "time"

type CardMeta struct {
	Updated   time.Time
	IsChanged bool
	Changes   []CardChange
}

type CardChange struct {
	PrevCard Card
	NewCard  Card
	Change   CardChangeType
}

type CardChangeType string

const (
	CardChangeTypeUpdated CardChangeType = "Updated"
	CardChangeTypeDeleted CardChangeType = "Deleted"
	CardChangeTypeAdded    CardChangeType = "Added"
)

func (c *CardMeta) UpdateCard(prevCard Card, newCard Card) {
	cardChange := CardChange{
		PrevCard: prevCard,
		NewCard:  newCard,
		Change:   CardChangeTypeUpdated,
	}
	c.Changes = append(c.Changes, cardChange)
	c.IsChanged = true
}

func (c *CardMeta) DeleteCard(card Card) {
	cardChange := CardChange{
		PrevCard: card,
		Change:   CardChangeTypeDeleted,
	}
	c.Changes = append(c.Changes, cardChange)
	c.IsChanged = true
}

func (c *CardMeta) AddCard(card Card) {
	cardChange := CardChange{
		NewCard: card,
		Change:  CardChangeTypeAdded,
	}
	c.Changes = append(c.Changes, cardChange)
	c.IsChanged = true
}
