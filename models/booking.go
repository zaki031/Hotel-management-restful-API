package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Booking struct {
	ID              primitive.ObjectID `bson:"_id,omitempty"`
	BookerFirstname string             `bson:"BookerFirstname,omitempty"`
	BookerLastname  string             `bson:"BookerLastname,omitempty"`
	RoomNumber      int                `bson:"RoomNumber,omitempty"`

	CheckInDate   primitive.DateTime `bson:"checkInDate,omitempty"`
	CheckOutDate  primitive.DateTime `bson:"checkOutDate,omitempty"`
	BookingStatus string             `bson:"status,omitempty" default:"Pending"`
}
