package main

import (
	"html/template"
	"os"
)

func main() {
	t, err := template.ParseFiles("hello.gohtml")

	if err != nil {
		panic(err)
	}

	data := struct {
		Name string
		Age  int
		Kids []string
		Phones map[string]int
	}{
		"Zuck",
		32,
		[]string{"John", "Luke", "Sara"},
		map[string]int{"home": 333222, "mobile": 222333},
	}

	if err = t.Execute(os.Stdout, data); err != nil {
		panic(err)
	}
}
