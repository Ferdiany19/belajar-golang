package main

import "fmt"

func main() {
	var getMinMax = func(n []int) (min int, max int) {
		for i, e := range n {
			switch {
			case i == 0:
				min, max = e, e
			case e > max:
				max = e
			case e < max:
				min = e
			}
		}
		return
	}

	var data = []int{2, 5, 3, 4, 32, 12, 32, 52, 11, 2, 3}
	var min, max = getMinMax(data)
	fmt.Printf("data : %v\nmin : %v\nmax %v\n ", data, min, max)
}
