package main

import (
	"autheticate_project/logger"
	"fmt"
	
)

func main(){
	fmt.Println("Starting the Program...")
	logger.SetLogLevel(logger.DEBUG)
	
}