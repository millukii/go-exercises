//https://www.youtube.com/watch?v=-oYitelECuQ
//muy útil para manejar grandes cantidades de datos
package main

import "fmt"

var count int

//Node representa el componente de una busqueda de arbol binario
type Node struct {
	Key   int
	Left  *Node
	Right *Node
}

//Insert va a agregar un nodo al arbol
//la key que se va a agregar no debe estar ya en el arbol
func (n *Node) Insert(k int) {
	//la idea es comparar la k paametro con la key del receiver e ir
	//al child en la derecha si es más largo o ir al child en la izquierda si es más pequeño
	if n.Key < k {
		//move right
		//si el node está vacío
		if n.Right == nil {
			n.Right = &Node{Key: k}
		} else {
			//si el nodo ya está ocupado vamos a repetir el llamado a insert a partir de ese nodo
			//llamado recursivo
			n.Right.Insert(k)
		}

	} else if n.Key > k {
		//move left
		if n.Left == nil {
			n.Left = &Node{Key: k}
		} else {
			n.Left.Insert(k)
		}
	}
}

//Search recibe una llave y devuelve true si existe un nodo con ese valor
func (n *Node) Search(k int) bool {
	count++
	if n == nil {
		return false
	}
	if n.Key < k {
		return n.Right.Search(k)
	} else if n.Key > k {
		return n.Left.Search(k)
	}
	return true
}
func main() {
	tree := &Node{Key: 100}
	tree.Insert(200)
	tree.Insert(300)
	fmt.Println(tree)

	tree.Insert(39)
	tree.Insert(23)
	tree.Insert(111)
	tree.Insert(492)
	tree.Insert(2)
	tree.Insert(34)
	tree.Insert(89)
	tree.Insert(344)
	tree.Insert(245)
	tree.Insert(357)

	fmt.Println(tree.Search(357))
	fmt.Println(count)
}
