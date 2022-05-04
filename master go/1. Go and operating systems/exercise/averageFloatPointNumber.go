package main

import(
	"fmt"
	"os"
	"strconv"
)

func main(){
	var floatAverageNumber float64
	arguments := os.Args

	if len(arguments) == 1 {
		fmt.Println("Very small number of arguments, please input 2 or more")
	} else {
		for i := 0; i < len(arguments); i++ {
			floatNumber, _ := strconv.ParseFloat(arguments[i], 64)
			floatAverageNumber = floatAverageNumber + floatNumber
		}
	}
	floatAverageNumber = floatAverageNumber / float64(len(arguments))
	fmt.Println("Float average number is:", floatAverageNumber)

}