package main

import (
	"fmt"
)

func reverseValue(a,b string) (string, string) {
return b,a
}

func main(){
	val1, val2 :=reverseValue("First value for a", "Second value for b")
	fmt.Println(val1, val2)
}
