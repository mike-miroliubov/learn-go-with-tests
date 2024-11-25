package arrays

func ArraySum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}

	return sum
}

func MultiArraySum(numbers ...[]int) []int {
	var result []int
	for _, array := range numbers {
		result = append(result, ArraySum(array))
	}

	return result
}

func ArrayTailSum(arrays ...[]int) []int {
	var result []int

	for _, array := range arrays {
		if len(array) == 0 {
			result = append(result, 0)
		} else {
			result = append(result, ArraySum(array[1:]))
		}
	}

	return result
}
