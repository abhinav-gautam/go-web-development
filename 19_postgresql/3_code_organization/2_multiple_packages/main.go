package main

import (
	"log"
	"net/http"
	"abhinav-web-dev/19_postgresql/3_code_organization/2_multiple_packages/books"
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
