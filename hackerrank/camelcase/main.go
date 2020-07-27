// go run main.go < camel.in > camel.out
package main

import (
	"fmt"
	"strings"
)

func main() {

	var input string

	fmt.Scanf("%s\n", &input)

	fmt.Println(solutionA(input))
	fmt.Println(solutionB(input))
}

func solutionA(input string) int {
	answer := 1
	//cómo podemos iterar un string?
	//un string detrás de scenas es solo un slice de bytes y un slice lo podemos iterar
	//facilmente con un for range

	for _, ch := range input {
		//imprime el byte de cada index
		//cada ch es una runa, basada en utf8-encoding
		//Los literales de runa son solo valores enteros de 32 bits
		//(sin embargo, son constantes sin tipo, por lo que su tipo puede cambiar)
		//Representan puntos de código de Unicode. Por ejemplo, la runa literal 'a' es en realidad el número 97.

		min, max := 'A', 'Z' //runas tipo int32
		// si el byte está dentro del mínimo-maximo es una Mayúscula
		//entonces intentamos contar las letras según el número de mayúsculas
		if ch >= min && ch <= max {
			answer = answer + 1
		}
		//fmt.Println("index:", i, "rune", ch)
	}
	return answer
}

func solutionB(input string) int {
	answer := 1
	for _, ch := range input {
		str := string(ch)
		if strings.ToUpper(str) == str {
			answer = answer + 1
		}
	}
	return answer
}
