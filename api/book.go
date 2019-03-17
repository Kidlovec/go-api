package api

import (
  "encoding/json"
  "net/http"
)

// Book type with Name, Author and ISBN
type Book struct {
  Title string
  Author string
  ISBN string

  // define the book
}

// ToJSON to be used for marshalling of Book type
func (b Book) ToJSON() [] byte {
  ToJSON, err := json.Marshal(b);
  if err != nil {
    panic(err)
  }
  return ToJSON
}

// FromJSON to be used for unmarshalling of Book type
func FromJSON(data []byte) Book {
  book := Book{}
  err := json.Unmarshal(data, &book)
  if err != nil {
    panic(err)
  }
  return book
}

// BooksHandleFunc to be used as http.HandleFunc for Book API
func BooksHandleFunc(w http.ResponseWriter, r *http.Request) {

}
