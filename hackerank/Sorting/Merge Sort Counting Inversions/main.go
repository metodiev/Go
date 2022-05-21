package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

/*
 * Complete the 'countInversions' function below.
 *
 * The function is expected to return a LONG_INTEGER.
 * The function accepts INTEGER_ARRAY arr as parameter.
 */

func countInversions(arr []int32) int64 {
	var countInvertors int64
	var j int32
	j = 1
	for i := 0; i < len(arr); i++ {
		if arr[i] > arr[j] {
			countInvertors++
		}
		
	}

	return int64(countInvertors)
}

func main() {
  arr := []int32{2, 1, 3, 1, 2}
  countInversions(arr)
}
