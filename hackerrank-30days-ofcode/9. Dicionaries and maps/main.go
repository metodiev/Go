package main

import (
	"fmt"
	"time"
)

func main() {
	n := 0
	fmt.Scanf("%d\n", &n)
	dict := make(map[string]string, n)
	for i := 0; i < n; i++ {
		var name, number string
		fmt.Scanf("%s %s\n", &name, & number)
		dict[name] = number
	}

	quit := make(chan struct{})
	queries := make([]string, 0)
	go func ()  {

		for {
			select{
			case <- quit:
				return
			default: 
			var line string
			n, _ := fmt.Scanf("%s\n", &line)
				if n > 0{
					queries = append(queries, line)
				}
			}
		}
		
	}()

	// Wait 1 second for all input
	time.Sleep(1 * time.Second)
	close(quit)

	//Print
	for _, v:= range queries {
		w, ok := dict[v]
		switch ok {
		case true:
			fmt.Printf("%s=%s\n",v , w)
		default:
			fmt.Println("Not found")
		}
	}


}
