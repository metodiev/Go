package main

import (
	"fmt"
	"os"
	"strconv"
)
//https://gobyexample.com/number-parsing
func main(){
	var sumNumbers int
	arguments := os.Args
	if len(arguments) == 1{
		fmt.Println("Please provide more intger arguments")
		os.Exit(1)
	} else {
		for i := 0; i < len(arguments); i++ {
			number, _ :=  strconv.ParseInt(arguments[i], 0, 64)
			sumNumbers = sumNumbers + int(number)
		}
	}

	fmt.Println("The sum is :" , sumNumbers)
}