package repositories

import (
	"context"

	"github.com/williamwinkler/hs-card-service/internal/domain"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type RarityRepository struct {
	db *gorm.DB
}

func NewRarityRepository(db *gorm.DB) *RarityRepository {
	return &RarityRepository{db: db}
}

func (c *RarityRepository) InsertMany(ctx context.Context, rarities []domain.Rarity) error {
	if len(rarities) == 0 {
		return nil
	}

	rows := make([]rarityRecord, 0, len(rarities))
	for _, item := range rarities {
		rows = append(rows, rarityRecord{
			Slug:         item.Slug,
			ID:           item.ID,
			CraftingCost: toInt64Array(item.CraftingCost),
			DustValue:    toInt64Array(item.DustValue),
			Name:         item.Name,
		})
	}

	return c.db.WithContext(ctx).Clauses(clause.OnConflict{
		Columns: []clause.Column{{Name: "id"}},
		DoUpdates: clause.AssignmentColumns([]string{
			"slug", "crafting_cost", "dust_value", "name",
		}),
	}).Create(&rows).Error
}

func (c *RarityRepository) DeleteAll(ctx context.Context) error {
	return c.db.WithContext(ctx).Exec("DELETE FROM rarities").Error
}

func (c *RarityRepository) FindAll(ctx context.Context) ([]domain.Rarity, error) {
	var rows []rarityRecord
	if err := c.db.WithContext(ctx).Order("name ASC").Find(&rows).Error; err != nil {
		return []domain.Rarity{}, err
	}

	rarities := make([]domain.Rarity, 0, len(rows))
	for _, row := range rows {
		rarities = append(rarities, domain.Rarity{
			Slug:         row.Slug,
			ID:           row.ID,
			CraftingCost: fromInt64Array(row.CraftingCost),
			DustValue:    fromInt64Array(row.DustValue),
			Name:         row.Name,
		})
	}

	return rarities, nil
}
