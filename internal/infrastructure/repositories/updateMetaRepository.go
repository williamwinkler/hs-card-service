package repositories

import (
	"github.com/williamwinkler/hs-card-service/internal/domain"
	"gorm.io/gorm"
)

type UpdateMetaRepository struct {
	db *gorm.DB
}

func NewUpdateMetaRepository(db *gorm.DB) *UpdateMetaRepository {
	return &UpdateMetaRepository{db: db}
}

func (c *UpdateMetaRepository) InsertOne(cardMeta domain.CardMeta) error {
	row := updateMetaRecord{
		UpdatedAt: cardMeta.Updated,
		IsChanged: cardMeta.IsChanged,
	}
	return c.db.Create(&row).Error
}

func (c *UpdateMetaRepository) FindNewest() (domain.CardMeta, error) {
	var row updateMetaRecord
	if err := c.db.Order("updated DESC").Limit(1).Take(&row).Error; err != nil {
		return domain.CardMeta{}, err
	}
	return domain.CardMeta{Updated: row.UpdatedAt, IsChanged: row.IsChanged}, nil
}
