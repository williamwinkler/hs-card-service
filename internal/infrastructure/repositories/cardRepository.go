package repositories

import (
	"context"
	"fmt"

	"github.com/williamwinkler/hs-card-service/internal/domain"
	"github.com/williamwinkler/hs-card-service/internal/infrastructure/migrations"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type CardRepository struct {
	cardsCollection *mongo.Collection
}

func NewCardRepository(db *mongo.Database) *CardRepository {
	cardsCollection := db.Collection(migrations.CARDS_COLLECTION)

	return &CardRepository{
		cardsCollection: cardsCollection,
	}
}

func (c *CardRepository) InsertOne(card domain.Card) error {
	_, err := c.cardsCollection.InsertOne(context.Background(), card)
	return err
}

func (c *CardRepository) InsertMany(cards []domain.Card) error {
	cardsInterface := make([]interface{}, len(cards))
	for i, card := range cards {
		cardsInterface[i] = card
	}
	_, err := c.cardsCollection.InsertMany(context.Background(), cardsInterface)
	return err
}

func (c *CardRepository) FindAll() ([]domain.Card, error) {
	sortByNameFilter := options.Find().SetSort(bson.M{"name": 1})
	cursor, err := c.cardsCollection.Find(context.Background(), bson.M{}, sortByNameFilter)
	if err != nil {
		return []domain.Card{}, err
	}

	return decodeToCards(cursor)
}

func (c *CardRepository) FindWithFilter(filter bson.M, page int, limit int) ([]domain.Card, error) {
	options := options.Find().SetSort(bson.M{"manacost": 1}).SetLimit(int64(limit)).SetSkip(int64(limit * (page - 1)))
	cursor, err := c.cardsCollection.Find(context.TODO(), filter, options)
	if err != nil {
		return []domain.Card{}, err
	}

	return decodeToCards(cursor)
}

func (c *CardRepository) UpdateOne(card domain.Card) error {
	cardIdFilter := bson.M{"id": card.ID}
	cardUpdate := bson.M{"$set": card}
	_, err := c.cardsCollection.UpdateOne(context.Background(), cardIdFilter, cardUpdate, nil)
	return err
}

func (c *CardRepository) DeleteOne(domain.Card) error {
	return nil
}

func (c *CardRepository) DeleteAll() error {
	result, err := c.cardsCollection.DeleteMany(context.Background(), bson.M{}, nil)
	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		return fmt.Errorf("zero documents were deleted")
	}

	return nil
}

func (c *CardRepository) Count() (int64, error) {
	return c.CountWithFilter(bson.M{})
}

func (c *CardRepository) CountWithFilter(filter bson.M) (int64, error) {
	return c.cardsCollection.CountDocuments(context.Background(), filter, nil)
}

func decodeToCards(cursor *mongo.Cursor) ([]domain.Card, error) {
	var cards []domain.Card
	for cursor.Next(context.Background()) {
		var card domain.Card
		err := cursor.Decode(&card)
		if err != nil {
			return []domain.Card{}, err
		}
		cards = append(cards, card)
	}
	return cards, nil
}
