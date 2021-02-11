package main

import "fmt"

func main() {
	var data = []int{2, 31, 3, 5, 32, 5, 7, 3, 0, 4}

	var newNumbers = func(min int) []int {
		var r []int
		for e := range data {
			if e < min {
				continue
			}
			r = append(r, e)
		}

		return r
	}(3)

	fmt.Println("original number\t:", data)
	fmt.Println("filtered number\t:", newNumbers)
}
