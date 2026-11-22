package main

import "fmt"

func main() {
	total := 0
	angka := []int{2, 4, 6, 8, 10}
	for i, n := range angka {
		fmt.Println(i, n)
		total += n
	}
	fmt.Println(total)
}
