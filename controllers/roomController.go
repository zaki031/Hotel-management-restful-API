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
	"strings"
)

func CreateRoom(w http.ResponseWriter, r *http.Request) {
	resp := Response{
		Message: "Room created successfully",
	}

	w.Header().Set("Content-Type", "application/json")

	collection := database.Connect("rooms")
	ctx := database.Ctx

	var room models.Room
	fmt.Println("Received request body")
	count, _ := collection.CountDocuments(ctx, bson.D{})
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&room); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	loweredCaseType := strings.ToLower(room.Type)
	if loweredCaseType == "basic" {
		room.Price = 100
	} else if loweredCaseType == "suit" {
		room.Price = 200
	} else if loweredCaseType == "luxury" {
		room.Price = 600
	}

	_, err := collection.InsertOne(ctx, bson.D{
		{Key: "roomNumber", Value: count + 1},
		{Key: "type", Value: room.Type},
		{Key: "availability", Value: "Not Booked"},
		{Key: "price", Value: room.Price},
	})
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		log.Println(err)
		return
	}

	json.NewEncoder(w).Encode(resp)
}

func DeleteRoom(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	objID, err := primitive.ObjectIDFromHex(vars["id"])
	if err != nil {
		log.Fatal(err)
	}
	w.Header().Set("Content-Type", "application/json")
	collection := database.Connect("rooms")
	ctx := database.Ctx
	_, err = collection.DeleteOne(ctx, bson.D{
		{Key: "_id", Value: objID},
	})
	if err != nil {
		log.Fatal(err)
	}
}

func GetAllRooms(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	collection := database.Connect("rooms")
	ctx := database.Ctx
	cursor, err := collection.Find(ctx, bson.D{})
	if err != nil {
		panic(err)
	}
	defer cursor.Close(ctx)

	var rooms []models.Room
	if err = cursor.All(ctx, &rooms); err != nil {
		panic(err)
	}

	json.NewEncoder(w).Encode(rooms)

}




func UpdateRoom(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*") 


	vars := mux.Vars(r)
	objID, err := primitive.ObjectIDFromHex(vars["id"])

	var room models.Room;


	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&room); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	updateFields := bson.M{}

	if room.RoomNumber != 0 {
    updateFields["roomNumber"] = room.RoomNumber
}
	if room.Type != "" {
    updateFields["type"] = room.Type
}
	if room.Price != 0 {
    updateFields["price"] = room.Price
}
if room.Availability != "Not Booked" {
    updateFields["availability"] = room.Availability
}

if room.BookerFirstname != "" {
	updateFields["bookerFirstname"] = room.BookerFirstname

}
if room.BookerLastname != "" {
	updateFields["bookerLastname"] = room.BookerLastname

}
if !room.CheckInDate.Time().IsZero() {
	updateFields["checkInDate"] = room.CheckInDate

}
if !room.CheckOutDate.Time().IsZero() {
	updateFields["checkOutDate"] = room.CheckOutDate

}


	if err != nil {
		log.Fatal(err)
	}
	w.Header().Set("Content-Type", "application/json")
	collection := database.Connect("rooms")
	ctx := database.Ctx
	_, err = collection.UpdateOne(ctx,bson.M{"_id": objID},bson.M{
		"$set": updateFields,
	},
)

	if err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(map[string]string{"message": "Room updated successfully"})

}