package routers

import (
	"github.com/gorilla/mux"
	"main/controllers"
)

func RoomsHundlers(r *mux.Router) {
	r.HandleFunc("/rooms/", controllers.GetAllRooms).Methods("GET")
	r.HandleFunc("/rooms/", controllers.CreateRoom).Methods("POST")
	// r.HandleFunc("/users/{id}/", controllers.UpdateUser).Methods("PUT")
	r.HandleFunc("/rooms/{id}", controllers.DeleteRoom).Methods("DELETE")
}
