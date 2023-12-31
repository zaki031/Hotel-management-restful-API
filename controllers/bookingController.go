package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"main/database"
	"main/models"
	"net/http"
)

func CreateBooking(w http.ResponseWriter, r *http.Request) {
	resp := Response{
		Message: "booking created successfully",
	}

	w.Header().Set("Content-Type", "application/json")

	collection := database.Connect("bookings")
	ctx := database.Ctx

	var booking models.Booking
	fmt.Println("Received request body")

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&booking); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	_, err := collection.InsertOne(ctx, bson.D{
		{Key: "BookerFirstname", Value: booking.BookerFirstname},
		{Key: "BookerLastname", Value: booking.BookerLastname},
		{Key: "status", Value: "Pending"},
		{Key: "checkInDate", Value: booking.CheckInDate},
		{Key: "checkOutDate", Value: booking.CheckOutDate},
	})
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		log.Println(err)
		return
	}

	json.NewEncoder(w).Encode(resp)
}

func DeleteBooking(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	objID, err := primitive.ObjectIDFromHex(vars["id"])
	if err != nil {
		log.Fatal(err)
	}
	w.Header().Set("Content-Type", "application/json")
	collection := database.Connect("bookings")
	ctx := database.Ctx
	_, err = collection.DeleteOne(ctx, bson.D{
		{Key: "_id", Value: objID},
	})
	if err != nil {
		log.Fatal(err)
	}
}

func GetAllBookings(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	collection := database.Connect("bookings")
	ctx := database.Ctx
	cursor, err := collection.Find(ctx, bson.D{})
	if err != nil {
		panic(err)
	}
	defer cursor.Close(ctx)

	var bookings []models.Booking
	if err = cursor.All(ctx, &bookings); err != nil {
		panic(err)
	}
	json.NewEncoder(w).Encode(bookings)

}
