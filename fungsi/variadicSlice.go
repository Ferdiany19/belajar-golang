package main

import "fmt"

func main() {
	var data = []int{2, 4, 3, 5, 6, 3, 2, 5, 6, 3, 2, 2, 4, 5, 6, 1, 23, 6, 2} // bedanya disini dia di slice
	var avg = calculate(data...)
	var msg = fmt.Sprintf("rata-rata : %.2f", avg)
	fmt.Printf(msg)
}

func calculate(numbers ...int) (avg float64) {
	var total int = 0
	for number := range numbers {
		total += number
	}

	avg = float64(total) / float64(len(numbers))
	return
}
