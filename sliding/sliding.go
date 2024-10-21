package sliding

func MinSubArrayLen(target int, nums []int) int {
	start, sum := 0, 0
	min := 0

	for end := range nums {
		sum += nums[end]

		for sum >= target {
			windowSize := (end - start) + 1
			if min == 0 || windowSize < min {
				min = windowSize
			}
			sum -= nums[start]
			start++
		}
	}
	return min
}
