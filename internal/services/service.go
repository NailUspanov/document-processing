package services

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"sstu/internal/models"
	"sstu/internal/repos"
)

type DocumentService interface {
	Create(request models.DocumentRequest, course models.CourseRequest) (interface{}, error)
	CreateContract(request models.Contract, course models.CourseRequest) (interface{}, error)
	GetAll() (interface{}, error)
	GetAllChildContracts() (interface{}, error)
	GetAllAdultContracts() (interface{}, error)
	GetAllClients() (interface{}, error)
	CreateChildContract(document models.Contract, course models.CourseRequest) (interface{}, error)
}

type CourseService interface {
	Create(request models.CourseRequest) (interface{}, error)
	GetAll() (interface{}, error)
	GetByName(string) (models.CourseRequest, error)
	IsPresent(string) (bool, error)
}

type AuthorizationService interface {
	CreateUser(user models.UserRequest) (any, error)
	GenerateJWT(username, password string) (string, error)
	ParseToken(s string) (primitive.ObjectID, error)
}

type Service struct {
	DocumentService
	CourseService
	AuthorizationService
}

func NewService(repos *repos.Repository) *Service {
	return &Service{
		DocumentService:      NewDocument(repos.Document),
		CourseService:        NewCourse(repos.Course),
		AuthorizationService: NewAuthService(repos.Authorization),
	}
}
