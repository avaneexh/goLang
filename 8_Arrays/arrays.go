package main

import "fmt"

// we use Array for fixed size
// Memory Optimization 
// constant time access


func main() {
	// Array is numbered sequence of any specific length
	// var nums [4]int


	// nums[0] = 1

	// fmt.Println(nums[0])
	// fmt.Println(nums)

	// array length
	// fmt.Println(len(nums))

	// var vals [4]string

	// vals[0] = "golang"

	// fmt.Println(vals)
	// var vals [4]bool

	// vals[2] =true

	// fmt.Println(vals)




	// Zeroed Vlaues

	//  int -> 0, String->"", bool -> False

	// To declare it in single line 
	// num := [3]int {1, 4, 6}


	// fmt.Println(num)

	// 2d arrays

	nums := [2][2]int {{3,4},{5,6}}

	fmt.Println(nums)

}