package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	var strUrl string = "https://www.theguardian.com/international"
	resp, err := http.Get(strUrl)

	if err != nil {
		panic(err)

	}
	defer resp.Body.Close()

	byteContent, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}

	fmt.Println("contet http:", string(byteContent))

}
