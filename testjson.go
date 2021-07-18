package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type toDO struct {
	UserID    int64
	ID        int
	Title     string
	Completed bool
}

func main() {

	//fmt.Println("content from URL :",
	//contentFromServer("https://jsonplaceholder.typicode.com/todos"))

	var contentInJson []toDO = parseJsonContent(contentFromServer("https://jsonplaceholder.typicode.com/todos"))

	fmt.Printf("content from URL w Json :%+v\n",
		contentInJson)

	var cntCompleted int = 0

	fmt.Println("++++++++++ Printing only Completed ones +++++++++++")
	for i := 0; i < len(contentInJson); i++ {
		if contentInJson[i].Completed {
			fmt.Println("UserID:", contentInJson[i].UserID,
				",ID:", contentInJson[i].ID,
				",Title:", contentInJson[i].Title,
				",Completed:", contentInJson[i].Completed)
			cntCompleted++
		}

	}
	fmt.Println("Total response from server:", len(contentInJson))
	fmt.Println("Total Completed task within response from server:", cntCompleted)
}

func checkErr(strErr error) {
	if strErr != nil {
		panic(strErr)

	}
}

func contentFromServer(strURL string) string {
	resp, err := http.Get(strURL)

	checkErr(err)
	defer resp.Body.Close()

	byteContent, err := ioutil.ReadAll(resp.Body)

	checkErr(err)

	return string(byteContent)

}

func parseJsonContent(strContent string) []toDO {
	todoInt := make([]toDO, 0)
	decoderToDo := json.NewDecoder(strings.NewReader(strContent))

	_, err := decoderToDo.Token()

	checkErr(err)

	var todo toDO
	for decoderToDo.More() {
		err := decoderToDo.Decode(&todo)
		checkErr(err)
		todoInt = append(todoInt, todo)
	}

	return todoInt
}
