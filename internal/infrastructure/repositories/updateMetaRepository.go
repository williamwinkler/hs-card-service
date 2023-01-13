package repositories

import (
	"context"

	"github.com/williamwinkler/hs-card-service/internal/domain"
	"github.com/williamwinkler/hs-card-service/internal/infrastructure/migrations"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UpdateMetaRepository struct {
	updateMeta *mongo.Collection
}

func NewUpdateMetaRepository(db *mongo.Database) *UpdateMetaRepository {
	metaCollection := db.Collection(migrations.CARDS_UPDATE_META_COLLECTION)

	return &UpdateMetaRepository{
		updateMeta: metaCollection,
	}
}

func (c *UpdateMetaRepository) InsertOne(cardMeta domain.CardMeta) error {
	_, err := c.updateMeta.InsertOne(context.Background(), cardMeta)
	return err
}

// TODO: query does not work
func (c *UpdateMetaRepository) FindNewest() (domain.CardMeta, error) {
	pipeline := []bson.M{{"$sort": bson.M{"updated": -1, "limit": 1}}}

	result, err := c.updateMeta.Aggregate(context.TODO(), pipeline)
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
