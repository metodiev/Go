package main

import (
	"fmt"
	"strings"
)

/*
 * Complete the 'twoStrings' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts following parameters:
 *  1. STRING s1
 *  2. STRING s2
 */

func twoStrings(s1 string, s2 string) string {
	var strContins string
	
	strContins = "NO"
	for i := 0; i < len(s1); i++ {
		//substring := s1[i:2]
		substring := string(s1[i])
		//fmt.Println(string(s1[i]))
		//fmt.Println(substring)
		if strings.Contains(s2, substring) {
			strContins = "YES"
			break
		}
	}
	if strContins == "NO" {
	for i := len(s1); i > 0; i-- {
		substring := s1[:i]
		//fmt.Println(substring)
		if strings.Contains(s2, substring) {
			strContins = "YES"
			break
		}
	}
 }
	return strContins
}

func main() {

	s1 := "wouldyoulikefries"
	s2 := "abcabcabcabcabcabc"

	str := twoStrings(s1, s2)
	fmt.Println(str)
}
