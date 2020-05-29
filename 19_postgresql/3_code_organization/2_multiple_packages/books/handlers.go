package books

import (
	"abhinav-web-dev/19_postgresql/3_code_organization/2_multiple_packages/config"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
)

// Fetching all books from database
func BookIndex(w http.ResponseWriter,r *http.Request){
	if r.Method != http.MethodGet{
		http.Error(w,http.StatusText(http.StatusMethodNotAllowed),http.StatusMethodNotAllowed)
		return
	}
	books,err:= AllBooks()
	if err != nil {
		http.Error(w,http.StatusText(http.StatusInternalServerError),500)
		return
	}
	w.Header().Set("Content-Type","text/html, charset=utf-8")
	config.TPL.ExecuteTemplate(w,"books.gohtml",books)
}
// Showing a book based on isbn
func BookShow(w http.ResponseWriter,r *http.Request){
	if r.Method != http.MethodGet{
		http.Error(w,http.StatusText(http.StatusMethodNotAllowed),405)
		return
	}
	book,err:=SingleBook(r)
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
func BookCreate(w http.ResponseWriter,r *http.Request){
	config.TPL.ExecuteTemplate(w,"create.gohtml",nil)
}
// Creating a new book
func BookCreateProcess(w http.ResponseWriter,r *http.Request){
	// Checking method
	if r.Method != http.MethodPost {
		http.Error(w,http.StatusText(http.StatusMethodNotAllowed),405)
		return
	}
	err := CreateBookProcess(r)
	if err != nil {
		http.Error(w,http.StatusText(http.StatusInternalServerError),500)
		return
	}
	// Confirm insertion
	config.TPL.ExecuteTemplate(w,"created.gohtml",nil)
}
// Book update form
func BookUpdate(w http.ResponseWriter,r *http.Request){
	if r.Method != http.MethodGet{
		http.Error(w,http.StatusText(http.StatusMethodNotAllowed),405)
		return
	}
	book,err:=UpdateBookForm(r)
	switch {
	case err == sql.ErrNoRows:
		http.NotFound(w,r)
		return
	case err !=nil:
		http.Error(w,http.StatusText(http.StatusInternalServerError),500)
		return
	}
	config.TPL.ExecuteTemplate(w,"update.gohtml",book)
}

//Updating book
func BookUpdateProcess(w http.ResponseWriter,r *http.Request){
	if r.Method != http.MethodPost{
		http.Error(w,http.StatusText(http.StatusMethodNotAllowed),405)
		return
	}
	err:=UpdateBookProcess(r)
	if err != nil {
		http.Error(w,http.StatusText(http.StatusInternalServerError),500)
		return
	}
	config.TPL.ExecuteTemplate(w,"updated.gohtml",nil)
}

// Deleting book
func BookDeleteProcess(w http.ResponseWriter,r *http.Request){
	if r.Method != http.MethodGet {
		http.Error(w,http.StatusText(http.StatusMethodNotAllowed),405)
		return
	}
	err:=DeleteBookProcess(r)
	if err != nil {
		http.Error(w,http.StatusText(http.StatusInternalServerError),500)
		return
	}
	http.Redirect(w,r,"/books",http.StatusSeeOther)
}

// Json of books
func ShowJson(w http.ResponseWriter,r *http.Request){
	if r.Method != http.MethodGet {
		http.Error(w,http.StatusText(http.StatusMethodNotAllowed),405)
		return
	}
	books,err := AllBooks()
	if err != nil {
		http.Error(w,http.StatusText(http.StatusInternalServerError),500)
		return
	}
	jsonData,err := json.Marshal(books)
	if err != nil {
		http.Error(w,http.StatusText(http.StatusInternalServerError),500)
		return
	}
	w.Header().Set("Content-Type","application/json")
	w.Write(jsonData)
}
