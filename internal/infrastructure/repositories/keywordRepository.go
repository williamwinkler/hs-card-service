package repositories

import (
	"github.com/williamwinkler/hs-card-service/internal/domain"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type KeywordRepository struct {
	db *gorm.DB
}

func NewKeywordRepository(db *gorm.DB) *KeywordRepository {
	return &KeywordRepository{db: db}
}

func (c *KeywordRepository) InsertMany(keywords []domain.Keyword) error {
	if len(keywords) == 0 {
		return nil
	}

	rows := make([]keywordRecord, 0, len(keywords))
	for _, item := range keywords {
		rows = append(rows, keywordRecord{
			ID:        item.ID,
			Slug:      item.Slug,
			Name:      item.Name,
			RefText:   item.RefText,
			Text:      item.Text,
			GameModes: toInt64Array(item.GameModes),
		})
	}

	return c.db.Clauses(clause.OnConflict{
		Columns: []clause.Column{{Name: "id"}},
		DoUpdates: clause.AssignmentColumns([]string{
			"slug", "name", "ref_text", "text", "game_modes",
		}),
	}).Create(&rows).Error
}

func (c *KeywordRepository) DeleteAll() error {
	return c.db.Exec("DELETE FROM keywords").Error
}

func (c *KeywordRepository) FindAll() ([]domain.Keyword, error) {
	var rows []keywordRecord
	if err := c.db.Order("name ASC").Find(&rows).Error; err != nil {
		return []domain.Keyword{}, err
	}

	keywords := make([]domain.Keyword, 0, len(rows))
	for _, row := range rows {
		keywords = append(keywords, domain.Keyword{
			ID:        row.ID,
			Slug:      row.Slug,
			Name:      row.Name,
			RefText:   row.RefText,
			Text:      row.Text,
			GameModes: fromInt64Array(row.GameModes),
		})
	}

	return keywords, nil
}
