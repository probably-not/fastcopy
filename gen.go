package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"text/template"
)

type fastcopy struct {
	T                  string
	CopyFuncGenerators []copyFuncGenerator
	MaxL               int
	MaxN               int
}

type copyFuncMapper struct {
	T string
	L int
}

type copyFuncGenerator struct {
	T string
	N int
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

	go117Bytes, err := ioutil.ReadFile("fastcopy_go117.go.tmpl")
	if err != nil {
		log.Fatalf("Unable to read template file fastcopy_go1.17.go.tmpl with error: %v", err)
	}

	notGo117Bytes, err := ioutil.ReadFile("fastcopy_go117_nobuild.go.tmpl")
	if err != nil {
		log.Fatalf("Unable to read template file fastcopy.go.tmpl with error: %v", err)
	}

	templateFuncs := template.FuncMap{
		"Title": strings.Title,
		"Iterate": func(count int) []int {
			iter := make([]int, count)
			for i := 0; i < count; i++ {
				iter[i] = i
			}
			return iter
		},
	}

	go117Tmpl, err := template.New("fastcopy_go117").Funcs(templateFuncs).Parse(string(go117Bytes))
	if err != nil {
		log.Fatalf("Unable to parse template file fastcopy_go1.17.go.tmpl with error: %v", err)
	}

	notGo117Tmpl, err := template.New("fastcopy_go117_nobuild").Funcs(templateFuncs).Parse(string(notGo117Bytes))
	if err != nil {
		log.Fatalf("Unable to parse template file fastcopy.go.tmpl with error: %v", err)
	}

	for _, v := range typesToGenerate {
		data := fastcopy{
			T:                  v,
			MaxL:               maxN + 1, // Since we use it as the array size it needs to be plus one (I think?)
			MaxN:               maxN,     // Since we use it as the array size it needs to be plus one (I think?)
			CopyFuncGenerators: make([]copyFuncGenerator, 0),
		}

		for i := 0; i <= maxN; i++ {
			data.CopyFuncGenerators = append(data.CopyFuncGenerators, copyFuncGenerator{v, i})
		}

		err := os.MkdirAll(fmt.Sprintf("./%s", v), 0700)
		if err != nil {
			log.Fatalf("Unable to create directory for ./%s with error: %v", v, err)
		}

		go117File, err := os.Create(fmt.Sprintf("./%s/fastcopy_gen_go117.go", v))
		if err != nil {
			log.Fatalf("Unable to create generated file fastcopy_gen_go1.17.go for %v with error: %v", v, err)
		}
		err = go117Tmpl.Execute(go117File, data)
		if err != nil {
			log.Fatalf("Unable to execute template for generated file fastcopy_gen_go1.17.go for %v with error: %v", v, err)
		}

		notGo117File, err := os.Create(fmt.Sprintf("./%s/fastcopy_gen_go117_nobuild.go", v))
		if err != nil {
			log.Fatalf("Unable to create generated file fastcopy_gen_go117_nobuild.go for %v with error: %v", v, err)
		}
		err = notGo117Tmpl.Execute(notGo117File, data)
		if err != nil {
			log.Fatalf("Unable to execute template for generated file fastcopy_gen_go117_nobuild.go for %v with error: %v", v, err)
		}
		log.Println("Generated", v, "Fastcopy Functions!")
	}

	log.Println("Finished!")
}
