package forcycle

func minSubArrayLen(target int, nums []int) int {
	numsLen := len(nums)
	minSubArrayLen := 0
	for i, num := range nums {
		sum := 0

		for j := i; j < numsLen; j++ {
			sum += nums[j]
			if sum >= target {
				break
			}
		}
	}
	return minSubArrayLen
}
