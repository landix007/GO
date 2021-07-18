package main

import "fmt"

func main() {

	var p *int

	if p == nil {
		fmt.Println("p is nil")
		v := 42
		p = &v
		fmt.Println("p is ", *p)

		*p = *p + 10
		//p = &v + 10
		fmt.Println("p after added is ", *p)
		fmt.Println("current v value ", v)
	} else {
		fmt.Println("p is NOT nil")
	}

}
