package main

import (
	"fmt"
)

/*
 * Complete the 'isBalanced' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts STRING s as parameter.
 */
type Node struct {
	Value string
	Next  *Node
}

var size = 0
var stack = new(Node)

func Push(v string) bool {
	if stack == nil {
		stack = &Node{v, nil}
		size = 1
		return true
	}

	temp := &Node{v, nil}
	temp.Next = stack
	stack = temp
	size++
	return true
}

func Pop(t *Node) (string, bool) {
	if size == 0 {
		return "", false
	}
	if size == 1 {
		size = 0
		stack = nil
		return t.Value, true
	}
	stack = stack.Next
	size--

	return t.Value, true
}

func isBalanced(s string) string {
	stack = nil
	isBalancedBrackets := "YES"

	for i := 0; i < len(s); i++ {
		substring := string(s[i])
		if (substring == ")" || substring == "]" || substring == "}")  {
			previousString, _ := Pop(stack)

			if !isBalancedBracketString(substring, previousString) {
				isBalancedBrackets = "NO"
			}

		} else {
			//fmt.Println(substring)
			Push(substring)
		}

	}

	// if isBalancedBrackets == "YES" {
	// 	fmt.Println(size)
	// 	for i := size / 2; i < size; i++ {
	// 		substring := string(s[i])
	// 		v, _ := Pop(stack)
	// 		//fmt.Println("Comparing substring to v", substring, v)

	// 		if !isBalancedBracketString(substring, v) {
	// 			//fmt.Println("brackets are different")
	// 			isBalancedBrackets = "NO"
	// 		}

	// 	}
	// }
	 fmt.Println(isBalancedBrackets)
	return isBalancedBrackets

}

func isBalancedBracketString(substring, v string) bool {
	var isBalancedString bool

	if substring == ")" && v == "(" {
		isBalancedString = true
	} else if substring == "]" && v == "[" {
		isBalancedString = true
	} else if substring == "}" && v == "{" {
		isBalancedString = true
	} else {
		isBalancedString = false
	}

	return isBalancedString

}

func main() {

	// {[()]}          first s = '{[()]}'
	// {[(])}          second s = '{[(])}'
	// {{[[(())]]}}    third s ='{{[[(())]]}}'

	 isBalanced("{[()]}")
	 isBalanced("{[(])} ")
	 isBalanced("{{[[(())]]}}")
	isBalanced("{(([])[])[]}")
}
