package repositories

import (
	"context"

	"github.com/williamwinkler/hs-card-service/internal/domain"
	"github.com/williamwinkler/hs-card-service/internal/infrastructure/migrations"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type RarityRepository struct {
	rarities *mongo.Collection
}

func NewRarityRepository(db *mongo.Database) *RarityRepository {
	Raritys := db.Collection(migrations.CARDS_RARITY_COLLECTION)

	return &RarityRepository{
		rarities: Raritys,
	}
}

func (c *RarityRepository) InsertMany(rarities []domain.Rarity) error {
	rarityInterfaces := make([]interface{}, len(rarities))
	for i, rarity := range rarities {
		rarityInterfaces[i] = rarity
	}
	_, err := c.rarities.InsertMany(context.TODO(), rarityInterfaces)
	return err
}

func (c *RarityRepository) DeleteAll() error {
	_, err := c.rarities.DeleteMany(context.TODO(), bson.M{}, nil)
	return err
}

func (c *RarityRepository) FindAll() ([]domain.Rarity, error) {
	cursor, err := c.rarities.Find(context.TODO(), bson.M{})
	if err != nil {
		return []domain.Rarity{}, err
	}

	return decodeToRarities(cursor)
}

func decodeToRarities(cursor *mongo.Cursor) ([]domain.Rarity, error) {
	var rarities []domain.Rarity
	for cursor.Next(context.TODO()) {
		var rarity domain.Rarity
		err := cursor.Decode(&rarity)
		if err != nil {
			return []domain.Rarity{}, err
		}
		rarities = append(rarities, rarity)
	}

	return rarities, nil
}
