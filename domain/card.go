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

func (c *Card) IsEqual(p Card) bool {
	return reflect.DeepEqual(c, p)
}
