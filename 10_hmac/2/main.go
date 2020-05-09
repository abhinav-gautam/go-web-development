package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"io"
	"log"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/",index)
	http.HandleFunc("/authenticate",auth)
	http.Handle("/favicon.ico",http.NotFoundHandler())
	log.Fatalln(http.ListenAndServe(":8080",nil))
}
func index(w http.ResponseWriter,r *http.Request){
	c,err:=r.Cookie("session")
	if err == http.ErrNoCookie {
		sId := uuid.Must(uuid.NewV4())
		c = &http.Cookie{
			Name:"session",
			Value:sId.String(),
		}
	}
	if r.Method == http.MethodPost{
		email := r.FormValue("email")
		c.Value = email + "|" + getCode(email)
	}
	http.SetCookie(w,c)
	_ , err = io.WriteString(w,`<!DOCTYPE html>
	<html>
	  <body>
	    <form method="POST">
	      <input type="email" name="email">
	      <input type="submit">
	    </form>
	    <a href="/authenticate">Validate This `+c.Value+`</a>
	  </body>
	</html>`)
	if err != nil {
		log.Fatalln(err)
	}
}
func auth(w http.ResponseWriter,r *http.Request){
	c,err := r.Cookie("session")
	if err != nil {
		http.Redirect(w,r,"/",303)
	}
	if c.Value == ""{
		http.Redirect(w,r,"/",303)
	}
	xs := strings.Split(c.Value,"|")
	if xs[0]=="" && xs[1]==""{
		http.Redirect(w,r,"/",303)
	}
	email := xs[0]
	codeRcvd := xs[1]
	codeCheck := getCode(email)

	if codeRcvd != codeCheck {
		fmt.Println("HMAC codes didn't match")
		fmt.Println(codeRcvd)
		fmt.Println(codeCheck)
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	_,err = io.WriteString(w, `<!DOCTYPE html>
	<html>
	  <body>
	  	<p>`+codeRcvd+` - RECEIVED </p>
	  	<p>`+codeCheck+` - RECALCULATED </p>
	  </body>
	</html>`)
	if err != nil {
		log.Fatalln(err)
	}
}
func getCode(s string) string{
	h := hmac.New(sha256.New,[]byte("privatekey"))
	io.WriteString(h,s)
	return fmt.Sprintf("%x",h.Sum(nil))
}
