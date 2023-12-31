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
