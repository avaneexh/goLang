package main

import "fmt"

//iterating over data structure

func main() {
	// nums := []int {6, 7,8}

	// for i :=0; i<len(nums); i++ {
	// 	fmt.Println(nums[i])
	// }

	// for _, num := range nums {
	// 	fmt.Println(num)
	// }
	m := map[string]string{"fname": "jhon", "lname": "doe"}

	for k, v := range m {
		fmt.Println(k, v)
	}

}
