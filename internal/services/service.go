package services

import (
	"sstu/internal/models"
	"sstu/internal/repos"
)

type DocumentService interface {
	Create(request models.DocumentRequest, course models.CourseRequest) (interface{}, error)
	GetAll() (interface{}, error)
	GetAllClients() (interface{}, error)
}

type CourseService interface {
	Create(request models.CourseRequest) (interface{}, error)
	GetAll() (interface{}, error)
	GetByName(string) (models.CourseRequest, error)
	IsPresent(string) (bool, error)
}

type Service struct {
	DocumentService
	CourseService
}

func NewService(repos *repos.Repository) *Service {
	return &Service{
		DocumentService: NewDocument(repos.Document),
		CourseService:   NewCourse(repos.Course),
	}
}
