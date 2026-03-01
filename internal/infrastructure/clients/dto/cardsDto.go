package dto

import (
	"encoding/json"
	"fmt"
)

type CopyOfCardID struct {
	Values []int
}

func (c *CopyOfCardID) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		c.Values = nil
		return nil
	}

	var single int
	if err := json.Unmarshal(data, &single); err == nil {
		c.Values = []int{single}
		return nil
	}

	var multiple []int
	if err := json.Unmarshal(data, &multiple); err == nil {
		c.Values = multiple
		return nil
	}

	return fmt.Errorf("copyOfCardId must be integer or integer array")
}

func (c CopyOfCardID) HasValue() bool {
	for _, value := range c.Values {
		if value != 0 {
			return true
		}
	}
	return false
}

type CardsDto struct {
	Cards []struct {
		ID            int           `json:"id"`
		Collectible   int           `json:"collectible"`
		Slug          string        `json:"slug"`
		ClassID       int           `json:"classId"`
		MultiClassIds []interface{} `json:"multiClassIds"`
		CardTypeID    int           `json:"cardTypeId"`
		CardSetID     int           `json:"cardSetId"`
		RarityID      int           `json:"rarityId"`
		ArtistName    string        `json:"artistName"`
		Health        int           `json:"health,omitempty"`
		Attack        int           `json:"attack,omitempty"`
		ManaCost      int           `json:"manaCost"`
		Name          string        `json:"name"`
		Text          string        `json:"text"`
		Image         string        `json:"image"`
		ImageGold     string        `json:"imageGold"`
		FlavorText    string        `json:"flavorText"`
		CropImage     string        `json:"cropImage"`
		ParentID      int           `json:"parentId"`
		KeywordIds    []int         `json:"keywordIds,omitempty"`
		CopyOfCardID  CopyOfCardID  `json:"copyOfCardId,omitempty"`
		Duels         struct {
			Relevant    bool `json:"relevant"`
			Constructed bool `json:"constructed"`
		} `json:"duels,omitempty"`
		MinionTypeID              int   `json:"minionTypeId,omitempty"`
		ChildIds                  []int `json:"childIds,omitempty"`
		Durability                int   `json:"durability,omitempty"`
		MultiTypeIds              []int `json:"multiTypeIds,omitempty"`
		SpellSchoolID             int   `json:"spellSchoolId,omitempty"`
		Armor                     int   `json:"armor,omitempty"`
		IsZilliaxFunctionalModule bool  `json:"isZilliaxFunctionalModule,omitempty"`
		IsZilliaxCosmeticModule   bool  `json:"isZilliaxCosmeticModule,omitempty"`
		RuneCost                  struct {
			Blood  int `json:"blood"`
			Frost  int `json:"frost"`
			Unholy int `json:"unholy"`
		} `json:"runeCost,omitempty"`
	} `json:"cards"`
	CardCount int `json:"cardCount"`
	PageCount int `json:"pageCount"`
	Page      int `json:"page"`
}
