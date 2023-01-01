package repositories

import (
	"context"
	"fmt"
	"os"

	"github.com/williamwinkler/hs-card-service/domain"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type CardRepository struct {
	cardsCollection *mongo.Collection
}

func NewCardRepository() (*CardRepository, error) {
	connection_string, present := os.LookupEnv("mongodb_connection_string")
	if !present {
		return &CardRepository{}, fmt.Errorf("mongodb_connection_string is not present in .env")
	}
	clientOptions := options.Client().ApplyURI(connection_string)

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return &CardRepository{}, err
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return &CardRepository{}, err
	}

	cardsCollection := client.Database("heartstone").Collection("cards")

	return &CardRepository{
		cardsCollection: cardsCollection,
	}, nil
}

func (c *CardRepository) Insert(card domain.Card) error {
	_, err := c.cardsCollection.InsertOne(context.TODO(), card)
	return err
}
