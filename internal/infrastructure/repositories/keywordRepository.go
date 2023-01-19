package repositories

import (
	"context"

	"github.com/williamwinkler/hs-card-service/internal/domain"
	"github.com/williamwinkler/hs-card-service/internal/infrastructure/migrations"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type KeywordRepository struct {
	keywords *mongo.Collection
}

func NewKeywordRepository(db *mongo.Database) *KeywordRepository {
	keywords := db.Collection(migrations.CARDS_KEYWORDS_COLLECTION)

	return &KeywordRepository{
		keywords: keywords,
	}
}

func (c *KeywordRepository) InsertMany(Keywords []domain.Keyword) error {
	keywordInterfaces := make([]interface{}, len(Keywords))
	for i, Keyword := range Keywords {
		keywordInterfaces[i] = Keyword
	}
	_, err := c.keywords.InsertMany(context.TODO(), keywordInterfaces)
	return err
}

func (c *KeywordRepository) DeleteAll() error {
	_, err := c.keywords.DeleteMany(context.TODO(), bson.M{}, nil)
	return err
}

func (c *KeywordRepository) FindAll() ([]domain.Keyword, error) {
	cursor, err := c.keywords.Find(context.TODO(), bson.M{})
	if err != nil {
		return []domain.Keyword{}, err
	}

	return decodeToKeywords(cursor)
}

func decodeToKeywords(cursor *mongo.Cursor) ([]domain.Keyword, error) {
	var keywords []domain.Keyword
	for cursor.Next(context.TODO()) {
		var keyword domain.Keyword
		err := cursor.Decode(&keyword)
		if err != nil {
			return []domain.Keyword{}, err
		}
		keywords = append(keywords, keyword)
	}

	return keywords, nil
}
