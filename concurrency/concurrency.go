package concurrency

import "fmt"

func IsEven(input int) (result bool) {
	if input%2 == 0 {
		fmt.Println(true)
		return true
	} else {
		fmt.Println(false)
		return false
	}
}
