package main

import (
	"net/http"
	"time"
)

//FaultTolerance returns a Decorator that extends a Client with fault tolerance
//configured with the given attemps and backoff duration
func FaultTolerance(attemps int, backoff time.Duration) Decorator {
	return func(c Client) Client {
		return ClientFunc(func(r *http.Request) (res *http.Response, err error) {
			for i := 0; i <= attemps; i++ {
				if res, err = c.Do(r); err == nil {
					break
				}
				time.Sleep(backoff * time.Duration(i))
				break
			}
			return res, err

		})
	}
}
