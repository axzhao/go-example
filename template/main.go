package main

import (
	"bytes"
	"fmt"
	"text/template"
)

type Type struct {
	Name string
	ID   int
}

var tmpl *template.Template

func main() {
	stu := Type{Name: "hello", ID: 11}
	str := "{{ . }} {{$.Name}} ID is {{ .ID }}"

	t, err := tmpl.Parse(str)
	if err != nil {
		panic(err)
	}
	var buf bytes.Buffer
	err = t.Execute(&buf, stu)
	if err != nil {
		panic(err)
	}
	fmt.Println(buf.String())
}

func init() {
	tmpl = template.New("test")
}
