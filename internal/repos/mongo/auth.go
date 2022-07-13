package mongo

import (
	"context"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"sstu/internal/models"
	"time"
)

type AuthPostgres struct {
	collection *mongo.Collection
}

func NewAuthPostgres(client *mongo.Client) *AuthPostgres {
	return &AuthPostgres{collection: GetCollection(client, "users")}
}

func (a *AuthPostgres) CreateUser(user models.UserRequest) (any, error) {
	tx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	result, err := a.collection.InsertOne(tx, user)
	if err != nil {
		logrus.Panicf("failed to insert into database: %s", err)
		return 0, err
	}
	return result.InsertedID, nil
}

func (a *AuthPostgres) GetUser(username, password string) (models.User, error) {
	tx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var user models.User

	err := a.collection.FindOne(tx, bson.D{{"username", username}, {"password", password}}).Decode(&user)
	if err != nil {
		logrus.Panicf("database find operation failed: %s", err)
	}

	return user, err
}
