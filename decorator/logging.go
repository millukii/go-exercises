package main

import (
	"log"
	"net/http"
)

// Logging returns a Decorator that logs a Client`s request
func Logging(l *log.Logger) Decorator {

	return func(c Client) Client {
		return ClientFunc(func(r *http.Request) (*http.Response, error) {
			l.Printf("%s: %s %s", r.UserAgent, r.Method, r.URL)
			return c.Do(r)
		})
	}
}

func LoggingDecorator(reqLogger HttpLogger) Decorator {
	return func(c Client) Client {
		return ClientFunc(func(req *http.Request) (*http.Response, error) {
			reqLogger(req)
			return c.Do(req)
		})
	}
}
