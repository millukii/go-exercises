//cuando los datos son guardados en un linked list son almacenados en una
//secuencia de nodos enlazados
//cada nodo tiene la dirección del siguiente
//https://www.youtube.com/watch?v=1S0_-VxPLJo
package main

import "fmt"

//Node
type node struct {
	data int
	next *node
}
type linkedList struct {
	head   *node
	length int
}

func (l *linkedList) prepend(n *node) {
	second := l.head
	l.head = n
	l.head.next = second
	l.length++
}
func (l linkedList) printListData() {
	toPrint := l.head
	for l.length != 0 {
		fmt.Printf("\n%d", toPrint.data)
		toPrint = toPrint.next
		l.length--
	}
	fmt.Printf("\n")
}
func (l *linkedList) deleteWithValue(value int) {
	if l.length == 0 {
		return
	}
	if l.head.data == value {
		l.head = l.head.next
		l.length--
		return
	}
	previousToDelete := l.head
	for previousToDelete.next.data != value {
		if previousToDelete.next.next == nil {
			return
		}
		previousToDelete = previousToDelete.next
	}
	previousToDelete.next = previousToDelete.next.next
	l.length--
}

func main() {
	newList := linkedList{}

	node1 := &node{data: 1}
	node2 := &node{data: 2}
	node3 := &node{data: 3}
	node4 := &node{data: 35}
	node5 := &node{data: 34}
	newList.prepend(node1)
	newList.prepend(node2)
	newList.prepend(node3)
	newList.prepend(node4)
	newList.prepend(node5)

	fmt.Println(newList)
	newList.deleteWithValue(2)
	newList.printListData()
}
