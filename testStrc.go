package main

import "fmt"

type customer struct {
	accountNo string
	name      string
	branch    string
}

func main() {
	var cust1 customer

	cust1 = customer{"4322", "Iwan", "Kuningan"}
	fmt.Printf("cust1 : %+v\n", cust1)
	fmt.Println("cust1 name :", cust1.name)
}
