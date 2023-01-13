package repositories

import (
	"context"

	"github.com/williamwinkler/hs-card-service/internal/domain"
	"github.com/williamwinkler/hs-card-service/internal/infrastructure/migrations"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ClassRepository struct {
	classes *mongo.Collection
}

func NewClassRepository(db *mongo.Database) *ClassRepository {
	classes := db.Collection(migrations.CARDS_CLASSES_COLLECTION)

	return &ClassRepository{
		classes: classes,
	}
}

func (c *ClassRepository) InsertMany(classes []domain.Class) error {
	classInterfaces := make([]interface{}, len(classes))
	for i, set := range classes {
		classInterfaces[i] = set
	}
	_, err := c.classes.InsertMany(context.TODO(), classInterfaces)
	return err
}

func (c *ClassRepository) DeleteAll() error {
	_, err := c.classes.DeleteMany(context.TODO(), bson.M{}, nil)
	return err
}
