package main

import "fmt"

func main(){
//creating na array
arr:=[6]string{"This", "is", "a", "Go", "Program", "questions"}

//Print Array
fmt.Println("original array:", arr)

//Create a slice
slicedArr:=arr[1:4]

//Display slice
fmt.Println("Sliced Array:", slicedArr)

//Length of slice calculated using len()
fmt.Println("Length of the slice: %d", len (slicedArr))

//Capacity of slice calculated using cap()
fmt.Println("Capacity of the slice: %d", cap(slicedArr))
}