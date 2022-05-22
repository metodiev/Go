package main

import (
	"fmt"
	"sort"
	"strings"
)

/*
 * Complete the 'isValid' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts STRING s as parameter.
 */

func isValid(s string) string {
	var isValidString string
	charactersCounterMap := make(map[string]int)
	var maxElement int32
	remainingCharecter := 2
	var valueDifferentelements int32

	for i := 0; i < len(s); i++ {
		charactersCounterMap[string(s[i])] = strings.Count(s, string(s[i]))
	}

	//sort values
	values := make([]int ,0)
	for _, v := range charactersCounterMap{
		values = append(values, v)
	}
	sort.Ints(values)

	for _, v := range values {
		//fmt.Println(v)
		if maxElement != int32(v) {
			maxElement = int32(v)
			remainingCharecter--
			fmt.Println(maxElement)
			//fmt.Println(remainingCharecter)
		}
	}

	for _, element := range charactersCounterMap {
		//fmt.Println("Key:", key, "=>", "Element:", element)

		if element == 1 {
			valueDifferentelements++
			//fmt.Println("value of different elements: ", valueDifferentelements)
		}
		
	}

	if valueDifferentelements >=2 {
		//fmt.Println("Different element :", valueDifferentelements)
		remainingCharecter = remainingCharecter - int(valueDifferentelements)
	}
	if remainingCharecter >= 0 {
		isValidString = "YES"
	} else if remainingCharecter < 0 {
		isValidString = "NO"
	}

	fmt.Println(isValidString)
	return isValidString
}

func main() {
	inputString := "aaaabbcc"
	isValid(inputString)
}
