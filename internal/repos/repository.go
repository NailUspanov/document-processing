package repos

import (
	"go.mongodb.org/mongo-driver/mongo"
	"sstu/internal/models"
	mongo2 "sstu/internal/repos/mongo"
)

type Document interface {
	Create(document models.DocumentRequest) (interface{}, error)
	FindAll() (interface{}, error)
	FindAllClients() (interface{}, error)
}

type Course interface {
	Create(course models.CourseRequest) (interface{}, error)
	FindAll() (interface{}, error)
	FindByName(string) (models.CourseRequest, error)
	ExistsByName(string) (bool, error)
}

type Repository struct {
	Document
	Course
}

func NewRepository(client *mongo.Client) *Repository {
	return &Repository{
		Document: mongo2.NewDocument(client),
		Course:   mongo2.NewCourse(client),
	}
}
