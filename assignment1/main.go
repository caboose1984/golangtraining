package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, i := range s {
		if i%2 == 0 {
			fmt.Printf("%v is even\n", i)
		} else {
			fmt.Printf("%v is odd\n", i)
		}

	}
}
