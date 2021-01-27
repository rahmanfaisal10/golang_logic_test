package test_soal

import "fmt"

func QuestionNumber1(arrayLong int) int {
	var data int
	fmt.Print("please enter the number 0 or 1: ")
	arrayData := make([]int, 0)
	for i := 0; i < arrayLong; i++ {
		fmt.Scan(&data)
		if data == 1 || data == 0 {
			arrayData = append(arrayData, data)
		} else {
			fmt.Println("the numbers allowed are only 1 or 0")
			return 0
		}
	}

	maxCount, currentCount := 0, 0
	for _, v := range arrayData {
		if v == 1 {
			currentCount++
		} else {
			currentCount = 0
		}
		if currentCount > maxCount {
			maxCount = currentCount
		}
	}
	return maxCount
}