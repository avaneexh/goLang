package main

// import "fmt"

// func add(a int, b int) int {
// 	return a + b
// }

// func getLanguages() (string, string, string, bool) {
// 	return "golang", "js", "c", true
// }

// func processIt(fn func(a int) int){
// 	fn(1)
// }

func processIt() func(a int) int{
	return func(a int) int {
		return 4
	}
}


func main() {
	// fn := func(a int) int {
	// 	return 2
	// }

	fn :=processIt()
	fn(6)

	// fmt.Println(processIt())


	// result := add(3, 5)
	// fmt.Println("Result",result)
	// lang1, lang2, _, tf := getLanguages()
	// fmt.Println(getLanguages())
	// fmt.Println(lang1)
	// fmt.Println(lang2)
	// fmt.Println(tf)
	
}