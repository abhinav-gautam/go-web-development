package books

import (
	"errors"
	"net/http"
	"strconv"
	"abhinav-web-dev/19_postgresql/3_code_organization/2_multiple_packages/config"
)

type book struct {
	Isbn string
	Title string
	Author string
	Price float32
}

func AllBooks()([]book, error){
	rows,err := config.DB.Query("SELECT * FROM books")
	if err != nil {
		return nil,err
	}
	defer rows.Close()

	books := make([]book,0)

	for rows.Next(){
		book := book{}
		err:= rows.Scan(&book.Isbn,&book.Title,&book.Author,&book.Price)
		if err != nil {
			return nil,err
		}
		books = append(books,book)
	}
	if err=rows.Err(); err!=nil{
		return nil,err
	}
	//for _,book := range books{
	//	fmt.Fprintf(w,"%s,%s,%s,$%.2f\n",book.Isbn,book.Title,book.Author,book.Price)
	//}
	return books,nil
}
func SingleBook(r *http.Request) (book,error){
	book := book{}
	isbn := r.FormValue("isbn")
	if isbn==""{
		return book,errors.New("Bad Request.")
	}

	row := config.DB.QueryRow("SELECT * FROM books WHERE isbn = $1",isbn)
	err:=row.Scan(&book.Isbn,&book.Title,&book.Author,&book.Price)
	if err != nil {
		return book,err
	}
	return book,nil
}
func CreateBookProcess(r *http.Request) error{
	// Getting form values
	book := book{}
	book.Isbn = r.FormValue("isbn")
	book.Title = r.FormValue("title")
	book.Author = r.FormValue("author")
	price := r.FormValue("price")

	// Validating form values
	if book.Isbn=="" || book.Title=="" || book.Author=="" || price==""{
		return errors.New("Bad Request.")
	}
	// Convert form values
	f64,err := strconv.ParseFloat(price,32)
	if err != nil {
		return err
	}
	book.Price = float32(f64)

	// Inserting values
	_,err = config.DB.Exec("INSERT INTO books(isbn,title,author,price) VALUES($1,$2,$3,$4)",book.Isbn,book.Title,book.Author,book.Price)
	if err != nil {
		return err
	}
	return nil
}
func UpdateBookForm(r *http.Request)(book,error){
	book := book{}
	isbn := r.FormValue("isbn")
	if isbn == ""{
		return book,errors.New("Bad Request.")
	}
	row := config.DB.QueryRow("SELECT * FROM books WHERE isbn = $1",isbn)
	err:=row.Scan(&book.Isbn,&book.Title,&book.Author,&book.Price)
	if err != nil {
		return book,err
	}
	return book,nil
}
func UpdateBookProcess(r *http.Request)error{
	book := book{}
	book.Isbn = r.FormValue("isbn")
	book.Title = r.FormValue("title")
	book.Author = r.FormValue("author")
	price := r.FormValue("price")

	f64,err:=strconv.ParseFloat(price,32)
	if err != nil {
		return err
	}
	book.Price = float32(f64)

	_,err = config.DB.Exec("UPDATE books SET isbn=$1, title=$2, author=$3, price=$4 WHERE isbn=$1",book.Isbn,book.Title,book.Author,book.Price)
	if err != nil {
		return err
	}
	return nil
}
func DeleteBookProcess(r *http.Request)error {
	isbn := r.FormValue("isbn")
	if isbn == "" {
		return errors.New("Bad Request.")
	}
	_, err := config.DB.Exec("DELETE FROM books WHERE isbn=$1", isbn)
	if err != nil {
		return err
	}
	return nil
}
