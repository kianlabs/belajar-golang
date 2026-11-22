package main

import "fmt"

func main() {
	fmt.Println(isGenap(7))
	fmt.Println(isGenap(8))
}

func isGenap(angka int) bool {
	return angka%2 == 0
}
