package repositories

import (
	"github.com/williamwinkler/hs-card-service/internal/domain"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type TypeRepository struct {
	db *gorm.DB
}

func NewTypeRepository(db *gorm.DB) *TypeRepository {
	return &TypeRepository{db: db}
}

func (c *TypeRepository) InsertMany(types []domain.Type) error {
	if len(types) == 0 {
		return nil
	}

	rows := make([]typeRecord, 0, len(types))
	for _, item := range types {
		rows = append(rows, typeRecord{
			Slug:      item.Slug,
			ID:        item.ID,
			Name:      item.Name,
			GameModes: toInt64Array(item.GameModes),
		})
	}

	return c.db.Clauses(clause.OnConflict{
		Columns: []clause.Column{{Name: "id"}},
		DoUpdates: clause.AssignmentColumns([]string{
			"slug", "name", "game_modes",
		}),
	}).Create(&rows).Error
}

func (c *TypeRepository) DeleteAll() error {
	return c.db.Exec("DELETE FROM types").Error
}

func (c *TypeRepository) FindAll() ([]domain.Type, error) {
	var rows []typeRecord
	if err := c.db.Order("name ASC").Find(&rows).Error; err != nil {
		return []domain.Type{}, err
	}

	types := make([]domain.Type, 0, len(rows))
	for _, row := range rows {
		types = append(types, domain.Type{
			Slug:      row.Slug,
			ID:        row.ID,
			Name:      row.Name,
			GameModes: fromInt64Array(row.GameModes),
		})
	}

	return types, nil
}
