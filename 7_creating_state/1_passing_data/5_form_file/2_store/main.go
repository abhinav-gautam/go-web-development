package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}
func main() {
	http.HandleFunc("/",index)
	http.Handle("/favicon.ico",http.NotFoundHandler())
	log.Fatalln(http.ListenAndServe(":8080",nil))
}
func index(w http.ResponseWriter,r *http.Request){
	var s string
	if r.Method == http.MethodPost{
		// Reading File
		f,fh,err:=r.FormFile("file")
		if err != nil {
			http.Error(w,err.Error(),http.StatusInternalServerError)
			return
		}
		defer f.Close()

		// Logging
		fmt.Println("File: ",fh.Filename,"\nFile Header: ",fh.Header,"\nError: ",err)

		bs ,err := ioutil.ReadAll(f)
		if err != nil {
			http.Error(w,err.Error(),http.StatusInternalServerError)
			return
		}

		s = string(bs)

		// Writing File
		dst,err := os.Create(filepath.Join("./users/",fh.Filename))

		if err != nil {
			http.Error(w,err.Error(),http.StatusInternalServerError)
			return
		}

		defer dst.Close()

		_,err = dst.Write(bs)

		if err != nil {
			http.Error(w,err.Error(),http.StatusInternalServerError)
			return
		}

	}
	tpl.ExecuteTemplate(w,"index.gohtml",s)
}
