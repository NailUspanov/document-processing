package repos

import (
	"go.mongodb.org/mongo-driver/mongo"
	"sstu/internal/models"
	mongo2 "sstu/internal/repos/mongo"
)

type Document interface {
	Create(document models.DocumentRequest) (interface{}, error)
}

type Repository struct {
	Document
}

func NewRepository(collection *mongo.Collection) *Repository {
	return &Repository{
		Document: mongo2.NewDocument(collection),
	}
}
