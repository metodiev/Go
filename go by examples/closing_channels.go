package main

import (
	"fmt"
	"time"
)

func main()  {
	jobs := make(chan int, 500)
	done := make(chan bool)

	go func(){
		for {
			j, more := <-jobs
		if more {
			//Calling sleep method
			time.Sleep(1* time.Second)
				fmt.Println("received job",j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()
	for j := 1; j <= 300; j++ {
		jobs <-j
		fmt.Println("send job", j)
	}
	close(jobs)
	fmt.Println("sent all jobs")

	<-done
}