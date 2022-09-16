package main

import (
	"fmt"
	"os"
)

func testCall() {
	for i := 0; i < 100; i++ {
		fmt.Println(i)
		if i == 10 {
			os.Exit(1)
		}
	}
}

func main() {
	defer fmt.Println("!")
	 fmt.Println("!")
	 fmt.Println("!")
	 testCall()

	//we use os.Exit  to imidiately exit wit a given status.
	//os.Exit(3)
	os.Exit(1)
}
