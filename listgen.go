package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

type listDesc struct {
	Package         string
	ListType        string
	ValueType       string
	CompareFunction string
}

func main() {
	outputFile := flag.String("out", "", "Output file. Leave blank for stdout.")
	packageName := flag.String("package", "", "Package name to use for the list.")
	listType := flag.String("list-type", "List", "List type.")
	valueType := flag.String("value-type", "", "Value type.")
	cmpFunc := flag.String("cmp", "", "Comparison function body. The argument names are `a' and `b'.")
	flag.Parse()

	exit := false

	if *valueType == "" {
		fmt.Println("Value type must be given using -type=[type].")
		exit = true
	}

	if *cmpFunc == "" {
		fmt.Println("Comparison function body must be given using -cmp='[body]'.")
		exit = true
	}

	if *packageName == "" {
		fmt.Println("Package name must be given using -package='[package]'.")
		exit = true
	}

	if exit {
		os.Exit(1)
	}

	var w io.Writer = os.Stdout

	if *outputFile != "" {
		f, err := os.Create(*outputFile)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		w = f
		defer f.Close()
	}

	sourceTempl.Execute(w, listDesc{
		Package:         *packageName,
		ListType:        *listType,
		ValueType:       *valueType,
		CompareFunction: *cmpFunc,
	})
}
