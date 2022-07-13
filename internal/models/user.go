package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id       primitive.ObjectID `json:"id" bson:"_id" binding:"required"`
	Name     string             `json:"name" binding:"required"`
	Username string             `json:"username" binding:"required"`
	Password string             `json:"password" binding:"required"`
}

type UserRequest struct {
	Name     string `json:"name" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
