package main

import (
	"fmt"
	"time"
)

func main() {

	var tm time.Time

	tm = time.Date(2010, time.February, 1, 1, 1, 1, 1, time.Now().Local().Location())
	fmt.Println("time is set into: ", tm.Location())
	fmt.Println("time is set into: ", tm)
	fmt.Println("Now is : ", time.Now())
	fmt.Printf("Now is in printf: %s\n", time.Now().UTC())

	//time1, _ := time.Parse("ddmmyyyy", "25102011")
	fmt.Println("time1 is : ", tm.Format("02012006"))
	fmt.Println("now is : ", time.Now().Format("02-01-2006"))

	fmt.Println("tomorrow is: ", time.Now().AddDate(0, 0, 1).Format("02-01-2006"))
	fmt.Println("1 month from now  is: ", time.Now().AddDate(0, 1, 0).Format("02-01-2006"))
	fmt.Println("1 year from now  is: ", time.Now().AddDate(1, 0, 0).Format("02-01-2006"))
}
