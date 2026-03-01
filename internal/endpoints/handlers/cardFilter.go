package handlers

import "github.com/williamwinkler/hs-card-service/internal/domain"

func newCardFilter(name *string, manaCost *int64, health *int64, attack *int64, class *int64, rarity *int64, types []int64, set *int64, keywords []int64) domain.CardFilter {
	filter := domain.CardFilter{}

	if name != nil {
		value := *name
		filter.NameContains = &value
	}

	if manaCost != nil {
		if *manaCost == 99 {
			value := 7
			filter.ManaCostGte = &value
		} else {
			value := int(*manaCost)
			filter.ManaCost = &value
		}
	}

	if health != nil {
		if *health == 99 {
			value := 7
			filter.HealthGte = &value
		} else {
			value := int(*health)
			filter.Health = &value
		}
	}

	if attack != nil {
		if *attack == 99 {
			value := 7
			filter.AttackGte = &value
		} else {
			value := int(*attack)
			filter.Attack = &value
		}
	}

	if class != nil {
		value := int(*class)
		filter.ClassID = &value
	}
	if rarity != nil {
		value := int(*rarity)
		filter.RarityID = &value
	}
	if set != nil {
		value := int(*set)
		filter.SetID = &value
	}

	if len(types) > 0 {
		filter.TypeIDs = make([]int, 0, len(types))
		for _, value := range types {
			filter.TypeIDs = append(filter.TypeIDs, int(value))
		}
	}

	if len(keywords) > 0 {
		filter.KeywordIDsAll = make([]int, 0, len(keywords))
		for _, value := range keywords {
			filter.KeywordIDsAll = append(filter.KeywordIDsAll, int(value))
		}
	}

	return filter
}
