package main

import (
    "fmt"
)

/*
 * Complete the 'alternatingCharacters' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts STRING s as parameter.
 */

func alternatingCharacters(s string) int32 {
    // Write your code here
	var deletingCharacters int32
	var firstCharachter string
	var characterFlag bool

	firstCharachter = string(s[0])
//	fmt.Println("first Character : ", firstCharachter)

	for i := 1; i < len(s); i++ {
		if firstCharachter == string(s[i]){
			characterFlag = true
		} else if firstCharachter != string(s[i]) {
			characterFlag = false
			firstCharachter = string(s[i])
		}

		if characterFlag {
			//fmt.Println("i, s[i] is equal to", i, s[i])
			deletingCharacters++
		}
		
	}
	//fmt.Println(deletingCharacters)

	return deletingCharacters


}

func main() {
   
	q := 5
	s := []string{"AAAA", "BBBBB", "ABABABAB", "BABABA", "AAABBB"}
	

	for i := 0; i < q; i++ {
		alternatingCharacters(s[i])	
	}

}
