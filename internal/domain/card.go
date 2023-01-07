package domain

import "reflect"

type Card struct {
	ID            int           `json:"id"`
	Collectible   int           `json:"collectible"`
	Slug          string        `json:"slug"`
	ClassID       int           `json:"classId"`
	MultiClassIds []interface{} `json:"multiClassIds"`
	CardTypeID    int           `json:"cardTypeId"`
	CardSetID     int           `json:"cardSetId"`
	RarityID      int           `json:"rarityId"`
	ArtistName    string        `json:"artistName"`
	Health        int           `json:"health"`
	Attack        int           `json:"attack"`
	ManaCost      int           `json:"manaCost"`
	Name          string        `json:"name"`
	Text          string        `json:"text"`
	Image         string        `json:"image"`
	ImageGold     string        `json:"imageGold"`
	FlavorText    string        `json:"flavorText"`
	CropImage     string        `json:"cropImage"`
	ParentID      int           `json:"parentId"`
	KeywordIds    []int         `json:"keywordIds"`
	Duels         Duels         `json:"duels"`
}

type Duels struct {
	Relevant    bool `json:"relevant"`
	Constructed bool `json:"constructed"`
}

func (c *Card) Equals(card2 Card) bool {
	return c.ID == card2.ID &&
		c.Collectible == card2.Collectible &&
		c.Slug == card2.Slug &&
		c.ClassID == card2.ClassID &&
		reflect.DeepEqual(c.MultiClassIds, card2.MultiClassIds) &&
		c.CardTypeID == card2.CardTypeID &&
		c.CardSetID == card2.CardSetID &&
		c.RarityID == card2.RarityID &&
		c.ArtistName == card2.ArtistName &&
		c.Health == card2.Health &&
		c.Attack == card2.Attack &&
		c.ManaCost == card2.ManaCost &&
		c.Name == card2.Name &&
		c.Text == card2.Text &&
		c.Image == card2.Image &&
		c.ImageGold == card2.ImageGold &&
		c.FlavorText == card2.FlavorText &&
		c.CropImage == card2.CropImage &&
		c.ParentID == card2.ParentID &&
		reflect.DeepEqual(c.KeywordIds, card2.KeywordIds) &&
		c.Duels.Relevant == card2.Duels.Relevant &&
		c.Duels.Constructed == card2.Duels.Constructed
}
