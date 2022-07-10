package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Course struct {
	Id         primitive.ObjectID `json:"id" bson:"_id"  binding:"required"`
	Name       string             `json:"name" bson:"name" binding:"required"`
	Hours      int32              `json:"hours" bson:"hours" binding:"required"`
	Cost       int32              `json:"cost" bson:"cost" binding:"required"`
	CourseType string             `json:"courseType" bson:"courseType" binding:"required"`
}

type CourseRequest struct {
	Name       string `json:"name" bson:"name" binding:"required"`
	Hours      int32  `json:"hours" bson:"hours" binding:"required"`
	Cost       int32  `json:"cost" bson:"cost" binding:"required"`
	CourseType string `json:"courseType" bson:"courseType" binding:"required"`
}
