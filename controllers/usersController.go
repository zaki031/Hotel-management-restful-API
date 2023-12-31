package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"main/database"
	"main/models"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Response struct {
	Message string `json:"message"`
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	resp := Response{
		Message: "User created successfully",
	}

	w.Header().Set("Content-Type", "application/json")

	collection := database.Connect("users")
	ctx := database.Ctx

	var user models.User
	fmt.Println("Received request body")

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	if user.Role == "" {
		user.Role = "user"
	}

	_, err := collection.InsertOne(ctx, bson.D{
		{Key: "firstname", Value: user.Firstname},
		{Key: "lastname", Value: user.Lastname},
		{Key: "email", Value: user.Email},
		{Key: "phonenumber", Value: user.PhoneNumber},
		{Key: "role", Value: user.Role},
	})
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		log.Println(err)
		return
	}

	json.NewEncoder(w).Encode(resp)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	objID, err := primitive.ObjectIDFromHex(vars["id"])
	if err != nil {
		log.Fatal(err)
	}
	w.Header().Set("Content-Type", "application/json")
	collection := database.Connect("users")
	ctx := database.Ctx
	_, err = collection.DeleteOne(ctx, bson.D{
		{Key: "_id", Value: objID},
	})
	if err != nil {
		log.Fatal(err)
	}
}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	collection := database.Connect("users")
	ctx := database.Ctx
	cursor, err := collection.Find(ctx, bson.D{})
	if err != nil {
		panic(err)
	}
	defer cursor.Close(ctx)

	var users []models.User
	if err = cursor.All(ctx, &users); err != nil {
		panic(err)
	}

	json.NewEncoder(w).Encode(users)

}
