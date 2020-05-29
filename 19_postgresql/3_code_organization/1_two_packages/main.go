package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"abhinav-web-dev/19_postgresql/3_code_organization/1_two_packages/models"
)
var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	http.HandleFunc("/",bookIndex)
	http.HandleFunc("/books",bookIndex)
	http.HandleFunc("/books/show",bookShow)
	http.HandleFunc("/books/create",bookCreate)
	http.HandleFunc("/books/create/process",bookCreateProcess)
	http.HandleFunc("/books/update",bookUpdate)
	http.HandleFunc("/books/update/process",bookUpdateProcess)
	http.HandleFunc("/books/delete/process",bookDeleteProcess)
	log.Fatalln(http.ListenAndServe(":8080",nil))
}

// Fetching all books from database
func bookIndex(w http.ResponseWriter,r *http.Request){
	if r.Method != http.MethodGet{
		http.Error(w,http.StatusText(http.StatusMethodNotAllowed),http.StatusMethodNotAllowed)
		return
	}
	books,err:= models.AllBooks()
	if err != nil {
		http.Error(w,http.StatusText(http.StatusInternalServerError),500)
		return
	}
	w.Header().Set("Content-Type","text/html, charset=utf-8")
	tpl.ExecuteTemplate(w,"books.gohtml",books)
}
// Showing a book based on isbn
func bookShow(w http.ResponseWriter,r *http.Request){
	if r.Method != http.MethodGet{
		http.Error(w,http.StatusText(http.StatusMethodNotAllowed),405)
		return
	}
	book,err:=models.SingleBook(r)
	switch {
	case err == sql.ErrNoRows:
		http.NotFound(w,r)
		return
	case err !=nil:
		http.Error(w,http.StatusText(http.StatusInternalServerError),500)
		return
	}
	fmt.Fprintf(w,"%s,%s,%s,$%.2f\n",book.Isbn,book.Title,book.Author,book.Price)
}
// Book create form
func bookCreate(w http.ResponseWriter,r *http.Request){
	tpl.ExecuteTemplate(w,"create.gohtml",nil)
}
// Creating a new book
func bookCreateProcess(w http.ResponseWriter,r *http.Request){
	// Checking method
	if r.Method != http.MethodPost {
		http.Error(w,http.StatusText(http.StatusMethodNotAllowed),405)
		return
	}
	err := models.CreateBookProcess(r)
	if err != nil {
		http.Error(w,http.StatusText(http.StatusInternalServerError),500)
		return
	}
	// Confirm insertion
	tpl.ExecuteTemplate(w,"created.gohtml",nil)
}
// Book update form
func bookUpdate(w http.ResponseWriter,r *http.Request){
	if r.Method != http.MethodGet{
		http.Error(w,http.StatusText(http.StatusMethodNotAllowed),405)
		return
	}
	book,err:=models.UpdateBookForm(r)
	switch {
	case err == sql.ErrNoRows:
		http.NotFound(w,r)
		return
	case err !=nil:
		http.Error(w,http.StatusText(http.StatusInternalServerError),500)
		return
	}
	tpl.ExecuteTemplate(w,"update.gohtml",book)
}

//Updating book
func bookUpdateProcess(w http.ResponseWriter,r *http.Request){
	if r.Method != http.MethodPost{
		http.Error(w,http.StatusText(http.StatusMethodNotAllowed),405)
		return
	}
	err:=models.UpdateBookProcess(r)
	if err != nil {
		http.Error(w,http.StatusText(http.StatusInternalServerError),500)
		return
	}
	tpl.ExecuteTemplate(w,"updated.gohtml",nil)
}

// Deleting book
func bookDeleteProcess(w http.ResponseWriter,r *http.Request){
	if r.Method != http.MethodGet {
		http.Error(w,http.StatusText(http.StatusMethodNotAllowed),405)
		return
	}
	err:=models.DeleteBookProcess(r)
	if err != nil {
		http.Error(w,http.StatusText(http.StatusInternalServerError),500)
		return
	}
	http.Redirect(w,r,"/books",http.StatusSeeOther)
}