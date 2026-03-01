package domain

type CardFilter struct {
	NameContains  *string
	ManaCost      *int
	ManaCostGte   *int
	Health        *int
	HealthGte     *int
	Attack        *int
	AttackGte     *int
	ClassID       *int
	RarityID      *int
	TypeIDs       []int
	SetID         *int
	KeywordIDsAll []int
}
