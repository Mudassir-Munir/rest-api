package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	// "github.com/Mudassir-Munir/rest-api/pkg/mocks"
	"github.com/Mudassir-Munir/rest-api/pkg/modles"
	"github.com/gorilla/mux"
)

func (h handler) DeleteBook(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	//iterate over all books
	// for index, book := range mocks.Books {
	// 	if book.Id == id {
	// 		//Delete book and send respons if id matches with dynamic id
	// 		mocks.Books = append(mocks.Books[:index], mocks.Books[index+1:]...)

	// 	}
	// }
	var book modles.Book
	if result := h.DB.First(&book, id); result.Error != nil {
		fmt.Println(result.Error)
	}

	h.DB.Delete(&book)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Deleted Book")

}
