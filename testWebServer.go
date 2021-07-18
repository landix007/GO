package main

import (
	"fmt"
	"log"
	"net/http"
)

type HelloWeb struct{}

func (hlo HelloWeb) ServeHTTP(wr http.ResponseWriter, req *http.Request) {
	log.Println("Receive input header:", req.Header)
	log.Println("Receive input body:", req.Body)
	//log.Println("Entered URL by user:", req.)
	fmt.Fprint(wr, "<h1>Hello From GOLANG Web</h1>")
}

func main() {
	var he HelloWeb

	log.Println("+++++++++++++ Starting http web server ++++++++++")
	err := http.ListenAndServe(":4000", he)
	checkErr(err)
}

func checkErr(strErr error) {
	if strErr != nil {
		panic(strErr)

	}
}
