package domain

type Type struct {
	Slug      string `json:"slug"`
	ID        int    `json:"id"`
	Name      string `json:"name"`
	GameModes []int  `json:"gameModes"`
}
