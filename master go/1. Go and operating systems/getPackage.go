package main

import (
	"fmt"
	// you need to download manually this package
	//go get -v github.com/mactsouk/go/simpleGitHub
	// ls -l ~/go/src/github.com/mactsouk/go/simpleGitHub/
	"github.com/mactsouk/go/simpleGitHub"
)

func main () {
	fmt.Println(simpleGitHub.AddTwo(5, 6))
}