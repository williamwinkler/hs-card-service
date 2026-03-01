package repositories

import (
	"sort"
	"strings"

	"github.com/williamwinkler/hs-card-service/internal/domain"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

const (
	cardUpsertBatchSize        = 400
	cardKeywordDeleteBatchSize = 2000
	cardKeywordInsertBatchSize = 2000
)

type CardRepository struct {
	db *gorm.DB
}

func NewCardRepository(db *gorm.DB) *CardRepository {
	return &CardRepository{db: db}
}

func (c *CardRepository) InsertOne(card domain.Card) error {
	return c.InsertMany([]domain.Card{card})
}

func (c *CardRepository) InsertMany(cards []domain.Card) error {
	if len(cards) == 0 {
		return nil
	}

	return c.db.Transaction(func(tx *gorm.DB) error {
		cardRows := make([]cardRecord, 0, len(cards))
		cardIDs := make([]int, 0, len(cards))
		keywordRows := make([]cardKeywordRecord, 0)

		for _, card := range cards {
			row, err := toCardRecord(card)
			if err != nil {
				return err
			}
			cardRows = append(cardRows, row)
			cardIDs = append(cardIDs, card.ID)

			for _, keywordID := range card.KeywordIds {
				keywordRows = append(keywordRows, cardKeywordRecord{
					CardID:    card.ID,
					KeywordID: keywordID,
				})
			}
		}

		if err := tx.Clauses(clause.OnConflict{
			Columns: []clause.Column{{Name: "id"}},
			DoUpdates: clause.AssignmentColumns([]string{
				"collectible", "slug", "classid", "multiclassids", "spellschoolid", "cardtypeid", "cardsetid",
				"rarityid", "artistname", "health", "attack", "manacost", "name", "text", "image", "imagegold",
				"flavortext", "cropimage", "parentid", "copyofcardids", "miniontypeid", "childids", "durability",
				"multitypeids", "armor", "iszilliaxfunctionalmodule", "iszilliaxcosmeticmodule", "duals_relevant", "duals_constructed",
			}),
		}).CreateInBatches(&cardRows, cardUpsertBatchSize).Error; err != nil {
			return err
		}

		for _, cardIDBatch := range splitInts(cardIDs, cardKeywordDeleteBatchSize) {
			if err := tx.Where("card_id IN ?", cardIDBatch).Delete(&cardKeywordRecord{}).Error; err != nil {
				return err
			}
		}

		for _, keywordBatch := range splitCardKeywords(keywordRows, cardKeywordInsertBatchSize) {
			if len(keywordBatch) == 0 {
				continue
			}
			if err := tx.Clauses(clause.OnConflict{DoNothing: true}).CreateInBatches(&keywordBatch, cardKeywordInsertBatchSize).Error; err != nil {
				return err
			}
		}

		return nil
	})
}

func (c *CardRepository) FindAll() ([]domain.Card, error) {
	var rows []cardRecord
	if err := c.db.Order("name ASC").Find(&rows).Error; err != nil {
		return []domain.Card{}, err
	}
	return c.toDomainCards(rows)
}

func (c *CardRepository) FindWithFilter(filter domain.CardFilter, page int, limit int) ([]domain.Card, error) {
	var rows []cardRecord
	query := c.applyFilter(c.db.Model(&cardRecord{}), filter)

	if err := query.Order("manacost ASC, name ASC").
		Offset((page - 1) * limit).
		Limit(limit).
		Find(&rows).Error; err != nil {
		return []domain.Card{}, err
	}

	return c.toDomainCards(rows)
}

func (c *CardRepository) FindRichWithFilter(filter domain.CardFilter, page int, limit int) ([]domain.RichCard, error) {
	cards, err := c.FindWithFilter(filter, page, limit)
	if err != nil {
		return []domain.RichCard{}, err
	}
	if len(cards) == 0 {
		return []domain.RichCard{}, nil
	}

	classNames, err := c.fetchClassNames()
	if err != nil {
		return []domain.RichCard{}, err
	}
	setNames, err := c.fetchSetNames()
	if err != nil {
		return []domain.RichCard{}, err
	}
	rarityNames, err := c.fetchRarityNames()
	if err != nil {
		return []domain.RichCard{}, err
	}
	typeNames, err := c.fetchTypeNames()
	if err != nil {
		return []domain.RichCard{}, err
	}
	keywordNames, err := c.fetchKeywordNames()
	if err != nil {
		return []domain.RichCard{}, err
	}

	richCards := make([]domain.RichCard, 0, len(cards))
	for _, card := range cards {
		keywords := make([]string, 0, len(card.KeywordIds))
		for _, keywordID := range card.KeywordIds {
			if name, ok := keywordNames[keywordID]; ok {
				keywords = append(keywords, name)
			}
		}
		sort.Strings(keywords)

		richCards = append(richCards, domain.RichCard{
			ID:            card.ID,
			Collectible:   card.Collectible,
			Slug:          card.Slug,
			Class:         classNames[card.ClassID],
			MultiClassIds: card.MultiClassIds,
			CardType:      typeNames[card.CardTypeID],
			CardSet:       setNames[card.CardSetID],
			Rarity:        rarityNames[card.RarityID],
			ArtistName:    card.ArtistName,
			Health:        card.Health,
			Attack:        card.Attack,
			ManaCost:      card.ManaCost,
			Name:          card.Name,
			Text:          card.Text,
			Image:         card.Image,
			ImageGold:     card.ImageGold,
			FlavorText:    card.FlavorText,
			CropImage:     card.CropImage,
			ParentID:      card.ParentID,
			Keywords:      keywords,
			Duels:         card.Duels,
		})
	}

	return richCards, nil
}

func (c *CardRepository) UpdateOne(card domain.Card) error {
	return c.InsertOne(card)
}

func (c *CardRepository) DeleteOne(card domain.Card) error {
	return c.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("card_id = ?", card.ID).Delete(&cardKeywordRecord{}).Error; err != nil {
			return err
		}
		return tx.Where("id = ?", card.ID).Delete(&cardRecord{}).Error
	})
}

