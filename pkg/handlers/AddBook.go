package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"

	"github.com/Mudassir-Munir/rest-api/pkg/mocks"
	"github.com/Mudassir-Munir/rest-api/pkg/modles"
)

func AddBook(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Println(err)
	}

	var book modles.Book
	json.Unmarshal(body, &book)

	book.Id = rand.Intn(100)
	mocks.Books = append(mocks.Books, book)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(book)
}
