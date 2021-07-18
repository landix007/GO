package main

import (
	"fmt"
	"strings"
)

func main() {

	//fmt.Println("part of: ",strings.Contains())
	if !strings.Contains("kampret", "ampr") {
		fmt.Println("OK")

	} else {
		fmt.Println("NOT OK")
	}
}
