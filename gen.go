package main

//go:generate listgen -package intlist -list-type IntList -value-type=int -cmp "return a-b" -out intlist/intlist.go
//go:generate listgen -package stringlist -list-type StringList -value-type=string -cmp "if a < b {return -1}; if a > b {return 1}; return 0;" -out stringlist/stringlist.go
//go:generate gofmt -w intlist/intlist.go
//go:generate gofmt -w stringlist/stringlist.go
