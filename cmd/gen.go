package main

import (
	"fmt"
	"os"
	"text/template"
)

type fastcopy struct {
	Copies []copy
}

type copy struct {
	N int
	T string
}

const maxN = 10

func main() {
	b, err := os.ReadFile("fastcopy.go.tmpl")
	if err != nil {
		panic(err)
	}

	t, err := template.New("config").Parse(string(b))
	if err != nil {
		panic(err)
	}

	file, err := os.Create("fastcopy.gen.go")
	if err != nil {
		panic(err)
	}

	data := fastcopy{make([]copy, 0)}
	for _, v := range []string{"byte", "string", "int", "float64"} {
		for i := 1; i <= maxN; i++ {
			data.Copies = append(data.Copies, copy{i, v})
		}
	}

	err = t.Execute(file, data)
	if err != nil {
		panic(err)
	}

	fmt.Println("Generated!")
}
