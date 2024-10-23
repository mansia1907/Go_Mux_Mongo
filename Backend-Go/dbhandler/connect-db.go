package dbhandler

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ConnectDB fun to configures the client to use the correct URI and check for errors.
func Database_handler() *mongo.Client {
	client, err := mongo.NewClient(options.Client().ApplyURI(MongoURI()))
	if err != nil {
		log.Fatal(err)
	}

	//timeout of 10 seconds
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	//check if there is an error while connecting to the database and cancel the connection if the connecting period exceeds 10 seconds

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	//ping the database
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB")
	return client
}

// Client instance
// DB is an instance of connectDB
var DB *mongo.Client = Database_handler()

//getting database collections
// to retrive and create collection on database

func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	collection := client.Database("user_api").Collection(collectionName)
	return collection
}

func MongoURI() string {
	err := godotenv.Load() // Load environment variables from .env file
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return os.Getenv("MONGOURI") // Return the MONGOURI variable
}
