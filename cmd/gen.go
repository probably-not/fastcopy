package main

import (
	"fmt"
	"os"
	"strings"
	"text/template"
)

type fastcopy struct {
	T                  string
	MaxL               int
	CopyFuncGenerators []copyFuncGenerator
}

type copyFuncMapper struct {
	T string
	L int
}

type copyFuncGenerator struct {
	N int
	T string
}

const maxN = 1024 // Generate functions for copying up until 1 kilo.

func main() {
	b, err := os.ReadFile("fastcopy.go.tmpl")
	if err != nil {
		panic(err)
	}

	t, err := template.New("config").Funcs(template.FuncMap{
		"Title": strings.Title,
		"Iterate": func(count int) []int {
			iter := make([]int, count)
			for i := 0; i < count; i++ {
				iter[i] = i
			}
			return iter
		},
	}).Parse(string(b))
	if err != nil {
		panic(err)
	}

	for _, v := range []string{
		"bool", "string", "int", "int8", "int16", "int32", "int64", "uint", "uint8", "uint16", "uint32",
		"uint64", "uintptr", "byte", "rune", "float32", "float64", "complex64", "complex128"} {
		data := fastcopy{
			T:                  v,
			MaxL:               maxN + 1, // Since we use it as the array size it needs to be plus one (I think?)
			CopyFuncGenerators: make([]copyFuncGenerator, 0),
		}

		for i := 0; i <= maxN; i++ {
			data.CopyFuncGenerators = append(data.CopyFuncGenerators, copyFuncGenerator{i, v})
		}

		file, err := os.Create(fmt.Sprintf("fastcopy.%s.gen.go", v))
		if err != nil {
			panic(err)
		}
		err = t.Execute(file, data)
		if err != nil {
			panic(err)
		}
		fmt.Println("Generated", v, "Fastcopy Functions!")
	}

	fmt.Println("Finished!")
}
