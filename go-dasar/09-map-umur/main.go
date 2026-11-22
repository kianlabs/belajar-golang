package main

import "fmt"

func main() {
	umur := map[string]int{
		"Kyan":  22,
		"Putri": 19,
		"Jamal": 24,
	}
	fmt.Println(umur["Kyan"])
	umur["Eman"] = 25
	fmt.Println(umur["Eman"])
}
