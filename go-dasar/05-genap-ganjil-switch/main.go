package main

import "fmt"

func main() {
	angka := 7
	switch angka % 2 {
	case 0:
		fmt.Println("Genap")
	default:
		fmt.Println("Ganjil")
	}
}
