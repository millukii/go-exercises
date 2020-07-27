package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {

	solutionA()
	solutionB()
}

func solutionA() {
	start := time.Now()

	_, err := http.Get("https://fb.com")

	if err != nil {
		fmt.Println(err.Error())
	}

	elapsed := time.Since(start).Seconds()
	fmt.Println(elapsed)
}

func solutionB() {
	start := time.Now()
	_, err := http.Get("https://fb.com")

	if err != nil {
		fmt.Println(err.Error())
	}
	elapsed := time.Now().Sub(start)
	fmt.Println(elapsed)
}
