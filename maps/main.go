package main

import "fmt"

func main() {

	mapaStringString := CrearMapaStringString()
	fmt.Println("Paises participantes")
	Imprimir(mapaStringString)

	fmt.Println("Todos los Paises")
	for i, v := range Paises {
		fmt.Printf("%d País %+v Capital %+v\n", i+1, v.Nombre, v.Capital)
	}

	delete(mapaStringString, "Ecuador")
	fmt.Println("Actualización")
	Imprimir(mapaStringString)

}

func Imprimir(mapa map[string]string) {

	for i, v := range mapa {
		fmt.Printf("Key %+v : Value %+v\n", i, v)
	}
}

func CrearMapaStringString() map[string]string {

	var mapa map[string]string

	mapa = make(map[string]string, 0)

	for i := 0; i < len(Paises); i++ {
		if Paises[i].Participa == true {
			mapa[Paises[i].Nombre] = Paises[i].Capital
		}
	}
	return mapa
}

var Paises = []*Pais{
	&Pais{
		Nombre: "Ecuador", Capital: "Quito", Participa: true,
	},
	&Pais{
		Nombre: "Peru", Capital: "Lima", Participa: true,
	},
	&Pais{
		Nombre: "Yugoslavia", Capital: "Belgrado", Participa: false,
	},
}

type Pais struct {
	Nombre    string
	Capital   string
	Participa bool
}
