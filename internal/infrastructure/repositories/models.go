package repositories

import (
	"encoding/json"
	"time"

	"github.com/lib/pq"
	"github.com/williamwinkler/hs-card-service/internal/domain"
	"gorm.io/datatypes"
)

type cardRecord struct {
	ID                        int            `gorm:"column:id;primaryKey"`
	Collectible               int            `gorm:"column:collectible"`
	Slug                      string         `gorm:"column:slug"`
	ClassID                   int            `gorm:"column:classid"`
	MultiClassIDs             datatypes.JSON `gorm:"column:multiclassids;type:jsonb"`
	SpellSchoolID             int            `gorm:"column:spellschoolid"`
	CardTypeID                int            `gorm:"column:cardtypeid"`
	CardSetID                 int            `gorm:"column:cardsetid"`
	RarityID                  int            `gorm:"column:rarityid"`
	ArtistName                string         `gorm:"column:artistname"`
	Health                    int            `gorm:"column:health"`
	Attack                    int            `gorm:"column:attack"`
	ManaCost                  int            `gorm:"column:manacost"`
	Name                      string         `gorm:"column:name"`
	Text                      string         `gorm:"column:text"`
	Image                     string         `gorm:"column:image"`
	ImageGold                 string         `gorm:"column:imagegold"`
	FlavorText                string         `gorm:"column:flavortext"`
	CropImage                 string         `gorm:"column:cropimage"`
	ParentID                  int            `gorm:"column:parentid"`
	CopyOfCardIDs             pq.Int64Array  `gorm:"column:copyofcardids;type:integer[]"`
	MinionTypeID              int            `gorm:"column:miniontypeid"`
	ChildIDs                  pq.Int64Array  `gorm:"column:childids;type:integer[]"`
	Durability                int            `gorm:"column:durability"`
	MultiTypeIDs              pq.Int64Array  `gorm:"column:multitypeids;type:integer[]"`
	Armor                     int            `gorm:"column:armor"`
	IsZilliaxFunctionalModule bool           `gorm:"column:iszilliaxfunctionalmodule"`
	IsZilliaxCosmeticModule   bool           `gorm:"column:iszilliaxcosmeticmodule"`
	DualsRelevant             bool           `gorm:"column:duals_relevant"`
	DualsConstructed          bool           `gorm:"column:duals_constructed"`
}

func (cardRecord) TableName() string { return "cards" }

type cardKeywordRecord struct {
	CardID    int `gorm:"column:card_id;primaryKey"`
	KeywordID int `gorm:"column:keyword_id;primaryKey"`
}

func (cardKeywordRecord) TableName() string { return "card_keywords" }

type setRecord struct {
	ID                          int           `gorm:"column:id;primaryKey"`
	Name                        string        `gorm:"column:name"`
	Slug                        string        `gorm:"column:slug"`
	Type                        string        `gorm:"column:type"`
	CollectibleCount            int           `gorm:"column:collectible_count"`
	CollectibleRevealedCount    int           `gorm:"column:collectible_revealed_count"`
	NonCollectibleCount         int           `gorm:"column:non_collectible_count"`
	NonCollectibleRevealedCount int           `gorm:"column:non_collectible_revealed_count"`
	AliasSetIDs                 pq.Int64Array `gorm:"column:alias_set_ids;type:integer[]"`
}

func (setRecord) TableName() string { return "sets" }

type classRecord struct {
	Slug                 string        `gorm:"column:slug"`
	ID                   int           `gorm:"column:id;primaryKey"`
	Name                 string        `gorm:"column:name"`
	CardID               int           `gorm:"column:card_id"`
	HeroPowerCardID      int           `gorm:"column:hero_power_card_id"`
	AlternateHeroCardIDs pq.Int64Array `gorm:"column:alternate_hero_card_ids;type:integer[]"`
}

func (classRecord) TableName() string { return "classes" }

type rarityRecord struct {
	Slug         string        `gorm:"column:slug"`
	ID           int           `gorm:"column:id;primaryKey"`
	CraftingCost pq.Int64Array `gorm:"column:crafting_cost;type:integer[]"`
	DustValue    pq.Int64Array `gorm:"column:dust_value;type:integer[]"`
	Name         string        `gorm:"column:name"`
}

func (rarityRecord) TableName() string { return "rarities" }

type typeRecord struct {
	Slug      string        `gorm:"column:slug"`
	ID        int           `gorm:"column:id;primaryKey"`
	Name      string        `gorm:"column:name"`
	GameModes pq.Int64Array `gorm:"column:game_modes;type:integer[]"`
}

func (typeRecord) TableName() string { return "types" }

