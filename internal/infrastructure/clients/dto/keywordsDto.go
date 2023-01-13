package dto

type KeywordsDto []struct {
	ID        int    `json:"id"`
	Slug      string `json:"slug"`
	Name      string `json:"name"`
	RefText   string `json:"refText"`
	Text      string `json:"text"`
	GameModes []int  `json:"gameModes"`
}
