package main

import "fmt"

func main() {
	//variables whose value can not changes later

	const name = "goLang"
	const port = 5000
	const host = "localhost"

	fmt.Println(name, port, host )

	// host = "server"  // not allowed in go

}