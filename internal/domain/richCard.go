package domain

type RichCard struct {
	ID            int           `json:"id"`
	Collectible   int           `json:"collectible"`
	Slug          string        `json:"slug"`
	Class         string        `json:"class"`
	MultiClassIds []interface{} `json:"multiClassIds"`
	CardType      string        `json:"cardType"`
	CardSet       string        `json:"cardSet"`
	Rarity        string        `json:"rarity"`
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
	Keywords      []string      `json:"keyword"`
	Duels         Duels         `json:"duels"`
}
