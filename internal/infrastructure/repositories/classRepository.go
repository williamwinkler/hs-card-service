package repositories

import (
	"context"

	"github.com/williamwinkler/hs-card-service/internal/domain"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ClassRepository struct {
	db *gorm.DB
}

func NewClassRepository(db *gorm.DB) *ClassRepository {
	return &ClassRepository{db: db}
}

func (c *ClassRepository) InsertMany(ctx context.Context, classes []domain.Class) error {
	if len(classes) == 0 {
		return nil
	}

	rows := make([]classRecord, 0, len(classes))
	for _, item := range classes {
		rows = append(rows, classRecord{
			Slug:                 item.Slug,
			ID:                   item.ID,
			Name:                 item.Name,
			CardID:               item.CardID,
			HeroPowerCardID:      item.HeroPowerCardID,
			AlternateHeroCardIDs: toInt64Array(item.AlternateHeroCardIds),
		})
	}

	return c.db.WithContext(ctx).Clauses(clause.OnConflict{
		Columns: []clause.Column{{Name: "id"}},
		DoUpdates: clause.AssignmentColumns([]string{
			"slug", "name", "card_id", "hero_power_card_id", "alternate_hero_card_ids",
		}),
	}).Create(&rows).Error
}

func (c *ClassRepository) DeleteAll(ctx context.Context) error {
	return c.db.WithContext(ctx).Exec("DELETE FROM classes").Error
}

func (c *ClassRepository) FindAll(ctx context.Context) ([]domain.Class, error) {
	var rows []classRecord
	if err := c.db.WithContext(ctx).Order("name ASC").Find(&rows).Error; err != nil {
		return []domain.Class{}, err
	}

	classes := make([]domain.Class, 0, len(rows))
	for _, row := range rows {
		classes = append(classes, domain.Class{
			Slug:                 row.Slug,
			ID:                   row.ID,
			Name:                 row.Name,
			CardID:               row.CardID,
			HeroPowerCardID:      row.HeroPowerCardID,
			AlternateHeroCardIds: fromInt64Array(row.AlternateHeroCardIDs),
		})
	}

	return classes, nil
}