type keywordRecord struct {
	ID        int           `gorm:"column:id;primaryKey"`
	Slug      string        `gorm:"column:slug"`
	Name      string        `gorm:"column:name"`
	RefText   string        `gorm:"column:ref_text"`
	Text      string        `gorm:"column:text"`
	GameModes pq.Int64Array `gorm:"column:game_modes;type:integer[]"`
}

func (keywordRecord) TableName() string { return "keywords" }

type updateMetaRecord struct {
	ID        uint64    `gorm:"column:id;primaryKey;autoIncrement"`
	UpdatedAt time.Time `gorm:"column:updated"`
	IsChanged bool      `gorm:"column:is_changed"`
}

func (updateMetaRecord) TableName() string { return "update_meta" }

func toCardRecord(card domain.Card) (cardRecord, error) {
	encodedMultiClassIDs, err := json.Marshal(card.MultiClassIds)
	if err != nil {
		return cardRecord{}, err
	}

	return cardRecord{
		ID:                        card.ID,
		Collectible:               card.Collectible,
		Slug:                      card.Slug,
		ClassID:                   card.ClassID,
		MultiClassIDs:             datatypes.JSON(encodedMultiClassIDs),
		SpellSchoolID:             card.SpellSchoolID,
		CardTypeID:                card.CardTypeID,
		CardSetID:                 card.CardSetID,
		RarityID:                  card.RarityID,
		ArtistName:                card.ArtistName,
		Health:                    card.Health,
		Attack:                    card.Attack,
		ManaCost:                  card.ManaCost,
		Name:                      card.Name,
		Text:                      card.Text,
		Image:                     card.Image,
		ImageGold:                 card.ImageGold,
		FlavorText:                card.FlavorText,
		CropImage:                 card.CropImage,
		ParentID:                  card.ParentID,
		CopyOfCardIDs:             toInt64Array(card.CopyOfCardIDs),
		MinionTypeID:              card.MinionTypeID,
		ChildIDs:                  toInt64Array(card.ChildIDs),
		Durability:                card.Durability,
		MultiTypeIDs:              toInt64Array(card.MultiTypeIDs),
		Armor:                     card.Armor,
		IsZilliaxFunctionalModule: card.IsZilliaxFunctionalModule,
		IsZilliaxCosmeticModule:   card.IsZilliaxCosmeticModule,
		DualsRelevant:             card.Duels.Relevant,
		DualsConstructed:          card.Duels.Constructed,
	}, nil
}

func toDomainCard(record cardRecord, keywordIDs []int) domain.Card {
	var multiClassIDs []interface{}
	if len(record.MultiClassIDs) > 0 {
		_ = json.Unmarshal(record.MultiClassIDs, &multiClassIDs)
	}

	return domain.Card{
		ID:                        record.ID,
		Collectible:               record.Collectible,
		Slug:                      record.Slug,
		ClassID:                   record.ClassID,
		MultiClassIds:             multiClassIDs,
		SpellSchoolID:             record.SpellSchoolID,
		CardTypeID:                record.CardTypeID,
		CardSetID:                 record.CardSetID,
		RarityID:                  record.RarityID,
		ArtistName:                record.ArtistName,
		Health:                    record.Health,
		Attack:                    record.Attack,
		ManaCost:                  record.ManaCost,
		Name:                      record.Name,
		Text:                      record.Text,
		Image:                     record.Image,
		ImageGold:                 record.ImageGold,
		FlavorText:                record.FlavorText,
		CropImage:                 record.CropImage,
		ParentID:                  record.ParentID,
		KeywordIds:                keywordIDs,
		CopyOfCardIDs:             fromInt64Array(record.CopyOfCardIDs),
		MinionTypeID:              record.MinionTypeID,
		ChildIDs:                  fromInt64Array(record.ChildIDs),
		Durability:                record.Durability,
		MultiTypeIDs:              fromInt64Array(record.MultiTypeIDs),
		Armor:                     record.Armor,
		IsZilliaxFunctionalModule: record.IsZilliaxFunctionalModule,
		IsZilliaxCosmeticModule:   record.IsZilliaxCosmeticModule,
		Duels: domain.Duels{
			Relevant:    record.DualsRelevant,
			Constructed: record.DualsConstructed,
		},
	}
}

func toInt64Array(values []int) pq.Int64Array {
	result := make(pq.Int64Array, 0, len(values))
	for _, value := range values {
		result = append(result, int64(value))
	}
	return result
}

func fromInt64Array(values pq.Int64Array) []int {
	result := make([]int, 0, len(values))
	for _, value := range values {
		result = append(result, int(value))
	}
	return result
}
