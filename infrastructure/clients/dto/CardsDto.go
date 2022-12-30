package dto

type GetCardsResponse struct {
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
		CopyOfCardID  int           `json:"copyOfCardId,omitempty"`
		Duels         struct {
			Relevant    bool `json:"relevant"`
			Constructed bool `json:"constructed"`
		} `json:"duels,omitempty"`
		MinionTypeID  int   `json:"minionTypeId,omitempty"`
		ChildIds      []int `json:"childIds,omitempty"`
		Durability    int   `json:"durability,omitempty"`
		MultiTypeIds  []int `json:"multiTypeIds,omitempty"`
		SpellSchoolID int   `json:"spellSchoolId,omitempty"`
		Armor         int   `json:"armor,omitempty"`
		RuneCost      struct {
			Blood  int `json:"blood"`
			Frost  int `json:"frost"`
			Unholy int `json:"unholy"`
		} `json:"runeCost,omitempty"`
	} `json:"cards"`
	CardCount int `json:"cardCount"`
	PageCount int `json:"pageCount"`
	Page      int `json:"page"`
}
