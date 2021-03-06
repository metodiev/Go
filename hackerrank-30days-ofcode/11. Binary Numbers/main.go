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

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	n := int32(nTemp)
	binaryNumbers(n)
}

func binaryNumbers(n int32){
	var  reminder, count, max int32
	
	for n > 0 {
		reminder = n % 2
		if reminder == 1 {
			count++
		} else {
			count = 0
		}
		if count > max {
			max = count
		}
		n /=2
	}
	fmt.Println(max)
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
