package repos

import (
	"go.mongodb.org/mongo-driver/mongo"
	"sstu/internal/models"
	mongo2 "sstu/internal/repos/mongo"
)

type Document interface {
	Create(document any) (interface{}, error)
	FindAll() (interface{}, error)
	FindAllChildContracts() (interface{}, error)
	FindAllAdultContracts() (interface{}, error)
	FindAllClients() (interface{}, error)
	FindAllContractsByClient(string, string, string) (interface{}, error)
}

type Course interface {
	Create(course models.CourseRequest) (interface{}, error)
	FindAll() (interface{}, error)
	FindByName(string) (models.CourseRequest, error)
	ExistsByName(string) (bool, error)
}

type Authorization interface {
	CreateUser(user models.UserRequest) (any, error)
	GetUser(username, password string) (models.User, error)
}

type Repository struct {
	Document
	Course
	Authorization
}

func NewRepository(client *mongo.Client) *Repository {
	return &Repository{
		Document:      mongo2.NewDocument(client),
		Course:        mongo2.NewCourse(client),
		Authorization: mongo2.NewAuthPostgres(client),
	}
}
