package main

import (
	"fmt"
)

func main() {
	// simple switch
	// i := 3

	// switch i {
	// case 1:
	// 	fmt.Println("one")
	// case 2:
	// 	fmt.Println("Two")
	// case 3:
	// 	fmt.Println("Three")
	// default:
	// 	fmt.Println("Other")
	// }

	// multiple condition switch

	// switch time.Now().Weekday() {
	// case time.Monday:
	// 	fmt.Println("Monday")
	// case time.Tuesday:
	// 	fmt.Println("Tuesday")
	// case time.Wednesday:
	// 	fmt.Println("Wednesday")
	// case time.Thursday:
	// 	fmt.Println("Thursday")
	// case time.Friday:
	// 	fmt.Println("Friday")
	// case time.Saturday:
	// 	fmt.Println("Saturday")
	// case time.Sunday:
	// 	fmt.Println("Sundayyy")
	// }

	// type switch

	whoAmI := func(i any) {
		switch i.(type) {
		case int:
			fmt.Println("Its an integer")
		case string:
			fmt.Println("Its a String")
		case bool:
			fmt.Println("Its a boolean")
		default:
			fmt.Println("Other")

		}
	}

	whoAmI(55)

}
