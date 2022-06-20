package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Mudassir-Munir/rest-api/pkg/modles"
)

func (h handler) GetAllBooks(w http.ResponseWriter, r *http.Request) {

	var book []modles.Book
	if result := h.DB.Find(&book); result.Error != nil {
		fmt.Println(result.Error)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(book)
}
