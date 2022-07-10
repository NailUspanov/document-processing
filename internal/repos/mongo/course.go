package mongo

import (
	"context"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"sstu/internal/models"
	"time"
)

type Course struct {
	collection *mongo.Collection
}

func NewCourse(client *mongo.Client) *Course {
	return &Course{collection: GetCollection(client, "courses")}
}

func (c *Course) Create(course models.CourseRequest) (interface{}, error) {
	tx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	result, err := c.collection.InsertOne(tx, course)
	if err != nil {
		logrus.Fatalf("failed to insert into database: %s", err)
		return 0, err
	}
	return result, nil
}

func (c *Course) FindAll() (interface{}, error) {
	tx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var courses []*models.CourseRequest

	cur, err := c.collection.Find(tx, bson.D{})
	if err != nil {
		return courses, err
	}

	for cur.Next(tx) {
		var course models.CourseRequest
		err := cur.Decode(&course)
		if err != nil {
			return courses, err
		}

		courses = append(courses, &course)
	}

	if err := cur.Err(); err != nil {
		return courses, err
	}

	// once exhausted, close the cursor
	cur.Close(tx)

	if len(courses) == 0 {
		return courses, mongo.ErrNoDocuments
	}

	return courses, nil
}

func (c *Course) FindByName(name string) (models.CourseRequest, error) {
	tx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var course models.CourseRequest

	err := c.collection.FindOne(tx, bson.D{{"name", name}}).Decode(&course)
	if err != nil {
		return course, err
	}

	return course, nil
}

func (c *Course) ExistsByName(name string) (bool, error) {
	_, err := c.FindByName(name)

	return err == nil, err
}
