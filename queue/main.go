package main

import "fmt"

// Queue struct representando una cola que contiene un slice
type Queue struct {
	items []int
}

//Enqueue agrega un valor al final
func (q *Queue) Enqueue(i int) {
	q.items = append(q.items, i)
}

//Dequeue quita un valor al principio y retorna el valor removido
func (q *Queue) Dequeue() int {
	toRemove := q.items[0]
	q.items = q.items[1:]
	return toRemove
}

func main() {
	newQueue := Queue{}
	fmt.Println(newQueue)
	newQueue.Enqueue(100)
	newQueue.Enqueue(200)
	newQueue.Enqueue(300)
	fmt.Println(newQueue)
	newQueue.Dequeue()
	fmt.Println(newQueue)
}
