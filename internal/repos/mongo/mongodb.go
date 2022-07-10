package mongo

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

func ConnectDB() (*mongo.Client, error) {
	//TODO
	opts := options.Client().ApplyURI("mongodb://localhost:27017")
	opts.SetAuth(options.Credential{
		Username: "admin",
		Password: "password",
	})

	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		logrus.Fatal(err)
		return nil, err
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	//ping the database
	err = client.Ping(ctx, nil)
	if err != nil {
		logrus.Fatal(err)
		return nil, err
	}

	fmt.Println("Connected to MongoDB")
	return client, nil
}

// GetCollection getting database collections
func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	collection := client.Database("sstudb").Collection(collectionName)
	return collection
}
