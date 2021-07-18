package main

import (
	"fmt"
	"sort"
)

func main() {
	var map1 = make(map[string]string)
	//var map1 map[string]string
	map1["kampret"] = "banget"
	map1["hitam"] = "putih"
	map1["jkt"] = "jakarta"
	map1["dps"] = "denpasar"
	map1["bdg"] = "bandung"
	map1["sby"] = "surabaya"

	fmt.Println("map1 content: ", map1)
	delete(map1, "jkt")
	fmt.Println("map1 content: ", map1)
	fmt.Printf("map1 content printf: %+v\n", map1)
	for k, v := range map1 {
		fmt.Printf("the value of %+v: %+v\n", k, v)
	}

	fmt.Println("++++++++++ Sorting map +++++++++++++++")
	var mkeys = make([]string, len(map1))

	var i = 0
	for k := range map1 {
		mkeys[i] = k
		i++
	}
	fmt.Println("content of mkeys:", mkeys)
	sort.Strings(mkeys)
	fmt.Println("content of mkeys sorted:", mkeys)
	for k := range mkeys {
		fmt.Println("Content of map1 sorted ", mkeys[k], ":", map1[mkeys[k]])
	}
}
