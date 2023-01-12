package repositories

import (
	"context"

	"github.com/williamwinkler/hs-card-service/internal/domain"
	"github.com/williamwinkler/hs-card-service/internal/infrastructure/migrations"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UpdateMetaRepository struct {
	metaCollection *mongo.Collection
}

func NewUpdateMetaRepository(db *mongo.Database) *UpdateMetaRepository {
	metaCollection := db.Collection(migrations.CARDS_UPDATE_META_COLLECTION)

	return &UpdateMetaRepository{
		metaCollection: metaCollection,
	}
}

func (c *UpdateMetaRepository) InsertOne(cardMeta domain.CardMeta) error {
	_, err := c.metaCollection.InsertOne(context.Background(), cardMeta)
	return err
}

func (c *UpdateMetaRepository) FindNewest() (domain.CardMeta, error) {
	pipeline := []bson.M{{"$sort": bson.M{"updated": -1}}, {"limt": 1}}

	result, err := c.metaCollection.Aggregate(context.TODO(), pipeline)
	if err != nil {
		return domain.CardMeta{}, err
	}
	var cardMeta domain.CardMeta
	err = result.Decode(&cardMeta)
	if err != nil {
		return domain.CardMeta{}, err
	}

	return cardMeta, nil
}
