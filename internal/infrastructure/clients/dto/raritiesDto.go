package dto

type RaritiesDto []struct {
	Slug         string `json:"slug"`
	ID           int    `json:"id"`
	CraftingCost []int  `json:"craftingCost"`
	DustValue    []int  `json:"dustValue"`
	Name         string `json:"name"`
}
