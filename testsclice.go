package main

import "fmt"

func main() {

	var color = []string{"blue", "red", "white"}

	fmt.Println("slices len: ", len(color))

	for i := 0; i < len(color); i++ {
		fmt.Println("content ", i, ":", color[i])
	}
	color = append(color, "black")
	for i := 0; i < len(color); i++ {
		fmt.Println("content ", i, ":", color[i])
	}
	color = append(color[2:len(color)])

	fmt.Println("+++++++++++++++++++++")
	for i := 0; i < len(color); i++ {
		fmt.Println("content ", i, ":", color[i])
	}

	//color = append(color[1:])
	color = append(color[:len(color)-1])

	fmt.Println("+++++++++++++++++++++")
	for i := 0; i < len(color); i++ {
		fmt.Println("content ", i, ":", color[i])
	}

	//var color1 = []string{}
	var color1 = make([]string, 5, 5)

	fmt.Println("color1 cap: ", cap(color1))
	color1 = append(color1, "white")
	fmt.Println("color1 len: ", len(color1))
	fmt.Println("color1 cap: ", cap(color1))

	for i := 0; i < len(color1); i++ {
		fmt.Println("content color1", i, ":", color1[i])
	}

}
