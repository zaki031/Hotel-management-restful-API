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
	r.Use(EnableCORS)


}

func EnableCORS(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Access-Control-Allow-Origin", "*")
        w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
        w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

        if r.Method == http.MethodOptions {
            w.WriteHeader(http.StatusOK)
            return
        }

        next.ServeHTTP(w, r)
    })
}