package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/intogit/go-blog-backend/handlers"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("This is main begining")

	godotenv.Load()
	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("PORT is not found in environment")
	}
	fmt.Println("PORT is", portString)
	// http.HandleFunc("/healthCheck", handleHealthCheck)
	// http.HandleFunc("/createUser", handleCreateUser)
	// log.Fatal(http.ListenAndServe(":"+ portString, nil))

	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	log.Println("Server started at port", portString)
	err := http.ListenAndServe(":"+portString, handlers.CreateRoute())
	if err != nil {
		log.Fatal(err)
	}
}
