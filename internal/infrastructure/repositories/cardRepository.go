package repositories

import (
	"context"
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

func (c *CardRepository) FindRichWithFilter(filter bson.M, page int, limit int) ([]domain.RichCard, error) {
	pipeline := []bson.M{
		{
			"$match": filter,
		},
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
		{
			"$lookup": bson.M{
				"from":         "types",
				"localField":   "cardtypeid",
				"foreignField": "id",
				"as":           "type",
			},
		},
		{
			"$unwind": "$keywordids",
		},
		{
			"$lookup": bson.M{
				"from":         "keywords",
				"localField":   "keywordids",
				"foreignField": "id",
				"as":           "keywords",
			},
		},
		{
			"$match": bson.M{"keywords": bson.M{"$ne": []interface{}{}}},
		},
		{
			"$unwind": "$keywords",
		},
		{
			"$group": bson.M{
				"_id":          "$_id",
				"artistname":   bson.M{"$first": "$artistname"},
				"attack":       bson.M{"$first": "$attack"},
				"set":          bson.M{"$first": "$set"},
				"type":         bson.M{"$first": "$type"},
				"class":        bson.M{"$first": "$class"},
				"collectiable": bson.M{"$first": "$collectiable"},
				"duals":        bson.M{"$first": "$duals"},
				"flavortext":   bson.M{"$first": "$flavortext"},
				"health":       bson.M{"$first": "$health"},
				"id":           bson.M{"$first": "$id"},
				"image":        bson.M{"$first": "$image"},
				"imagegold":    bson.M{"$first": "$imagegold"},
				"keywords":     bson.M{"$push": "$keywords.name"},
				"manacost":     bson.M{"$first": "$manacost"},
				"name":         bson.M{"$first": "$name"},
				"parentid":     bson.M{"$first": "$parentid"},
				"rarity":       bson.M{"$first": "$rarity"},
				"text":         bson.M{"$first": "$text"},
			},
		},
		{
			"$sort": bson.M{"manacost": 1},
		},
		{
			"$skip": int64(limit * (page - 1)),
		},
		{
			"$limit": int64(limit),
		},
		{
			"$project": bson.M{
				"artistname": 1,
				"attack":     1,
				"cardType": bson.M{
					"$arrayElemAt": []interface{}{"$type.name", 0},
				},
				"cardSet": bson.M{
					"$arrayElemAt": []interface{}{"$set.name", 0},
				},
				"class": bson.M{
					"$arrayElemAt": []interface{}{"$class.name", 0},
				},
				"collectible": 1,
				"duals":       1,
				"flavortext":  1,
				"health":      1,
				"id":          1,
				"image":       1,
				"imagegold":   1,
				"keywords":    1,
				"manacost":    1,
				"name":        1,
				"parentid":    1,
				"rarity": bson.M{
					"$arrayElemAt": []interface{}{"$rarity.name", 0},
				},
				"text": 1,
			},
		},
	}

	cursor, err := c.cards.Aggregate(context.TODO(), pipeline)
	if err != nil {
		log.Fatalf("error in aggregate: %v", err)
	}

	var richCards []domain.RichCard
	for cursor.Next(context.TODO()) {
		var richCard domain.RichCard
		err := cursor.Decode(&richCard)
		if err != nil {
			return []domain.RichCard{}, err
		}
		richCards = append(richCards, richCard)
	}

	return richCards, nil
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
	_, err := c.cards.DeleteMany(context.Background(), bson.M{}, nil)
	return err
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
