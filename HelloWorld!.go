package main

import (
	"fmt"
)

func main() {
	var input int
	fmt.Print("Masukkan angka: ")
	fmt.Scanf("%d", &input)

	if input == 42 {
		fmt.Println("Hello Universe!")
	} else {
		fmt.Printf("Input Angkamu adalah: %d\n", input)
	}
}
