package main

import (
	"fmt"
	"strings"
)

//Main Function
func main(){
	//Create and initialize
	string1:="Welcome to the program"
	string2:="Golang Program"

	//Check for presence Using Contains method of strings package
	res1:=strings.Contains(string1, "Welcome")
	res2:=strings.Contains(string2, "Program")

	//Display the result
	fmt.Println("Is Welcome word present in string1:", res1)
	fmt.Println("Is Program word presejted int string2", res2)



}