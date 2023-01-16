package dto

import (
	"github.com/williamwinkler/hs-card-service/internal/domain"
)

func MapToCards(cardsDto CardsDto) []domain.Card {
	var cards []domain.Card
	for _, p := range cardsDto.Cards {
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

func MapToSets(setsDto SetsDto) []domain.Set {
	var sets []domain.Set
	for _, r := range setsDto {
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

func MapToClasses(classesDto ClassesDto) []domain.Class {
	var classes []domain.Class
	for _, r := range classesDto {
		var s domain.Class
		s.Slug = r.Slug
		s.ID = r.ID
		s.Name = r.Name
		s.CardID = r.CardID
		s.HeroPowerCardID = r.HeroPowerCardID
		s.AlternateHeroCardIds = r.AlternateHeroCardIds

		classes = append(classes, s)
	}
	return classes
}

func MapToRariteis(raritiesDto RaritiesDto) []domain.Rarity {
	var rarities []domain.Rarity
	for _, r := range raritiesDto {
		var s domain.Rarity
		s.Slug = r.Slug
		s.ID = r.ID
		s.CraftingCost = r.CraftingCost
		s.DustValue = r.DustValue
		s.Name = r.Name

		rarities = append(rarities, s)
	}
	return rarities
}

func MapToTypes(typesDto TypesDto) []domain.Type {
	var types []domain.Type
	for _, r := range typesDto {
		var s domain.Type
		s.Slug = r.Slug
		s.ID = r.ID
		s.Name = r.Name
		s.GameModes = r.GameModes

		types = append(types, s)
	}
	return types
}

func MapToKeywords(keywordsDto KeywordsDto) []domain.Keyword {
	var keywords []domain.Keyword
	for _, r := range keywordsDto {
		var s domain.Keyword
		s.ID = r.ID
		s.Slug = r.Slug
		s.Name = r.Name
		s.RefText = r.RefText
		s.Text = r.Text
		s.GameModes = r.GameModes

		keywords = append(keywords, s)
	}
	return keywords
}
