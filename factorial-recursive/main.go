package main

import "fmt"

func main() {
	fmt.Println(factorial(5, 1))
}

func factorial(n, m int) int {
	if n == 0 {
		return m
	}

	m *= n

	return factorial(n-1, m)
}
