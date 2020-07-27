package main

import (
	"fmt"
)

func main() {

	//Duck type
	arreglo := crearArreglo()
	fmt.Printf("Arreglo %+v\n", arreglo)

	slice := crearSlice()
	fmt.Printf("Slice %+v\n", slice)

	sliceCopiado := make([]int, len(slice))

	//copy no devuelve un nuevo array sino que reemplaza los valores
	numeroElementosCopiados := copy(sliceCopiado[3:], slice[2:])
	fmt.Printf("Slice Copiado %+v\n", sliceCopiado)
	fmt.Printf("Elementos Copiados %d\n", numeroElementosCopiados)

	//slice literal
	sliceLiteral := []string{"Hola", "Chicos", ":)"}
	fmt.Printf("Slice Literal %+v\n", sliceLiteral)

	sliceDeArray := arreglo[1:4]
	fmt.Printf("Slice de Array %+v\n", sliceDeArray)

	sliceDeSlice := slice[2:4]
	fmt.Printf("Slice de Slice %+v\n", sliceDeSlice)

	sliceElementoEliminado := eliminarElementoSlide(1, sliceCopiado)
	/*	sliceElementoEliminado = eliminarElementoSlide(2, sliceElementoEliminado)
		sliceElementoEliminado = eliminarElementoSlide(3, sliceElementoEliminado)
		sliceElementoEliminado = eliminarElementoSlide(4, sliceElementoEliminado)*/
	fmt.Printf("Slice index eliminado %d\n", sliceElementoEliminado)
}

func crearArreglo() [5]int {
	var arreglo [5]int

	for i := 0; i < len(arreglo); i++ {
		arreglo[i] = i
	}

	return arreglo
}

func crearSlice() []int {

	var slice []int

	slice = make([]int, 0)

	for i := 0; i <= 10; i++ {
		slice = append(slice, i)
	}

	return slice
}

func eliminarElementoSlide(index int, slice []int) []int {
	newslice := make([]int, len(slice)-1)
	newslice = append(slice[:index], slice[index+1:]...)
	return newslice
}
