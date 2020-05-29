package books

import (
	"abhinav-web-dev/18_mongodb/2_bookstore_crud/config"
	"errors"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"strconv"
)

type book struct {
	Isbn string `json:"isbn" bson:"isbn"`
	Title string `json:"title" bson:"title"`
	Author string `json:"author" bson:"author"`
	Price float32 `json:"price" bson:"price"`
}

func AllBooks()([]book, error){
	books := make([]book,0)
	err:=config.Books.Find(bson.M{}).All(&books)
	if err != nil {
		return nil,err
	}
	return books,nil
}
func SingleBook(r *http.Request) (book,error){
	book := book{}
	isbn := r.FormValue("isbn")
	if isbn==""{
		return book,errors.New("Bad Request.")
	}

	err := config.Books.Find(bson.M{"isbn":isbn}).One(&book)
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
	err = config.Books.Insert(book)
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
	err:=config.Books.Find(bson.M{"isbn":isbn}).One(&book)
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

	err=config.Books.Update(bson.M{"isbn":book.Isbn},&book)
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
	err := config.Books.Remove(bson.M{"isbn":isbn})
	if err != nil {
		return err
	}
	return nil
}
