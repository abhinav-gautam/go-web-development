package main

import (
	"abhinav-web-dev/18_mongodb/2_bookstore_crud/books"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/",books.BookIndex)
	http.HandleFunc("/books",books.BookIndex)
	http.HandleFunc("/books/show",books.BookShow)
	http.HandleFunc("/books/create",books.BookCreate)
	http.HandleFunc("/books/create/process",books.BookCreateProcess)
	http.HandleFunc("/books/update",books.BookUpdate)
	http.HandleFunc("/books/update/process",books.BookUpdateProcess)
	http.HandleFunc("/books/delete/process",books.BookDeleteProcess)
	http.HandleFunc("/books/json",books.ShowJson)
	log.Fatalln(http.ListenAndServe(":8080",nil))
}
