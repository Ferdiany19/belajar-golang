package main

import "fmt"

// LOOPING
// func main() {
// outerLoop:
// 	for i := 0; i < 5; i++ {
// 		for j := 0; j < 5; j++ {
// 			if i == 3 {
// 				break outerLoop
// 			}
// 			fmt.Print("Matriks ['", i, "][", j, "]", "\n")
// 		}
// 	}
// }

func main() {
	// var names = [4]string{
	// 	"ferdian",
	// 	"Yuliansyah",
	// 	"ganteng",
	// 	"banget",
	// }

	// fmt.Println(names[0], names[1], names[2], names[3])

	// LOOPING ARRAY

	var fruits = [...]string{
		"apple",
		"grape",
		"banana",
		"rambutans",
		"durians",
	}

	// FOOR LOOP BIASA
	// for i := 0; i < len(fruits); i++ {
	// 	fmt.Printf("elemen %d : %s\n", i, fruits[i])
	// }

	// FOR RANGE
	// for i, fruit := range fruits {
	// 	fmt.Printf("elemen %d : %s\n", i, fruit)
	// }

	// FOR RANGE MENGGUNAKAN UNDERSCORE

	for _, fruit := range fruits {
		fmt.Printf("ini adalah : %s\n", fruit)
	}
}