func (c *CardRepository) DeleteAll() error {
	return c.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Exec("DELETE FROM card_keywords").Error; err != nil {
			return err
		}
		return tx.Exec("DELETE FROM cards").Error
	})
}

func (c *CardRepository) Count() (int64, error) {
	return c.CountWithFilter(domain.CardFilter{})
}

func (c *CardRepository) CountWithFilter(filter domain.CardFilter) (int64, error) {
	var count int64
	err := c.applyFilter(c.db.Model(&cardRecord{}), filter).Count(&count).Error
	return count, err
}

func (c *CardRepository) toDomainCards(rows []cardRecord) ([]domain.Card, error) {
	if len(rows) == 0 {
		return []domain.Card{}, nil
	}

	cardIDs := make([]int, 0, len(rows))
	for _, row := range rows {
		cardIDs = append(cardIDs, row.ID)
	}

	keywordIDsMap, err := c.fetchKeywordIDsByCard(cardIDs)
	if err != nil {
		return []domain.Card{}, err
	}

	cards := make([]domain.Card, 0, len(rows))
	for _, row := range rows {
		cards = append(cards, toDomainCard(row, keywordIDsMap[row.ID]))
	}

	return cards, nil
}

func (c *CardRepository) fetchKeywordIDsByCard(cardIDs []int) (map[int][]int, error) {
	var mappings []cardKeywordRecord
	if err := c.db.Where("card_id IN ?", cardIDs).Order("keyword_id ASC").Find(&mappings).Error; err != nil {
		return nil, err
	}

	result := make(map[int][]int, len(cardIDs))
	for _, id := range cardIDs {
		result[id] = []int{}
	}
	for _, mapping := range mappings {
		result[mapping.CardID] = append(result[mapping.CardID], mapping.KeywordID)
	}

	return result, nil
}

func (c *CardRepository) fetchClassNames() (map[int]string, error) {
	var rows []classRecord
	if err := c.db.Find(&rows).Error; err != nil {
		return nil, err
	}
	result := make(map[int]string, len(rows))
	for _, row := range rows {
		result[row.ID] = row.Name
	}
	return result, nil
}

