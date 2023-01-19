package repositories

import (
	"context"

	"github.com/williamwinkler/hs-card-service/internal/domain"
	"github.com/williamwinkler/hs-card-service/internal/infrastructure/migrations"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type SetRepository struct {
	sets *mongo.Collection
}

func NewSetRepository(db *mongo.Database) *SetRepository {
	sets := db.Collection(migrations.CARDS_SETS_COLLECTION)

	return &SetRepository{
		sets: sets,
	}
}

func (c *SetRepository) InsertMany(sets []domain.Set) error {
	setInterfaces := make([]interface{}, len(sets))
	for i, set := range sets {
		setInterfaces[i] = set
	}
	_, err := c.sets.InsertMany(context.TODO(), setInterfaces)
	return err
}

func (c *SetRepository) DeleteAll() error {
	_, err := c.sets.DeleteMany(context.TODO(), bson.M{}, nil)
	return err
}

func (c *SetRepository) FindAll() ([]domain.Set, error) {
	cursor, err := c.sets.Find(context.TODO(), bson.M{})
	if err != nil {
		return []domain.Set{}, err
	}

	return decodeToSets(cursor)
}

func decodeToSets(cursor *mongo.Cursor) ([]domain.Set, error) {
	var sets []domain.Set
	for cursor.Next(context.TODO()) {
		var set domain.Set
		err := cursor.Decode(&set)
		if err != nil {
			return []domain.Set{}, err
		}
		sets = append(sets, set)
	}

	return sets, nil
}
