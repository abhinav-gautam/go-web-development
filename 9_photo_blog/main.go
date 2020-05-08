package main

import (
	"crypto/sha1"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}
func main() {
	http.HandleFunc("/",index)
	http.Handle("/favicon.ico",http.NotFoundHandler())
	http.Handle("/public/",http.StripPrefix("/public",http.FileServer(http.Dir("./public"))))
	log.Fatalln(http.ListenAndServe(":8080",nil))
}
func index(w http.ResponseWriter,r *http.Request){
	c := getCookie(w,r)
	if r.Method == http.MethodPost{
		// Getting file from form
		mf,fh,err := r.FormFile("nf")
		if err != nil {
			log.Fatal(err)
		}
		defer mf.Close()
		// Creating SHA1 for filename
		ext := strings.Split(fh.Filename,".")[1]
		h:=sha1.New()
		io.Copy(h,mf)
		fName := fmt.Sprintf("%x",h.Sum(nil)) + "." + ext
		// Creating new file
		wd,err := os.Getwd()
		if err != nil {
			log.Fatalln(err)
		}
		path := filepath.Join(wd,"public","pics",fName)
		nf,err := os.Create(path)
		if err != nil {
			log.Fatalln(err)
		}
		defer nf.Close()
		// Copy
		mf.Seek(0,0)
		io.Copy(nf,mf)
		// Adding file name to user's cookie
		c = appendValue(w,c,fName)
	}
	xs := strings.Split(c.Value,"|")
	tpl.ExecuteTemplate(w,"index.gohtml",xs[1:])
}
func getCookie(w http.ResponseWriter,r *http.Request) *http.Cookie{
	c,err := r.Cookie("session")
	if err == http.ErrNoCookie {
		sID := uuid.Must(uuid.NewV4())
		c = &http.Cookie{
			Name:"session",
			Value:sID.String(),
		}
		http.SetCookie(w,c)
	}
	return c
}
func appendValue(w http.ResponseWriter,c *http.Cookie,fName string) *http.Cookie{
	s:=c.Value
	if !strings.Contains(s,fName){
		s += "|" + fName
	}
	c.Value = s
	http.SetCookie(w,c)
	return c
}
