package domain

type Set struct {
	ID                          int    `json:"id"`
	Name                        string `json:"name"`
	Slug                        string `json:"slug"`
	Type                        string `json:"type"`
	CollectibleCount            int    `json:"collectibleCount"`
	CollectibleRevealedCount    int    `json:"collectibleRevealedCount"`
	NonCollectibleCount         int    `json:"nonCollectibleCount"`
	NonCollectibleRevealedCount int    `json:"nonCollectibleRevealedCount"`
	AliasSetIds                 []int  `json:"aliasSetIds,omitempty"`
}
