package main

//go:generate listgen -package intlist -type=int -cmp "return a-b" -out intlist/intlist.go
//go:generate listgen -package stringlist -type=string -cmp "if a < b {return -1}; if a > b {return 1}; return 0;" -out stringlist/stringlist.go
//go:generate gofmt -w intlist/intlist.go
//go:generate gofmt -w stringlist/stringlist.go
