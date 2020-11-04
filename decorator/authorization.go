package main

import "net/http"

// Authorization return a Decorator that authorizes every Client request
// with the given token.
func Authorization(token string) Decorator {
	return Header("Authorization", token)
}

func Header(name, value string) Decorator {
	return func(c Client) Client {
		return ClientFunc(func(r *http.Request) (*http.Response, error) {
			r.Header.Add(name, value)
			return c.Do(r)
		})
	}
}
