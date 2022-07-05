package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Document struct {
	Id       primitive.ObjectID `json:"id"`
	Name     string             `json:"name" bson:"name"`
	Filepath string             `json:"filepath" bson:"filepath"`
}
type DocumentRequest struct {
	Name      string `json:"name" bson:"name"`
	Job       string `json:"job" bson:"job"`
	Education string `json:"education" bson:"education"`
	Passport  string `json:"passport" bson:"passport"`
	Address   string `json:"address" bson:"address"`
	Snils     string `json:"snils" bson:"snils"`
	Email     string `json:"email" bson:"email"`
	Filepath  string `json:"filepath,omitempty" bson:"filepath"`
	Street    string `json:"street" bson:"street"`
	House     string `json:"house" bson:"house"`
	Flat      string `json:"flat" bson:"flat"`
	City      string `json:"city" bson:"city"`
	Index     string `json:"index" bson:"index"`
	Phone     string `json:"phone" bson:"phone"`
}
