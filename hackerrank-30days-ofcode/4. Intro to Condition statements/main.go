package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	NTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	N := int32(NTemp)
	printConditionActions(N)
}


func printConditionActions(number int32){

	if number % 2 != 0 {
		fmt.Println("Weird")
	} else if number % 2 == 0  {
		if number >= 2 && number < 6 {
			fmt.Println("Not Weird")
		} else if  number >= 6 && number <21 {
			fmt.Println("Weird")
		}else if  number >= 20 {
			fmt.Println("Not Weird")
		}
	} 
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
