package routers

import (
	"github.com/gorilla/mux"
)

func Routers() *mux.Router {

	r := mux.NewRouter().StrictSlash(true)

	UsersHundlers(r)
	BookingsHundlers(r)
	RoomsHundlers(r)

	return r
}
