package main

import (
	"net/http"
	"time"
)

type Counter struct {
	Name  string
	Count int
}

type Histogram struct {
	Name string
	Vals []uint64
}

func (h Histogram) Observe(start int64) {

}
func (c Counter) Add(number int64) {

}
func NewCounter(name string) Counter {
	return Counter{
		Name: name,
	}
}

func NewHistogram(name string, v ...uint64) Histogram {
	var vals []uint64
	for _, v := range v {
		vals = append(vals, v)
	}
	return Histogram{
		Name: name,
		Vals: vals,
	}
}

//Instrumentation returns a Decorator that instruments a Client with the given metrics
func Instrumentation(request Counter, latency Histogram) Decorator {
	return func(c Client) Client {
		return ClientFunc(func(r *http.Request) (*http.Response, error) {
			//After finish
			defer func(start time.Time) {
				latency.Observe(time.Since(start).Nanoseconds())
				request.Add(1)
			}(time.Now())
			return c.Do(r)
		})
	}
}
