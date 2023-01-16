package repositories

import (
	"context"

	"github.com/williamwinkler/hs-card-service/internal/domain"
	"github.com/williamwinkler/hs-card-service/internal/infrastructure/migrations"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type RarityRepository struct {
	Raritys *mongo.Collection
}

func NewRarityRepository(db *mongo.Database) *RarityRepository {
	Raritys := db.Collection(migrations.CARDS_RARITY_COLLECTION)

	return &RarityRepository{
		Raritys: Raritys,
	}
}

func (c *RarityRepository) InsertMany(rarities []domain.Rarity) error {
	rarityInterfaces := make([]interface{}, len(rarities))
	for i, rarity := range rarities {
		rarityInterfaces[i] = rarity
	}
	_, err := c.Raritys.InsertMany(context.TODO(), rarityInterfaces)
	return err
}

func (c *RarityRepository) DeleteAll() error {
	_, err := c.Raritys.DeleteMany(context.TODO(), bson.M{}, nil)
	return err
}
