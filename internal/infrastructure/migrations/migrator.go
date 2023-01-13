package migrations

import (
	"context"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const DATABASE = "hs-card"
const CARDS_COLLECTION = "cards"
const CARDS_UPDATE_META_COLLECTION = "update-meta"
const CARDS_SETS_COLLECTION = "sets"
const CARDS_CLASSES_COLLECTION = "classes"
const CARDS_KEYWORDS_COLLECTION = "keywords"

type Database struct {
	Client *mongo.Client
	Db     *mongo.Database
}

// TODO make sure that a db and collections are created
// use the mongoDB Go Driver to do so. No need to care about migrations at this point

func SetupDatabase() (*Database, error) {
	// Connect to client
	client, err := getClient()
	if err != nil {
		return &Database{}, err
	}

	// Apply migrations
	createDatabase(client, DATABASE)
	db := client.Database(DATABASE)

	err = createCollection(db, CARDS_COLLECTION)
	if err != nil {
		return &Database{}, err
	}
	err = createCollection(db, CARDS_UPDATE_META_COLLECTION)
	if err != nil {
		return &Database{}, err
	}
	err = createCollection(db, CARDS_SETS_COLLECTION)
	if err != nil {
		return &Database{}, err
	}
	err = createCollection(db, CARDS_CLASSES_COLLECTION)
	if err != nil {
		return &Database{}, err
	}
	err = createCollection(db, CARDS_KEYWORDS_COLLECTION)
	if err != nil {
		return &Database{}, err
	}

	return &Database{
		Client: client,
		Db:     db,
	}, nil
}

func createDatabase(client *mongo.Client, dbName string) {
	// Check if the database exists
	err := client.Database(dbName).RunCommand(context.TODO(), bson.M{"ping": 1}).Err()
	if err != nil {
		// Create the database
		client.Database(dbName).RunCommand(context.TODO(), bson.M{"create": dbName})
		log.Printf("Database '%s' created successfully", dbName)
	}

	log.Printf("Database '%s' already exists", dbName)
}

func createCollection(db *mongo.Database, collectionName string) error {
	err := db.CreateCollection(context.TODO(), collectionName)
	if err != nil {
		if err.Error() == fmt.Sprintf("(NamespaceExists) Collection %s.%s already exists.", db.Name(), collectionName) { // TODO: better error checking
			log.Printf("Collection '%s' already exists", collectionName)
			return nil
		}
		return fmt.Errorf("Failed to create collection '%s' in database '%s': %v", collectionName, db.Name(), err)
	}

	log.Printf("Collection '%s' was created", collectionName)
	return nil
}

func getClient() (*mongo.Client, error) {
	connection_string, present := os.LookupEnv("mongodb_connection_string")
	if !present {
		return &mongo.Client{}, fmt.Errorf("mongodb_connection_string is not present in .env")
	}
	clientOptions := options.Client().ApplyURI(connection_string)

	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return &mongo.Client{}, err
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		return &mongo.Client{}, err
	}

	return client, nil
}
