package main

import (
	"fmt"
	"strings"
)

/*
 * Complete the 'makeAnagram' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. STRING a
 *  2. STRING b
 */

func makeAnagram(a string, b string) int32 {
	var deletingCharacters int32
	deletingCharacters = 0
	//var m map[string]string
	stringSet := map[string]struct{}{}

	for i := 0; i < len(b); i++ {
		if !strings.Contains(a, string(b[i])) {
			stringSet[string(b[i])] = struct{}{}
			fmt.Println(string(b[i]))
			deletingCharacters++
		}
	}

	for i := 0; i < len(a); i++ {
		if !strings.Contains(b, string(a[i])) {
			stringSet[string(a[i])] = struct{}{}
			fmt.Println(string(a[i]))
			deletingCharacters++
		}
	}

	// for i := 0; i < len(a); i++ {
	// 	char := string(a[i])
	// 	var validateChar int32
	// 	validateChar = 0
	// 	for j := 0; j < len(b); j++ {
	// 		if char == string(b[j]) {
	// 			//fmt.Println("Validate Char: ", validateChar)
	// 			validateChar++
	// 		}
	// 		if (j+1) == len(b) && validateChar == 0 {
	// 			stringSet[char] = struct{}{}
	// 			if !strings.Contains(b, string(b[j])) ||
	// 			   !strings.Contains(b, string(b[j])) {
	// 				deletingCharacters++
	// 			}
	// 			//deletingCharacters++
	// 			if char != string(b[j]) {
	// 				fmt.Println("a", char)
	// 				fmt.Println("b", string(b[j]))
	// 				fmt.Println(j)

	// 				//stringSet[string(b[j])] = struct{}{}
	// 				//m[char] = string(b[j])
	// 				//deletingCharacters = deletingCharacters + 2
	// 			}
	// 		}
	// 	}
	// }

	// for str := range stringSet {
	// 	//fmt.Println(str)
	// 	if strings.Contains(a, str) || strings.Contains(b, str) {
	// 		fmt.Println("delete string set" , str)
	// 		delete(stringSet, str)
	// 	}
	// }

	// for str := range stringSet {
	// 	fmt.Println(str)
	// 	//deletingCharacters++
	// }
	fmt.Println(deletingCharacters)

	return deletingCharacters
}

func main() {
	a := "jxwtrhvujlmrpdoqbisbwhmgpmeoke"
	b := "fcrxzwscanmligyxyvym"

	makeAnagram(a, b)
}
