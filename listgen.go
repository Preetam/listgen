package main

import (
	"flag"
	"io"
	"log"
	"os"
)

type listDesc struct {
	Package         string
	ValueType       string
	CompareFunction string
}

func main() {
	outputFile := flag.String("out", "", "output file")
	packageName := flag.String("package", "", "package name")
	valueType := flag.String("type", "", "value type")
	cmpFunc := flag.String("cmp", "", "comparison function body")
	flag.Parse()

	if *valueType == "" || *cmpFunc == "" || *packageName == "" {
		log.Fatal("value type, comparison function body, and package name must be given")
	}

	var w io.Writer = os.Stdout

	if *outputFile != "" {
		f, err := os.Create(*outputFile)
		if err != nil {
			log.Fatal(err)
		}
		w = f
		defer f.Close()
	}

	sourceTempl.Execute(w, listDesc{
		Package:         *packageName,
		ValueType:       *valueType,
		CompareFunction: *cmpFunc,
	})
}
