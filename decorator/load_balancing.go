package main

import (
	"encoding/base64"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"sync/atomic"
)

type BalanceStrategy func(...*url.URL) *url.URL

//Director modifies an http,Request to follow a load balancing strategy
type Director func(*http.Request)

//LoadBalancing returns a Decorator that load balances a Client`s request accoss
//multiple backends using the given Director
func LoadBalancing(dir Director) Decorator {
	return func(c Client) Client {
		return ClientFunc(func(r *http.Request) (*http.Response, error) {
			dir(r)
			return c.Do(r)
		})
	}
}

//RoundRobin returns a Balancer which round-robins across the given backends
func RoundRobin(robin uint64, backends ...string) Director {
	return func(r *http.Request) {
		if len(backends) > 0 {
			r.URL.Host = backends[atomic.AddUint64(&robin, 1)%uint64(len(backends))]
		}

	}
}

//Random returns a Balancer which randomly picks one of the given backends
func Random(seed int64, backends ...string) Director {
	rnd := rand.New(rand.NewSource(seed))
	return func(r *http.Request) {
		if len(backends) > 0 {
			r.URL.Host = backends[rnd.Intn(len(backends))]
		}
	}
}

func LoadBalancerDecorator(strategy BalanceStrategy, urls ...*url.URL) Decorator {
	return func(c Client) Client {
		return ClientFunc(func(req *http.Request) (*http.Response, error) {
			req.URL = strategy(urls...)
			return c.Do(req)
		})
	}
}

func RandomStrategy(urls ...*url.URL) *url.URL {
	idx := rand.Intn(len(urls))
	return urls[idx]
}

type HttpLogger func(req *http.Request)

func StandardHttpLogger(log *log.Logger) HttpLogger {
	return func(req *http.Request) {
		log.Printf("%s, %s", req.Method, req.URL)
	}
}

func VerboseHttpLogger(log *log.Logger) HttpLogger {
	return func(req *http.Request) {
		log.Printf("%s, %s %s %s", req.Method, req.URL, req.Body)
	}
}

func basicAuthHeader(username, password string) string {
	return "Basic " + base64Auth(username, password)
}

func base64Auth(username, password string) string {
	data := []byte(username + ":" + password)
	return base64.StdEncoding.EncodeToString(data)
}
