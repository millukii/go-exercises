package main

import "fmt"

func main() {
	var length, delta int
	var input string

	fmt.Scanf("%d\n", &length)
	fmt.Scanf("%s\n", &input)
	fmt.Scanf("%d\n", &delta)

	/* 	fmt.Printf("lenght: %d\n", length)
	   	fmt.Printf("input: %s\n", input)
			 fmt.Printf("delta: %d\n", delta) */

	ascii_lowercase := []rune("abcdefghijklmnopqrstuvwxyz")

	newRune := rotate('z', 2, ascii_lowercase)
	fmt.Println(string(newRune))
}

func rotate(s rune, delta int, key []rune) rune {
	//no se encuentra
	index := -1

	for i, r := range key {
		if r == s {
			index = i
			break
		}

	}

	if index < 0 {
		panic("id < 0")
	}

	for i := 0; i < delta; i++ {
		index++
		if index >= len(key) {
			index = 0
		}
	}
	return key[index]
}
