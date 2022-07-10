package mongo

import (
	"context"
	"github.com/sirupsen/logrus"
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
