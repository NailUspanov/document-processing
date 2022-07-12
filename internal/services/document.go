package services

import (
	"github.com/nguyenthenguyen/docx"
	"github.com/sirupsen/logrus"
	"log"
	"sstu/internal/models"
	"sstu/internal/repos"
	"strconv"
	"time"
)

type Document struct {
	repo repos.Document
}

func NewDocument(repo repos.Document) *Document {
	return &Document{repo: repo}
}

func (d *Document) Create(document models.DocumentRequest, course models.CourseRequest) (interface{}, error) {

	// Read from docx file
	r, err := docx.ReadDocxFile("./data/templates/Zayavlenie_svezhak.docx")
	if err != nil {
		log.Println(err)
		panic(err)
	}
	defer r.Close()

	docx1 := r.Editable()
	docx1.Replace("name", document.Name, -1)
	docx1.Replace("job", document.Job, -1)
	docx1.Replace("education", document.Education, -1)
	docx1.Replace("passport", document.Passport, -1)
	docx1.Replace("address", document.Address, -1)
	docx1.Replace("snils", document.Snils, -1)
	docx1.Replace("email", document.Email, -1)
	docx1.Replace("street", document.Street, -1)
	docx1.Replace("house", document.House, -1)
	docx1.Replace("flat", document.Flat, -1)
	docx1.Replace("city", document.City, -1)
	docx1.Replace("index", document.Index, -1)
	docx1.Replace("phone", document.Phone, -1)
	docx1.Replace("courseName", document.Course, -1)
	docx1.Replace("courseType", course.CourseType, -1)

	str := "document" + strconv.Itoa(time.Now().Nanosecond())

	err = docx1.WriteToFile("./data/documents/" + str + ".docx")
	if err != nil {
		logrus.Fatalln("failed to create new docx")
		return nil, err
	}

	document.Filepath = str
	return d.repo.Create(document)
}

func (d *Document) GetAll() (interface{}, error) {
	return d.repo.FindAll()
}

func (d *Document) GetAllClients() (interface{}, error) {
	return d.repo.FindAllClients()
}
