package services

import (
	"sstu/internal/models"
	"sstu/internal/repos"
)

type DocumentService interface {
	Create(request models.DocumentRequest) (interface{}, error)
}

type Service struct {
	DocumentService
}

func NewService(repos *repos.Repository) *Service {
	return &Service{
		DocumentService: NewDocument(repos.Document),
	}
}
