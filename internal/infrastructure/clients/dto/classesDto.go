package dto

type ClassesDto []struct {
	Slug                 string `json:"slug"`
	ID                   int    `json:"id"`
	Name                 string `json:"name"`
	CardID               int    `json:"cardId,omitempty"`
	HeroPowerCardID      int    `json:"heroPowerCardId,omitempty"`
	AlternateHeroCardIds []int  `json:"alternateHeroCardIds,omitempty"`
}
