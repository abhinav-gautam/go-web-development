package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

var db *sql.DB
var tpl *template.Template

type book struct {
	Isbn string
	Title string
	Author string
	Price float32
}
func init() {
	var err error
	db,err = sql.Open("postgres","postgres://abhinav:91189@localhost/bookstore?sslmode=disable")
	if err != nil {
		panic(err)
	}
	fmt.Println("[+] Database Connected")

	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}
func main() {
	defer db.Close()
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
	rows,err := db.Query("SELECT * FROM books")
	if err != nil {
		fmt.Println(err)
		http.Error(w,http.StatusText(http.StatusInternalServerError),http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	books := make([]book,0)

	for rows.Next(){
		book := book{}
		err:= rows.Scan(&book.Isbn,&book.Title,&book.Author,&book.Price)
		if err != nil {
			http.Error(w,http.StatusText(http.StatusInternalServerError),http.StatusInternalServerError)
			return
		}
		books = append(books,book)
	}
	if err=rows.Err(); err!=nil{
		panic(err)
	}
	//for _,book := range books{
	//	fmt.Fprintf(w,"%s,%s,%s,$%.2f\n",book.Isbn,book.Title,book.Author,book.Price)
	//}
	w.Header().Set("Content-Type","text/html, charset=utf-8")
	tpl.ExecuteTemplate(w,"books.gohtml",books)
}
// Showing a book based on isbn
func bookShow(w http.ResponseWriter,r *http.Request){
	if r.Method != http.MethodGet{
		http.Error(w,http.StatusText(http.StatusMethodNotAllowed),405)
		return
	}
	isbn := r.FormValue("isbn")
	if isbn==""{
		http.Error(w,http.StatusText(http.StatusBadRequest),400)
		return
	}
	book := book{}

	row := db.QueryRow("SELECT * FROM books WHERE isbn = $1",isbn)
	err:=row.Scan(&book.Isbn,&book.Title,&book.Author,&book.Price)
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

	// Getting form values
	book := book{}
	book.Isbn = r.FormValue("isbn")
	book.Title = r.FormValue("title")
	book.Author = r.FormValue("author")
	price := r.FormValue("price")

	// Validating form values
	if book.Isbn=="" || book.Title=="" || book.Author=="" || price==""{
		http.Error(w,http.StatusText(http.StatusBadRequest),400)
		return
	}
	// Convert form values
	f64,err := strconv.ParseFloat(price,32)
	if err != nil {
		http.Error(w,http.StatusText(http.StatusNotAcceptable),406)
		return
	}
	book.Price = float32(f64)

	// Inserting values
	_,err = db.Exec("INSERT INTO books(isbn,title,author,price) VALUES($1,$2,$3,$4)",book.Isbn,book.Title,book.Author,book.Price)
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

	isbn := r.FormValue("isbn")
	if isbn == ""{
		http.Error(w,http.StatusText(http.StatusBadRequest),400)
		return
	}
	book := book{}
	row := db.QueryRow("SELECT * FROM books WHERE isbn = $1",isbn)
	err:=row.Scan(&book.Isbn,&book.Title,&book.Author,&book.Price)
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

	book := book{}
	book.Isbn = r.FormValue("isbn")
	book.Title = r.FormValue("title")
	book.Author = r.FormValue("author")
	price := r.FormValue("price")

	f64,err:=strconv.ParseFloat(price,32)
	if err != nil {
		http.Error(w,http.StatusText(http.StatusNotAcceptable),406)
		return
	}
	book.Price = float32(f64)

	_,err = db.Exec("UPDATE books SET isbn=$1, title=$2, author=$3, price=$4 WHERE isbn=$1",book.Isbn,book.Title,book.Author,book.Price)
	if err != nil {
		http.Error(w,http.StatusText(500),500)
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
	isbn := r.FormValue("isbn")
	if isbn == ""{
		http.Error(w,http.StatusText(http.StatusBadRequest),400)
		return
	}
	_,err := db.Exec("DELETE FROM books WHERE isbn=$1",isbn)
	if err != nil {
		http.Error(w,http.StatusText(http.StatusInternalServerError),500)
		return
	}

	http.Redirect(w,r,"/books",http.StatusSeeOther)
}