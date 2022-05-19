package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	EVEN = iota
	ODD
)

func print(oddOrEven int, words string) {
	for i := oddOrEven; i < len(words); i += 2 {
		fmt.Printf("%c", words[i])
	}
	fmt.Print(" ")
}

func main() {
	var T int

	fmt.Scan(&T)
	words := make([]string, T)
	reader := bufio.NewReader(os.Stdin)
	for i := 0; i < T; i++ {
		words[i], _ = reader.ReadString('\n')
		words[i] = strings.TrimSuffix(words[i], "\n")
	}

	for word := range words {
		print(EVEN, words[word])
		print(ODD, words[word])

		fmt.Printf("\n")
	}
}
