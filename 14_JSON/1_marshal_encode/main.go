package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type person struct {
	FirstName string
	LastName  string
	Items     []string
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/marshal", marshal)
	http.HandleFunc("/encode", encode)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	log.Fatalln(http.ListenAndServe(":8080", nil))

}
func index(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte(`<!doctype html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport"
          content="width=device-width, user-scalable=no, initial-scale=1.0, maximum-scale=1.0, minimum-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>JSON</title>
</head>
<body>
<h3>JSON</h3>
<a href="/marshal">Marshal</a>
<a href="/encode">Encode</a>
</body>
</html>`))
	if err != nil {
		log.Fatalln(err)
	}
}
func marshal(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	p1 := person{
		FirstName: "Abhinav",
		LastName:  "Gautam",
		Items:     []string{"Laptop", "Car"},
	}
	jsonData, err := json.Marshal(p1)
	if err != nil {
		log.Fatalln(err)
	}
	w.Write(jsonData)
}
func encode(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	p1 := person{
		FirstName: "Abhinav",
		LastName:  "Gautam",
		Items:     []string{"Laptop", "Car"},
	}
	err := json.NewEncoder(w).Encode(p1)
	if err != nil {
		log.Fatalln(err)
	}
}
