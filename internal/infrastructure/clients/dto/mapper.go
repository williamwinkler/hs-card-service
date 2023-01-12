package dto

import "github.com/williamwinkler/hs-card-service/internal/domain"

func MapToCards(cardResp CardsDto) []domain.Card {
	var cards []domain.Card
	for _, p := range cardResp.Cards {
		if p.CopyOfCardID != 0 {
			continue // if CopyOfCard is not 0, it's an unwanted outdated version
		}

		var c domain.Card
		c.ID = p.ID
		c.Collectible = p.Collectible
		c.Slug = p.Slug
		c.ClassID = p.ClassID
		c.MultiClassIds = p.MultiClassIds
		c.CardTypeID = p.CardTypeID
		c.CardSetID = p.CardSetID
		c.RarityID = p.RarityID
		c.ArtistName = p.ArtistName
		c.Health = p.Health
		c.Attack = p.Attack
		c.ManaCost = p.ManaCost
		c.Name = p.Name
		c.Text = p.Text
		c.Image = p.Image
		c.ImageGold = p.ImageGold
		c.FlavorText = p.FlavorText
		c.CropImage = p.CropImage
		c.ParentID = p.ParentID
		c.KeywordIds = p.KeywordIds
		c.Duels = p.Duels

		cards = append(cards, c)
	}
	return cards
}

func MapToSets(setsResp SetsDto) []domain.Set {
	var sets []domain.Set
	for _, r := range setsResp {
		var s domain.Set
		s.ID = r.ID
		s.Name = r.Name
		s.Slug = r.Slug
		s.Type = r.Type
		s.CollectibleCount = r.CollectibleCount
		s.CollectibleRevealedCount = r.CollectibleRevealedCount
		s.NonCollectibleCount = r.NonCollectibleCount
		s.NonCollectibleRevealedCount = r.NonCollectibleRevealedCount
		s.AliasSetIds = r.AliasSetIds

		sets = append(sets, s)
	}

	return sets
}
