package main

import (
	"errors"
	"fmt"
)

func Bagi(a int, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("Tidak bisa bagi nol")
	}
	return a / b, nil
}

func main() {
	hasil, err := Bagi(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(hasil)
	}

	hasil2, err := Bagi(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(hasil2)
	}
}
