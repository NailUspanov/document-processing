package services

import (
	"sstu/internal/models"
	"sstu/internal/repos"
)

type Course struct {
	repo repos.Course
}

func NewCourse(repo repos.Course) *Course {
	return &Course{repo: repo}
}

func (d *Course) Create(course models.CourseRequest) (interface{}, error) {
	return d.repo.Create(course)
}

func (d *Course) GetAll() (interface{}, error) {
	return d.repo.FindAll()
}
func (d *Course) GetByName(name string) (models.CourseRequest, error) {
	return d.repo.FindByName(name)
}
func (d *Course) IsPresent(name string) (bool, error) {
	return d.repo.ExistsByName(name)
}
