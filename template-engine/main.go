// go tiene su propio motor de plantillas ajaja <3 cool
// ver tuto base https://www.youtube.com/watch?v=dWchPTi9Vc0
package main

import (
	"fmt"
	"html/template"
	"os"
	"strings"
)

func main() {
	//qué significa el . ?
	//en otros lenguajes el punto sirve para acceder a los métodos del objeto que estamos llamando
	//podríamos compararlo al uso de this
	//se refiere al objeto altual que estemos llamando en ese scope

	//este . estaría a un scope general correspondiente a la variable say
	templateString := "Hey, {{.}}"

	say := "Baby"

	tmpl := template.Must(template.New("Hey").Parse(templateString))
	//para cumplir con la firma le pasamos una interfaz io.Writer para que escriba la salida
	err := tmpl.Execute(os.Stdout, say)

	if err != nil {
		fmt.Println(err.Error())
	}
	//en este caso el . hace referencia al número 1
	templateString = "{{ if . }}Hey{{ end }}"
	tmpl = template.Must(template.New("Baby Hey baby hey\n").Parse(templateString))
	_ = tmpl.Execute(os.Stdout, 1)

	names := []string{"Tanjiro", "Oboro-chan", "Sho"}
	//iterar -> range
	//en este contexto el . es lo que sea que se este pasando a nuestra plantilla, es decir cada name de names
	//si estuvieramos pasando como data a la plantilla estructuras y quisieramos llamar a uno de sus propiedades
	//o fields sería el .nombreattr para referenciarlo
	templateString = "{{ range . }}Hey, {{.}}\n{{ end }}"
	tmpl = template.Must(template.New("range").Parse(templateString))
	_ = tmpl.Execute(os.Stdout, names)

	//se puede inyectar incluso funcionalidad que normalmente no estaba antes de procesar la plantilla
	//usando el tipo FuncMap
	//func(*Template)Funcs(FuncMap FuncMap) *Template
	//Funcs agrega los elementos del argumento map a la funcion map del template.
	//debe ser llamada antes de que el template sea parseado
	templateString = "{{ range . }}Hey, {{ toUpper . }}\n{{ end }}"
	funcMap := map[string]interface{}{
		"toUpper": strings.ToUpper,
	}
	//ahora imprime todo en mayuscula
	//esto es super poderoso para crear proyectos con distintas funcionalidades
	tmpl = template.Must(template.New("funcs").Funcs(funcMap).Parse(templateString))
	_ = tmpl.Execute(os.Stdout, names)
}
