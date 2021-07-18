package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func sayHello(wr http.ResponseWriter, req *http.Request) {
	log.Println("Receive input header:", req.Header)
	log.Println("Receive input body:", req.Body)
	log.Println("Entered URL by user:", req.Header.Get("Referer"))
	fmt.Fprint(wr, "<h1>Hello From GOLANG Web</h1>")
}

func main() {

	http.HandleFunc("/login", login)
	http.HandleFunc("/", sayHello)
	log.Println("+++++++++++++ Starting http web server ++++++++++")
	err := http.ListenAndServe(":4000", nil)
	checkErr(err)
}

func login(wr http.ResponseWriter, req *http.Request) {
	log.Println("Receive input header:", req.Header)
	log.Println("Receive input body:", req.Body)
	log.Println("Entered URL by user:", req.Header.Get("Referer"))
	fmt.Fprint(wr, "<h1>handle login</h1>")

	if req.Method == "GET" {
		t, _ := template.ParseFiles("login.gtpl")
		t.Execute(wr, nil)
	} else {
		req.ParseForm()

		fmt.Fprint(wr, "username: ", req.Form["username"], "<br>")
		fmt.Fprint(wr, "password: ", req.Form["password"])
	}
}

func checkErr(strErr error) {
	if strErr != nil {
		panic(strErr)

	}
}
