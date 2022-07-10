package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Document struct {
	Id        primitive.ObjectID `json:"id" bson:"_id" binding:"required"`
	Name      string             `json:"name" bson:"name" binding:"required"`
	Job       string             `json:"job" bson:"job" binding:"required"`
	Education string             `json:"education" bson:"education" binding:"required"`
	Passport  string             `json:"passport" bson:"passport" binding:"required"`
	Address   string             `json:"address" bson:"address" binding:"required"`
	Snils     string             `json:"snils" bson:"snils" binding:"required"`
	Email     string             `json:"email" bson:"email" binding:"required"`
	Filepath  string             `json:"filepath,omitempty" bson:"filepath"`
	Street    string             `json:"street" bson:"street" binding:"required"`
	House     string             `json:"house" bson:"house" binding:"required"`
	Flat      string             `json:"flat" bson:"flat"`
	City      string             `json:"city" bson:"city" binding:"required"`
	Index     string             `json:"index" bson:"index" binding:"required"`
	Phone     string             `json:"phone" bson:"phone" binding:"required"`
	Course    string             `json:"course" bson:"course"`
}
type DocumentRequest struct {
	Name      string `json:"name" bson:"name" binding:"required"`
	Job       string `json:"job" bson:"job" binding:"required"`
	Education string `json:"education" bson:"education" binding:"required"`
	Passport  string `json:"passport" bson:"passport" binding:"required"`
	Address   string `json:"address" bson:"address" binding:"required"`
	Snils     string `json:"snils" bson:"snils" binding:"required"`
	Email     string `json:"email" bson:"email" binding:"required"`
	Filepath  string `json:"filepath,omitempty" bson:"filepath"`
	Street    string `json:"street" bson:"street" binding:"required"`
	House     string `json:"house" bson:"house" binding:"required"`
	Flat      string `json:"flat" bson:"flat"`
	City      string `json:"city" bson:"city" binding:"required"`
	Index     string `json:"index" bson:"index" binding:"required"`
	Phone     string `json:"phone" bson:"phone" binding:"required"`
	Course    string `json:"course" bson:"course"`
}
