package sum

func Sum(numbers []int) int {
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	lenOfTheNumbers := len(numbersToSum)
	sums := make([]int, lenOfTheNumbers)

	for i, nums := range numbersToSum {
		sums[i] = Sum(nums)
	}
	return sums
}

func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int
	for _, nums := range numbersToSum {
		if len(nums) == 0 {
			sums = append(nums, 0)
		} else {
			tail := nums[1:]
			sums = append(sums, Sum(tail))
		}
	}
	return sums
}
