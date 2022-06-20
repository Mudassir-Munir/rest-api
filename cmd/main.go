package main

import (
	"log"
	"net/http"

	"github.com/Mudassir-Munir/rest-api/pkg/db"
	"github.com/Mudassir-Munir/rest-api/pkg/handlers"
	"github.com/gorilla/mux"
)

func main() {

	db := db.Init()
	h := handlers.New(db)

	router := mux.NewRouter()

	router.HandleFunc("/books", h.GetAllBooks).Methods(http.MethodGet)
	router.HandleFunc("/books", h.AddBook).Methods(http.MethodPost)
	router.HandleFunc("/books/{id}", h.GetBookById).Methods(http.MethodGet)
	router.HandleFunc("/books/{id}", h.UpdateBookById).Methods(http.MethodPut)
	router.HandleFunc("/books/{id}", h.DeleteBook).Methods(http.MethodDelete)

	log.Println("API is running")
	http.ListenAndServe(":4000", router)
}
