package routers

import (
	"github.com/gorilla/mux"
	"main/controllers"
)

func BookingsHundlers(r *mux.Router) {
	r.HandleFunc("/bookings/", controllers.GetAllBookings).Methods("GET")
	r.HandleFunc("/bookings/", controllers.CreateBooking).Methods("POST")
	// r.HandleFunc("/bookings/{id}/", controllers.UpdateBooking).Methods("PUT")
	r.HandleFunc("/bookings/{id}", controllers.DeleteBooking).Methods("DELETE")
}
