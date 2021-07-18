package main

import (
	"fmt"
	"strings"
)

func main() {
	count := 100

	for i := 0; i < count; i++ {
		fmt.Println("Hello World")

	}

	fmt.Println(strings.ToUpper("test wriTE"))
	str := "test123"
	str1 := "test1234"
	fmt.Println("new scring:", str, str1)
	fmt.Println("length:", len(str+str1))
}
