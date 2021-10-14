package main

import "fmt"

func main() {
	fmt.Println("Hallo Haris Torang Bisa")
	// komentar kode
	// menampilkan pesan hello world
	var firstname string = "Haris"

	var lastname string
	lastname = "Manaf"

	fmt.Println("Hallo !", firstname, lastname)

	/////////////////////////////////////////////////////////////

	var point = 3

	switch {
	case point == 8:
		fmt.Println("perfect")
	case (point < 8) && (point > 3):
		fmt.Println("awesome")
		fallthrough
	case point < 5:
		fmt.Println("you must be better broo")
	default:
		{
			fmt.Println("not bad")
			fmt.Println("you need to learn more")
		}
	}
	////////////////////

	// for i := 1; i < 5; i++ {
	// 	fmt.Println("angka", i)
	// }

	// var i = 0

	// for {
	// 	fmt.Println("angka", i)
	// 	i++

	// 	if i == 4 {
	// 		fmt.Println("mantap jiwa")
	// 		break

	// 	}
	// }

outerLoop:
	for i := 0; i < 5; i++ {
		for j := 1; j < 5; j++ {
			if i == 3 {
				break outerLoop
			}
			fmt.Print("Matriks, [", i, "][", j, "]", "\n")
		}
	}
}
