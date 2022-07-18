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
	Doctype   string             `json:"doctype"`
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
	Doctype   string `json:"doctype"`
}

type Contract struct {
	Name     string `json:"name" bson:"name" binding:"required"`
	Email    string `json:"email" bson:"email" binding:"required"`
	Birthday string `json:"birthday" bson:"birthday" binding:"required"`
	Phone    string `json:"phone" bson:"phone" binding:"required"`
	Address  string `json:"address" bson:"address" binding:"required"`
	Snils    string `json:"snils" bson:"snils" binding:"required"`

	Date           string `json:"date" bson:"date" binding:"required"`
	ContractNumber string `json:"contractNumber" bson:"contractNumber" binding:"required"`

	PassportSerial    string `json:"passport_serial" bson:"passport_serial" binding:"required"`
	PassportNumber    string `json:"passport_number" bson:"passport_number" binding:"required"`
	PassportIssueDate string `json:"passport_issue_date" bson:"passportIssueDate" binding:"required"`
	PassportIssuedBy  string `json:"passport_issued_by" bson:"passport_issued_by" binding:"required"`

	Filepath string `json:"filepath,omitempty" bson:"filepath"`
	Doctype  string `json:"doctype"`

	ChildName  string `json:"childName,omitempty" bson:"childName"`
	ChildPhone string `json:"childPhone,omitempty" bson:"childPhone"`
	ChildEmail string `json:"childEmail,omitempty" bson:"childEmail"`
	Education  string `json:"education,omitempty" bson:"education"`

	Course string `json:"course" bson:"course"`
}
