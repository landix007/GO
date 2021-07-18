package main

import (
	"fmt"
	"math/big"
)

func main() {

	int1, str := 12, "test"
	fmt.Println("test:", float64(int1)+5)
	fmt.Println("test str:", str)

	var b1, b2, b3, bigSum big.Float

	b1.SetFloat64(1.3)
	b2.SetFloat64(3.3)
	b3.SetFloat64(1.8)
	bigSum.Add(&b1, &b2).Add(&bigSum, &b3)
	fmt.Println("big summary float: ", bigSum.String())
	fmt.Printf("big summary float in printf: %.2f\n", &bigSum)
}
