package mongo

import (
	"context"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"sstu/internal/models"
	"time"
)

type Document struct {
	collection *mongo.Collection
}

func NewDocument(client *mongo.Client) *Document {
	collection := GetCollection(client, "documents")
	return &Document{collection: collection}
}

func (d *Document) Create(document models.DocumentRequest) (interface{}, error) {
	tx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	result, err := d.collection.InsertOne(tx, document)
	if err != nil {
		logrus.Fatalf("failed to insert into database: %s", err)
		return 0, err
	}
	return result, nil
}

func (d *Document) FindAll() (interface{}, error) {
	tx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var documents []models.Document

	cur, err := d.collection.Find(tx, bson.D{})

	defer cur.Close(tx)

	if err != nil {
		return documents, err
	}

	if err = cur.All(tx, &documents); err != nil {
		return documents, err
	}

	if err := cur.Err(); err != nil {
		return documents, err
	}

	if len(documents) == 0 {
		return documents, mongo.ErrNoDocuments
	}

	return documents, nil
}

func (d *Document) FindAllClients() (interface{}, error) {
	tx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var clients []bson.M

	groupState := bson.D{{"$group", bson.D{
		{
			"_id",
			bson.D{{"name", "$name"}, {"passport", "$passport"}}},
		{
			"documentsCount",
			bson.D{{"$sum", 1}}},
	}}}

	cur, err := d.collection.Aggregate(tx, mongo.Pipeline{groupState})

	defer cur.Close(tx)

	if err != nil {
		return clients, err
	}

	if err = cur.All(tx, &clients); err != nil {
		return clients, err
	}

	if err := cur.Err(); err != nil {
		return clients, err
	}

	if len(clients) == 0 {
		return clients, mongo.ErrNoDocuments
	}

	return clients, nil
}
