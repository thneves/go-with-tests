package main

func Sum(numbers []int) int {
	sum := 0

	for _, number := range numbers {
		sum += number
	}

	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	sizeOfSlicesToSum := len(numbersToSum)

	sums := make([]int, sizeOfSlicesToSum)

	for _, numbers := range numbersToSum {
		newSum := Sum(numbers)
		sums = append(sums, newSum)
	}

	return sums
}

func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int

	for _, numbers := range numbersToSum {
		tail := numbers[1:]
		sums = append(sums, Sum(tail))
	}

	return sums
}
