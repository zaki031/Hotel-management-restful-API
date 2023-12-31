package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Firstname   string             `bson:"firstname,omitempty"`
	Lastname    string             `bson:"lastname,omitempty"`
	Email       string             `bson:"email,omitempty"`
	Password    string             `bson:"password,omitempty"`
	PhoneNumber string             `bson:"phonenumber,omitempty"`
	Role        string             `bson:"role,omitempty" default:"user"`
}
