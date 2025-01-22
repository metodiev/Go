package main

import "fmt"

func main(){
	switch v:=param.(type)
default:
	fmt.Printf("Unexpected type %T", v)
case uint64:
	fmt.Println("Integer type")
case string:
	smt.Println("String type")
}