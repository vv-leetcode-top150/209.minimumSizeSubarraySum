package forcycle

func MinSubArrayLen(target int, nums []int) int {
	numsLen := len(nums)
	minSubArrayLen := numsLen + 1
	for i := range nums {
		sum := 0
		subArrayLen := 0
		found := false
		for j := i; j < numsLen; j++ {
			sum += nums[j]
			subArrayLen++
			if sum >= target {
				found = true
				break
			}
		}
		if (subArrayLen < minSubArrayLen) && found {
			minSubArrayLen = subArrayLen
		}
	}
	if minSubArrayLen > numsLen {
		minSubArrayLen = 0
	}
	return minSubArrayLen
}
