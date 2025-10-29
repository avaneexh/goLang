package main

import (
	"fmt"
)

// slices -> dynamic size
//  most used construct in go
// + useful methods
func main(){
	// uninitialized slice is nil
	// var nums []int  

	// fmt.Println(nums)
	// fmt.Println(len(nums))

	// var nums = make([]int, 0, 5)
	// 2 is initial size and 5 is initial capacity
	// fmt.Println(len(nums))
    // fmt.Println(nums)
	// fmt.Println(cap(nums)) // capacity maximum no of elements scan fit 
	// resize always doubles the size of slice

	// nums = append(nums, 1)
	// nums = append(nums, 2)
	// nums = append(nums, 3)
	// nums = append(nums, 4)
	// nums = append(nums, 5)
	// nums = append(nums, 6)
	// nums = append(nums, 7)
	// nums = append(nums, 8)
	// nums = append(nums, 9)
	// nums = append(nums, 10)

	// nums := []int{}
	// nums = append(nums, 1)
	// nums = append(nums, 1)
	// nums = append(nums, 1)

	// var nums = make([]int, 2, 5)

	// nums[0] = 3
	// nums[1] = 5

	// fmt.Println(nums)
	// fmt.Println(cap(nums))
	// fmt.Println(len(nums))


	// copy functions

	// var nums = make([]int, 0, 5)
	// nums = append(nums, 2)
	// var nums2 = make([]int, len(nums) )



	// copy(nums2, nums)

	// fmt.Println(nums, nums2)


	// slice operator 

	// var nums = []int{1,2,3}


	// fmt.Println((nums[0:2]))
	// fmt.Println((nums[:1]))
	// fmt.Println((nums[1:]))
	

	// slices package

	// var nums = []int{1,2,3}
	// var nums2 = []int{1,2}

	// fmt.Println(slices.Equal(nums, nums2))

	var nums = [][]int{{1,2,3},{2,3,4}}
	fmt.Println(nums)
}