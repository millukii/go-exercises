//https://www.youtube.com/watch?v=fsbm1FOSDJ0
//Stacks and queues are also very basic linear data structures.
//Los stacks funcionan bajo LIFO ultimo en entrar primero en salir
//se usan metodos push(agregar) and pop(quitar) ej pila de libros
//una queue funciona con FIFO primero en entrar primero en salir ej cola supermercado
//enqueue sería agregarse a la cola y dequeue salir de la cola
package main

import "fmt"

//la estructura stack representara el concepto que estamos tocando conteniendo un  slice
type Stack struct {
	items []int
}

//Push
func (s *Stack) Push(i int) {
	s.items = append(s.items, i)
}

//Pop remueve un valor al final y retorna el valor removido
func (s *Stack) Pop() int {
	l := len(s-item) - 1
	toRemove := s.items[l]
	s.items = s.items[:l]
	return toRemove
}

func main() {
	newStack := Stack{}
	fmt.Println(newStack)
	newStack.Push(100)
	newStack.Push(200)
	newStack.Push(300)
	fmt.Println(newStack)
	newStack.Pop()
	fmt.Println(newStack)
}
