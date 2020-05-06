package main

import (
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

func alreadyLoggedIn(w http.ResponseWriter,r *http.Request) bool{
	// Checking if already logged in
	c,err := r.Cookie("session")
	if err == http.ErrNoCookie {
		return false
	}
	// Refreshing session
	s,ok := dbSession[c.Value]
	if ok{
		s.LastActivity = time.Now()
		dbSession[c.Value] = s
	}
	c.MaxAge = sessionLength
	http.SetCookie(w,c)

	username := dbSession[c.Value].Username
	_,ok = dbUser[username]
	return ok
}
func signUp(w http.ResponseWriter,r *http.Request){
	if alreadyLoggedIn(w,r){
		http.Redirect(w,r,"/dashboard",http.StatusSeeOther)
		return
	}
	if r.Method == http.MethodPost {
		c,err := r.Cookie("session")
		if err == http.ErrNoCookie {
			// Creating new cookie
			sID := uuid.Must(uuid.NewV4())
			c = &http.Cookie{
				Name:"session",
				Value:sID.String(),
			}
		}
		http.SetCookie(w,c)
		r.ParseForm()
		username := r.FormValue("username")
		name := r.FormValue("name")
		password := r.FormValue("password")
		role := r.FormValue("role")

		// Encrypting Password
		encPassword, err:= bcrypt.GenerateFromPassword([]byte(password),bcrypt.MinCost)
		if err != nil {
			http.Error(w,"Internal Server Error",http.StatusInternalServerError)
		}

		// Checking if empty
		if username=="" || name=="" || password==""{
			http.Redirect(w,r,"/",http.StatusSeeOther)
			return
		}

		// Checking username availability
		if _,ok := dbUser[username]; ok{
			http.Error(w,"Username Not Available",http.StatusForbidden)
			return
		}
		// Adding user to map
		u := user{username,name,encPassword,role}
		dbUser[username] = u
		dbSession[c.Value] = session{username,time.Now()}
		http.Redirect(w,r,"/dashboard",http.StatusSeeOther)
	}
	tpl.ExecuteTemplate(w,"signUpForm.gohtml",nil)
}
func getUser(r *http.Request) user{
	c,err := r.Cookie("session")
	var u user
	if err != nil {
		return u
	}
	username := dbSession[c.Value].Username
	u = dbUser[username]
	return u
}
func logout(w http.ResponseWriter,r *http.Request){
	c, err:=r.Cookie("session")
	if err == http.ErrNoCookie {
		http.Redirect(w,r,"/",http.StatusSeeOther)
		return
	}
	delete(dbSession,c.Value)
	c.MaxAge = -1
	http.SetCookie(w,c)

	// clean up dbSessions
	if time.Now().Sub(dbSessionsCleaned) > (time.Second * 30) {
		go cleanSessions()
	}

	http.Redirect(w,r,"/",http.StatusSeeOther)
}
func login(w http.ResponseWriter,r *http.Request){
	if alreadyLoggedIn(w,r){
		http.Redirect(w,r,"/dashboard",http.StatusSeeOther)
		return
	}
	if r.Method == http.MethodPost{
		r.ParseForm()
		username := r.FormValue("username")
		password := r.FormValue("password")

		var u user
		// Checking User Exist
		u,ok := dbUser[username]
		if !ok{
			http.Error(w,"Incorrect Username and/or Password",http.StatusForbidden)
			return
		}
		// Checking password match
		err := bcrypt.CompareHashAndPassword(u.Password,[]byte(password))
		if err != nil {
			http.Error(w,"Incorrect Username and/or Password",http.StatusForbidden)
			return
		}
		// Creating new session
		sID := uuid.Must(uuid.NewV4())
		c := &http.Cookie{
			Name:"session",
			Value:sID.String(),
		}
		http.SetCookie(w,c)
		// Adding to session map
		dbSession[c.Value] = session{username,time.Now()}
		// Redirecting to dashboard
		http.Redirect(w,r,"/dashboard",http.StatusSeeOther)
	}

	tpl.ExecuteTemplate(w,"loginForm.gohtml",nil)
}

func cleanSessions(){
	for k,v := range dbSession{
		if time.Now().Sub(v.LastActivity)>(time.Second*30){
			delete(dbSession,k)
		}
	}
}