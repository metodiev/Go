package main

import (
	"aPackage"
	"fmt"
)

func main() {
	fmt.Println("Using aPackage!")
	aPackage.A()
	aPackage.B()
	fmt.Println(aPackage.MyConstant)
}

// $ mkdir ~/go/src/aPackage
// $ cp aPackage.go ~/go/src/aPackage/
// $ go install aPackage
// $ cd ~/go/pkg/darwin_amd64/
// $ ls -l aPackage.a
// -rw-r--r-- 1 mtsouk staff 4980 Dec 22 06:12 aPackage.a

//go tool compile aPackage.go
