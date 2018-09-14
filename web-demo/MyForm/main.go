package main

import (
	"fmt"
	"log"
	"net/http"
	"html/template"
)

func main() {
	http.HandleFunc("/", sayHello)
	http.HandleFunc("/login", login)
	log.Println("Start Server .......... ")
	log.Fatal(http.ListenAndServe(":9090", nil))
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World");
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method)
	if r.Method == "GET" {
		t, _ := template.ParseFiles("login.html")
		log.Println(t.Execute(w, nil))
	} else {
		fmt.Println("username:", r.FormValue("username"))
		fmt.Println("password:", r.FormValue("password"))
	}
}