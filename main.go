package main

import (
	"fmt"
	"log"
	"main/database"
	"main/routers"
	"net/http"
)

func main() {
	ctx := database.Ctx
	defer database.Client.Disconnect(ctx)

	r := routers.Routers()
	log.Fatal(http.ListenAndServe(":8000", r))
	fmt.Println("Listening at 8000")

}
