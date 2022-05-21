package main

import (
	"fmt"
	"sort"
)

/*
 * Complete the 'activityNotifications' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER_ARRAY expenditure
 *  2. INTEGER d
 */

func activityNotifications(expenditure []int32, d int32) int32 {
	// Write your code here
	//order the slice
	var numberOfNotifications int32

	for i := 0; i < len(expenditure); i++ {
		var median int32
		//after the bank get enough transaction data
		if int32(i) >= d {
			medianArr := make([]int32, d)
			copy(medianArr, expenditure[:i])
            fmt.Println(medianArr)
			sort.Slice(medianArr, func(i, j int) bool { return medianArr[i] < medianArr[j] })

			if len(medianArr)%2 == 0 {
				median = (medianArr[len(medianArr)/2] + medianArr[(len(medianArr)/2)+1]) / 2

			} else if len(medianArr)%2 != 0 {
				median = int32(medianArr[len(medianArr)/2])
			}
			fmt.Println("coutner", i)
			fmt.Println(median)
			
		}
	}

	return numberOfNotifications
}

//2, 2, 3, 3, 4, 4, 5, 6, 8

func main() {
	// reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

	// stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	// checkError(err)

	// defer stdout.Close()

	// writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

	// firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	// nTemp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
	// checkError(err)
	// n := int32(nTemp)

	// dTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
	// checkError(err)
	// d := int32(dTemp)

	// expenditureTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	// var expenditure []int32

	// for i := 0; i < int(n); i++ {
	//     expenditureItemTemp, err := strconv.ParseInt(expenditureTemp[i], 10, 64)
	//     checkError(err)
	//     expenditureItem := int32(expenditureItemTemp)
	//     expenditure = append(expenditure, expenditureItem)
	// }

	expenditure := []int32{2, 3, 4, 2, 3, 6, 8, 4, 5}
	var d int32
	d = 5

	activityNotifications(expenditure, d)

	// fmt.Fprintf(writer, "%d\n", result)

	// writer.Flush()
}
