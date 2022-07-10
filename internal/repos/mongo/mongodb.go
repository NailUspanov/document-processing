package mongo

import (
	"context"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mongodb"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

func ConnectDB() (*mongo.Client, error) {
	//TODO
	opts := options.Client().ApplyURI("mongodb://sstu_mongo_1:27017")
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

	m, err := migrate.New(
		"file://./migrations",
		"mongodb://admin:password@localhost:27017/sstudb?authSource=admin&readPreference=primary&ssl=false")

	err = m.Up()

	if err != nil && err != migrate.ErrNoChange {
		logrus.Panicln(err.Error())
	}

	return client, nil
}

// GetCollection getting database collections
func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	collection := client.Database("sstudb").Collection(collectionName)
	return collection
}
