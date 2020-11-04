package main

import "net/http"

type Client interface {
	Do(*http.Request) (*http.Response, error)
}

type ClientFunc func(r *http.Request) (*http.Response, error)

func (f ClientFunc) Do(r *http.Request) (*http.Response, error) {
	return f(r)
}
