package main

import (
	"bytes"
	"fmt"
	"go/format"
	"text/template"

	"github.com/prometheus/common/log"
)

type Data struct {
	Name   string
	Desc   string
	Fields []Field
}

type Field struct {
	Name     string
	TypeName string
}

func main() {
	data := []Data{
		{
			Name: "Server",
			Desc: "Details about a server",
			Fields: []Field{
				{"Id", "int"},
				{"Url", "string"},
				{"Active", "boolean"},
			},
		},
		{
			Name: "Channel",
			Desc: "A channel to chat in the server",
			Fields: []Field{
				{"Id", "int"},
				{"Name", "string"},
				{"Private", "boolean"},
			},
		},
		{
			Name: "Role",
			Desc: "Used to group permissions for users",
			Fields: []Field{
				{"Id", "int"},
				{"Name", "string"},
				{"Color", "string"},
			},
		},
	}

	tmpl, err := template.ParseFiles("struct.tmpl")
	if err != nil {
		log.Fatalf("Could not parse struct template: %v\n", err)
	}
	var processed bytes.Buffer
	err = tmpl.Execute(&processed, data)
	if err != nil {
		log.Fatalf("unable to parse data into template: %v\n", err)
	}
	formatted, err := format.Source(processed.Bytes())
	if err != nil {
		log.Fatalf("Could not format processed template: %v\n", err)
	}
	fmt.Println(string(formatted))

}
