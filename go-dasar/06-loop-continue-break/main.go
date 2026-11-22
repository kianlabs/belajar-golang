package main

import "fmt"

func main() {
	for angka := 1; angka <= 10; angka++ {
		if angka%3 == 0 {
			continue
		}
		if angka == 7 {
			break
		}
		fmt.Println(angka)
	}
}
