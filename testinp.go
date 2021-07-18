package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	//var str string

	//fmt.Scanln(&str)
	//fmt.Println(str)
	var strinp string

	fmt.Println("++++++++++++++ PASSING STRING ++++++++++++++++++")
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	strinp, _ = reader.ReadString('\n')
	fmt.Println("text inputed: ", strinp)
	fmt.Println("length of text inputed: ", len(strinp))

	//passing number

	fmt.Println("++++++++++++++ PASSING NUMBER ++++++++++++++++++")

	fmt.Print("Enter a number:")
	strinp, _ = reader.ReadString('\n')
	f, errinp := strconv.ParseFloat(strings.TrimSpace(strinp), 64)
	if errinp != nil {
		fmt.Println(errinp)
	} else {
		fmt.Println("Value of number: ", f)
	}
}
