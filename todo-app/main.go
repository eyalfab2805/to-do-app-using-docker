package main

import (
	"log"
	"net/http"

	"github.com/eyalfab2805/todo-app/database"
	"github.com/eyalfab2805/todo-app/handlers"
	"github.com/gorilla/mux"
)

func main() {
	// Initialize the database
	database.InitDB()

	// Set up the router
	router := mux.NewRouter()

	// Define routes

	router.HandleFunc("/tasks", handlers.CreateTask).Methods("POST")
	router.HandleFunc("/tasks", handlers.GetTasks).Methods("GET")
	router.HandleFunc("/tasks/{id}", handlers.GetTask).Methods("GET")
	router.HandleFunc("/tasks/{id}", handlers.UpdateTask).Methods("PUT")
	router.HandleFunc("/tasks/{id}", handlers.DeleteTask).Methods("DELETE")
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))
	// Start the server
	log.Println("Server is running on port 8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
