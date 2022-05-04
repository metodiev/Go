package main

import (
	"fmt"
	"os"
	"strconv"
)

func main(){
	arguments := os.Args

	if len(arguments) == 1 {
		fmt.Println("please input more than one argument")
		os.Exit(1)
	} else {
		for i := 0; i < len(arguments); i++ {
			//integerNumber
		}
	}
}