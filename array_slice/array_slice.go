package main

import "fmt"

func Sum(numbers []int) (sum int) {
	for _, number := range numbers {
		sum += number
	}

	return
}

func SumAll(numbersToSum ...[]int) (result []int) {
	result = make([]int, len(numbersToSum))

	for index, numbers := range numbersToSum {
		result[index] = Sum(numbers)
	}

	return
}

func SumAllTails(numbersToSum ...[]int) (result []int) {
	result = make([]int, len(numbersToSum))

	for index, numbers := range numbersToSum {
		if len(numbers) <= 1 {
			result[index] = 0
			continue
		}
		result[index] = Sum(numbers[1:])
	}

	return
}

func main() {
	numbers := [4]int{1, 3, 5, 2}

	fmt.Println(Sum(numbers[:]))
}
