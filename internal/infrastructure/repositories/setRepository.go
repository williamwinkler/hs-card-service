package repositories

import (
	"github.com/williamwinkler/hs-card-service/internal/domain"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type SetRepository struct {
	db *gorm.DB
}

func NewSetRepository(db *gorm.DB) *SetRepository {
	return &SetRepository{db: db}
}

func (c *SetRepository) InsertMany(sets []domain.Set) error {
	if len(sets) == 0 {
		return nil
	}

	rows := make([]setRecord, 0, len(sets))
	for _, item := range sets {
		rows = append(rows, setRecord{
			ID:                          item.ID,
			Name:                        item.Name,
			Slug:                        item.Slug,
			Type:                        item.Type,
			CollectibleCount:            item.CollectibleCount,
			CollectibleRevealedCount:    item.CollectibleRevealedCount,
			NonCollectibleCount:         item.NonCollectibleCount,
			NonCollectibleRevealedCount: item.NonCollectibleRevealedCount,
			AliasSetIDs:                 toInt64Array(item.AliasSetIds),
		})
	}

	return c.db.Clauses(clause.OnConflict{
		Columns: []clause.Column{{Name: "id"}},
		DoUpdates: clause.AssignmentColumns([]string{
			"name", "slug", "type", "collectible_count", "collectible_revealed_count",
			"non_collectible_count", "non_collectible_revealed_count", "alias_set_ids",
		}),
	}).Create(&rows).Error
}

func (c *SetRepository) DeleteAll() error {
	return c.db.Exec("DELETE FROM sets").Error
}

func (c *SetRepository) FindAll() ([]domain.Set, error) {
	var rows []setRecord
	if err := c.db.Order("name ASC").Find(&rows).Error; err != nil {
		return []domain.Set{}, err
	}

	sets := make([]domain.Set, 0, len(rows))
	for _, row := range rows {
		sets = append(sets, domain.Set{
			ID:                          row.ID,
			Name:                        row.Name,
			Slug:                        row.Slug,
			Type:                        row.Type,
			CollectibleCount:            row.CollectibleCount,
			CollectibleRevealedCount:    row.CollectibleRevealedCount,
			NonCollectibleCount:         row.NonCollectibleCount,
			NonCollectibleRevealedCount: row.NonCollectibleRevealedCount,
			AliasSetIds:                 fromInt64Array(row.AliasSetIDs),
		})
	}

	return sets, nil
}
