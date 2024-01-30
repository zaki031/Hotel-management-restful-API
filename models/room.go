package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Room struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	RoomNumber   int                `bson:"roomNumber,omitempty"`
	Type         string             `bson:"type,omitempty"`
	BookerFirstname         string             `bson:"bookerFirstname,omitempty"`
	BookerLastname        string             `bson:"bookerLastname,omitempty"`
	CheckInDate   primitive.DateTime `bson:"checkInDate,omitempty"`
	CheckOutDate  primitive.DateTime `bson:"checkOutDate,omitempty"`
	Price        int                `bson:"price,omitempty"`
	Availability string              `bson:"availability,omitempty" default:"false"`
}
