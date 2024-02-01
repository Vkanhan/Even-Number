// package main

// import (
// 	"errors"
// 	"fmt"
// )

// func main() {
// 	num := 4
// 	if ok, err := isEven(num); ok {
// 		fmt.Printf("%d is divisible by 2: %v \n", num, ok)
// 	} else {
// 		fmt.Println(err)
// 	}
// }

// func isEven(num int) (bool, error) {
// 	if num%2 != 0 {
// 		return false, errors.New("Number not divisible by 2")
// 	}
// 	return true, nil

// }

package main

import (
	"fmt"
	"errors"
)
func main() {
	num := 4
	if ok, err := isEven(num); ok {
		fmt.Printf("%d is divisible by 2: %v", num, ok)
	} else {
		fmt.Println(err)
	}

}

func isEven(num int) (bool, error) {
	if num % 2 != 0 {
		return false, errors.New("The number is not divisible by 2")
	}
	return true, nil
}