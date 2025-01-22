package main

import (
	"bytes"
	"fmt"
)

func main(){
	sl1:=[]byte{'P','R','O','G','R','A','M'}
	sl2:=[]byte{'B','I','T'}

	//Use Compare function to compare slices
	res:=bytes.Compare(sl1,sl2)

	if res==0 {
		fmt.Println("Equal Slices")
	} else{
		fmt.Println("Uniqual Slices")
	}
}