func (c *CardRepository) fetchSetNames() (map[int]string, error) {
	var rows []setRecord
	if err := c.db.Find(&rows).Error; err != nil {
		return nil, err
	}
	result := make(map[int]string, len(rows))
	for _, row := range rows {
		result[row.ID] = row.Name
	}
	return result, nil
}

func (c *CardRepository) fetchRarityNames() (map[int]string, error) {
	var rows []rarityRecord
	if err := c.db.Find(&rows).Error; err != nil {
		return nil, err
	}
	result := make(map[int]string, len(rows))
	for _, row := range rows {
		result[row.ID] = row.Name
	}
	return result, nil
}

func (c *CardRepository) fetchTypeNames() (map[int]string, error) {
	var rows []typeRecord
	if err := c.db.Find(&rows).Error; err != nil {
		return nil, err
	}
	result := make(map[int]string, len(rows))
	for _, row := range rows {
		result[row.ID] = row.Name
	}
	return result, nil
}

func (c *CardRepository) fetchKeywordNames() (map[int]string, error) {
	var rows []keywordRecord
	if err := c.db.Find(&rows).Error; err != nil {
		return nil, err
	}
	result := make(map[int]string, len(rows))
	for _, row := range rows {
		result[row.ID] = row.Name
	}
	return result, nil
}

func (c *CardRepository) applyFilter(query *gorm.DB, filter domain.CardFilter) *gorm.DB {
	if filter.NameContains != nil {
		value := strings.TrimSpace(*filter.NameContains)
		if value != "" {
			// Keep substring matching and add trigram fuzzy matching for typo tolerance.
			query = query.Where("(name ILIKE ? OR name % ?)", "%"+value+"%", value)
		}
	}
	if filter.ManaCost != nil {
		query = query.Where("manacost = ?", *filter.ManaCost)
	}
	if filter.ManaCostGte != nil {
		query = query.Where("manacost >= ?", *filter.ManaCostGte)
	}
	if filter.Health != nil {
		query = query.Where("health = ?", *filter.Health)
	}
	if filter.HealthGte != nil {
		query = query.Where("health >= ?", *filter.HealthGte)
	}
	if filter.Attack != nil {
		query = query.Where("attack = ?", *filter.Attack)
	}
	if filter.AttackGte != nil {
		query = query.Where("attack >= ?", *filter.AttackGte)
	}
	if filter.ClassID != nil {
		query = query.Where("classid = ?", *filter.ClassID)
	}
	if filter.RarityID != nil {
		query = query.Where("rarityid = ?", *filter.RarityID)
	}
	if len(filter.TypeIDs) > 0 {
		query = query.Where("cardtypeid IN ?", filter.TypeIDs)
	}
	if filter.SetID != nil {
		query = query.Where("cardsetid = ?", *filter.SetID)
	}
	if len(filter.KeywordIDsAll) > 0 {
		subQuery := c.db.Table("card_keywords").
			Select("card_id").
			Where("keyword_id IN ?", filter.KeywordIDsAll).
			Group("card_id").
			Having("COUNT(DISTINCT keyword_id) = ?", len(filter.KeywordIDsAll))

		query = query.Where("id IN (?)", subQuery)
	}
	return query
}

func splitInts(values []int, batchSize int) [][]int {
	if len(values) == 0 {
		return nil
	}
	if batchSize <= 0 {
		batchSize = len(values)
	}

	var batches [][]int
	for start := 0; start < len(values); start += batchSize {
		end := start + batchSize
		if end > len(values) {
			end = len(values)
		}
		batches = append(batches, values[start:end])
	}
	return batches
}

func splitCardKeywords(values []cardKeywordRecord, batchSize int) [][]cardKeywordRecord {
	if len(values) == 0 {
		return nil
	}
	if batchSize <= 0 {
		batchSize = len(values)
	}

	var batches [][]cardKeywordRecord
	for start := 0; start < len(values); start += batchSize {
		end := start + batchSize
		if end > len(values) {
			end = len(values)
		}
		batches = append(batches, values[start:end])
	}
	return batches
}
