package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Course struct {
	Id       primitive.ObjectID `json:"id"`
	Name     string             `json:"name" bson:"name"`
	Filepath string             `json:"filepath" bson:"filepath"`
}

type CourseRequest struct {
	Id         primitive.ObjectID `json:"id" bson:"_id"`
	Name       string             `json:"name" bson:"name" binding:"required"`
	Hours      int32              `json:"hours" bson:"hours" binding:"required"`
	Cost       int32              `json:"cost" bson:"cost" binding:"required"`
	CourseType string             `json:"courseType" bson:"courseType" binding:"required"`
}
