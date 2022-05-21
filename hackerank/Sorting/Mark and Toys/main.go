package main

import (
	"fmt"
	"sort"
)

/*
 * Complete the 'maximumToys' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER_ARRAY prices
 *  2. INTEGER k
 */

func maximumToys(prices []int32, k int32) int32 {
	var maxToysNumbers int32
    sArr := make([]int32, len(prices))
    
    if len(prices) == 0 || k == 0 {
        return maxToysNumbers
    }

    copy(sArr, prices[:len(prices)])
    sort.Slice(sArr, func(i, j int) bool { return sArr[i] < sArr[j] })

	for i := 0; i < len(sArr); i++ {
         k = k - sArr[i]
         fmt.Println(sArr[i])
         if k < 0 {
             break
         }
         maxToysNumbers++
	}
	return maxToysNumbers
}

func main() {
	
    var k int32
    k = 50
    var prices []int32
    prices = []int32{1, 12, 5, 111, 200, 1000, 10}
    
	maximumToys(prices, k)

}


