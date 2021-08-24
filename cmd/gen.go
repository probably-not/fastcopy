package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"text/template"
)

type fastcopy struct {
	T                  string
	MaxL               int
	MaxN               int
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

var (
	maxN            int
	typesToGenerate typesToGenerateList
	primitiveTypes  = typesToGenerateList{
		"bool", "string", "int", "int8", "int16", "int32", "int64", "uint", "uint8", "uint16", "uint32",
		"uint64", "uintptr", "byte", "rune", "float32", "float64", "complex64", "complex128",
	}
	showHelp bool
)

type typesToGenerateList []string

func (t *typesToGenerateList) Set(value string) error {
	*t = append(*t, strings.Split(value, ",")...)
	return nil
}

func (t *typesToGenerateList) String() string {
	return fmt.Sprint(*t)
}

func init() {
	flag.BoolVar(&showHelp, "help", false, "Show usage information.")
	flag.IntVar(&maxN, "maxn", 8192, "The highest length for which to generate a copy function. Higher values will result in higher memory consumption and binary sizes.")
}

func main() {
	flag.Var(&typesToGenerate, "types", "A comma-separated list of the types to generate fastcopy functions for. (default Go's basic primitive types)")
	flag.Parse()

	if len(typesToGenerate) == 0 {
		typesToGenerate = primitiveTypes
	}

	if showHelp {
		flag.Usage()
		os.Exit(0)
	}

	b, err := os.ReadFile("fastcopy_go1.17.go.tmpl")
	if err != nil {
		log.Fatalf("Unable to read template file fastcopy.go.tmpl with error: %v", err)
	}

	t, err := template.New("fastcopy_go1.17").Funcs(template.FuncMap{
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
		log.Fatalf("Unable to parse template file fastcopy_go1.17.go.tmpl with error: %v", err)
	}

	for _, v := range typesToGenerate {
		data := fastcopy{
			T:                  v,
			MaxL:               maxN + 1, // Since we use it as the array size it needs to be plus one (I think?)
			MaxN:               maxN,     // Since we use it as the array size it needs to be plus one (I think?)
			CopyFuncGenerators: make([]copyFuncGenerator, 0),
		}

		for i := 0; i <= maxN; i++ {
			data.CopyFuncGenerators = append(data.CopyFuncGenerators, copyFuncGenerator{i, v})
		}

		file, err := os.Create(fmt.Sprintf("fastcopy.%s.gen.go", v))
		if err != nil {
			log.Fatalf("Unable to create generated file fastcopy.%s.gen.go with error: %v", v, err)
		}
		err = t.Execute(file, data)
		if err != nil {
			log.Fatalf("Unable to execute template for generated file fastcopy.%s.gen.go with error: %v", v, err)
		}
		log.Println("Generated", v, "Fastcopy Functions!")
	}

	log.Println("Finished!")
}
