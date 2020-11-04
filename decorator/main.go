package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	/* 	c := Circle{5}
	   	PrintPerimeter(c) */

	cli := Decorate(http.DefaultClient,
		Authorization("TOKEN"),
		LoadBalancing(RoundRobin(0, "web1", "web2")),
		Logging(log.New(os.Stdout, "client: ", log.LstdFlags)),
		Instrumentation(
			NewCounter("client.request"),
			NewHistogram("client.latency", 0, 10e9, 3, 50),
		),
		FaultTolerance(5, time.Second),
	)

	reader := bytes.NewReader([]byte(`{"holo":"holo"}`))
	req, _ := http.NewRequest("url", "GET", reader)
	res, _ := cli.Do(req)
	fmt.Println(res)
}

/*
type Boundary interface {
	Perimeter() float64
}

func (c Circle) Perimeter() float64 {
	return 2 * c.Radius * 3.14
}

func PrintPerimeter(b Boundary) {
	fmt.Printf("%.2f", b.Perimeter())
}

type Circle struct {
	Radius float64
}
*/
