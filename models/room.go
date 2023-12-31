package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Room struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	RoomNumber   int                `bson:"roomNumber,omitempty"`
	Type         string             `bson:"type,omitempty"`
	Price        int                `bson:"price,omitempty"`
	Availability bool               `bson:"availability,omitempty" default:"false"`
}
