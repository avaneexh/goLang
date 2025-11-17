package main

import (
	"fmt"
	"maps"
)

// maps -> hash, object, dict

func main() {
	//  creating maps

	// m:=make(map[string]string)

	// // setting an elements

	// m["name"] = "golang"
	// m["area"] = "backend"
	// m["price"] = "50"

	// // get element

	// fmt.Println(m["name"])
	// fmt.Println(m["area"])
	// fmt.Println(m["price"])

	// // if key does not exists in map then it returs 0
	// fmt.Println(m["phone"])

	// n:=make(map[string]int)

	// n["age"] = 20
	// fmt.Println(n["age"])

	// fmt.Println(len(m))

	// delete function

	// m:=make(map[string]string)

	// m["name"] = "golang"
	// m["area"] = "backend"
	// m["price"] = "50"

	// fmt.Println(m)

	// delete(m, "price")

	// fmt.Println(m)

	// clear(m)

	// fmt.Println(m)

	// m:= map[string]int{"price": 40, "phone":3}

	// fmt.Println(m);

	// m:= map[string]int{"price": 40, "phone":3}

	// k, ok := m["price"]

	// fmt.Println(k)

	// if ok {
	// 	fmt.Println("all ok")
	// } else {
	// 	fmt.Println("not ok ")
	// }

	m1 := map[string]int{"price": 40, "phone": 3}
	m2 := map[string]int{"price": 40, "phone": 3}

	fmt.Println(maps.Equal(m1, m2))
}
