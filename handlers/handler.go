package handlers

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/luis501angel/go-course/middleware"
	"github.com/luis501angel/go-course/routers"
	"github.com/rs/cors"
)

func Handlers() {
	router := mux.NewRouter()

	router.HandleFunc("/sign-in", middleware.ValidateDB(routers.SignIn)).Methods("POST")

	PORT := "8080"
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
