package main

//Decorator wraps a Client with extra behaviour
type Decorator func(Client) Client

//Decorate decorates a Client c with all the given Decorators, in order
func Decorate(c Client, ds ...Decorator) Client {
	decorated := c
	for _, decorate := range ds {
		decorated = decorate(decorated)
	}
	return decorated
}
