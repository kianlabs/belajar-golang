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
	tests := []struct {
		nama     string
		a, b     int
		mau      int
		mauError bool
	}{
		{"bagi normal", 10, 2, 5, false},
		{"bagi nol", 10, 0, 0, true},
		{"bagi habis", 9, 3, 3, false},
	}

	for _, tc := range tests {
		hasil, err := Bagi(tc.a, tc.b)
		if tc.mauError {
			if err == nil {
				fmt.Println(tc.nama, "GAGAL: harusnya error")
			} else {
				fmt.Println(tc.nama, "PASS")
			}
		} else if err != nil {
			fmt.Println(tc.nama, "GAGAL: tidak boleh error")
		} else if hasil != tc.mau {
			fmt.Println(tc.nama, "GAGAL: dapat", hasil, "mau", tc.mau)
		} else {
			fmt.Println(tc.nama, "PASS")
		}
	}
}
