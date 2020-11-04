package main

import (
	"log"
	"net/http"
	"time"
)

func RetryDecorator(attempts int, retryInterval time.Duration, logger *log.Logger) Decorator {
	return func(c Client) Client {
		return ClientFunc(func(req *http.Request) (res *http.Response, err error) {
			try := func(req *http.Request) (*http.Response, error) {
				return c.Do(req)
			}

			for attempts > 0 {
				if res, err = try(req); err == nil {
					return res, err
				}
				logger.Printf("retrying.. %+v", err)
				attempts--
				time.Sleep(retryInterval)
			}

			return res, err
		})
	}
}
