package main

import "fmt"

func main() {

	var color [5]string

	color[0] = "merah"
	color[1] = "putih"
	color[2] = "hitam"
	color[3] = "yellow"
	color[4] = "blue"

	for i := 0; i < len(color); i++ {
		fmt.Println("content ", i, ":", color[i])
	}

	var arrNum = [5]int{7, 4, 2, 3, 4}

	for i := 0; i < len(arrNum); i++ {
		fmt.Println("content ", i, ":", arrNum[i])
	}
}
