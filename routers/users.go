package routers

import (
	"github.com/gorilla/mux"
	"main/controllers"
)

func UsersHundlers(r *mux.Router) {
	r.HandleFunc("/users/", controllers.GetAllUsers).Methods("GET")
	r.HandleFunc("/users/", controllers.CreateUser).Methods("POST")
	r.HandleFunc("/users/{id}", controllers.UpdateUser).Methods("PUT")
	r.HandleFunc("/users/{id}", controllers.DeleteUser).Methods("DELETE")
}
