package repositories

import (
	"context"

	"github.com/williamwinkler/hs-card-service/internal/domain"
	"github.com/williamwinkler/hs-card-service/internal/infrastructure/migrations"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type TypeRepository struct {
	types *mongo.Collection
}

func NewTypeRepository(db *mongo.Database) *TypeRepository {
	Types := db.Collection(migrations.CARDS_TYPES_COLLECTION)

	return &TypeRepository{
		types: Types,
	}
}

func (c *TypeRepository) InsertMany(types []domain.Type) error {
	typeInterfaces := make([]interface{}, len(types))
	for i, Type := range types {
		typeInterfaces[i] = Type
	}
	_, err := c.types.InsertMany(context.TODO(), typeInterfaces)
	return err
}

func (c *TypeRepository) DeleteAll() error {
	_, err := c.types.DeleteMany(context.TODO(), bson.M{}, nil)
	return err
}
