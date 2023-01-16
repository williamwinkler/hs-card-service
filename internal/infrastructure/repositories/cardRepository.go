package repositories

import (
	"context"
	"fmt"
	"log"

	"github.com/williamwinkler/hs-card-service/internal/domain"
	"github.com/williamwinkler/hs-card-service/internal/infrastructure/migrations"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type CardRepository struct {
	cards *mongo.Collection
}

func NewCardRepository(db *mongo.Database) *CardRepository {
	cardsCollection := db.Collection(migrations.CARDS_COLLECTION)

	return &CardRepository{
		cards: cardsCollection,
	}
}

func (c *CardRepository) InsertOne(card domain.Card) error {
	_, err := c.cards.InsertOne(context.Background(), card)
	return err
}

func (c *CardRepository) InsertMany(cards []domain.Card) error {
	cardsInterface := make([]interface{}, len(cards))
	for i, card := range cards {
		cardsInterface[i] = card
	}
	_, err := c.cards.InsertMany(context.Background(), cardsInterface)
	return err
}

func (c *CardRepository) FindAll() ([]domain.Card, error) {
	sortByNameFilter := options.Find().SetSort(bson.M{"name": 1})
	cursor, err := c.cards.Find(context.Background(), bson.M{}, sortByNameFilter)
	if err != nil {
		return []domain.Card{}, err
	}

	return decodeToCards(cursor)
}

func (c *CardRepository) FindWithFilter(filter bson.M, page int, limit int) ([]domain.Card, error) {
	options := options.Find().SetSort(bson.M{"manacost": 1}).SetLimit(int64(limit)).SetSkip(int64(limit * (page - 1)))
	cursor, err := c.cards.Find(context.TODO(), filter, options)
	if err != nil {
		return []domain.Card{}, err
	}

	return decodeToCards(cursor)
}

func (c *CardRepository) FindWithAggregate() {
	pipeline := []bson.M{
		{
			"$lookup": bson.M{
				"from":         "sets",
				"localField":   "cardsetid",
				"foreignField": "id",
				"as":           "set",
			},
		},
		{
			"$lookup": bson.M{
				"from":         "classes",
				"localField":   "classid",
				"foreignField": "id",
				"as":           "class",
			},
		},
		{
			"$lookup": bson.M{
				"from":         "rarities",
				"localField":   "rarityid",
				"foreignField": "id",
				"as":           "rarity",
			},
		},
		{"$unwind": "$keywordids"},
		{
			"$lookup": bson.M{
				"from":         "keywords",
				"localField":   "keywordids",
				"foreignField": "id",
				"as":           "keywords_names",
			},
		},
		{"$match": bson.M{"keywords_names": bson.M{"$ne": []interface{}{}}}},
		{
			"$project": bson.M{
				"name":                1,
				"keywords_names.name": 1,
				"set": bson.M{
					"$arrayElemAt": []interface{}{"$set.name", 0},
				},
				"class": bson.M{
					"$arrayElemAt": []interface{}{"$class.name", 0},
				},
				"rarity": bson.M{
					"$arrayElemAt": []interface{}{"$rarity.name", 0},
				},
			},
		},
	}

	cursor, err := c.cards.Aggregate(context.TODO(), pipeline)
	if err != nil {
		log.Fatalf("error in aggregate: %v", err)
	}

	i := 0
	for cursor.Next(context.TODO()) {
		if i > 0 {
			break
		}
		var result bson.M
		err := cursor.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(result)
		i++
	}
}

func (c *CardRepository) UpdateOne(card domain.Card) error {
	cardIdFilter := bson.M{"id": card.ID}
	cardUpdate := bson.M{"$set": card}
	_, err := c.cards.UpdateOne(context.Background(), cardIdFilter, cardUpdate, nil)
	return err
}

func (c *CardRepository) DeleteOne(domain.Card) error {
	return nil
}

func (c *CardRepository) DeleteAll() error {
	result, err := c.cards.DeleteMany(context.Background(), bson.M{}, nil)
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
	return c.cards.CountDocuments(context.Background(), filter, nil)
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
