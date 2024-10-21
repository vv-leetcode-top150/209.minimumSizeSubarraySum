package sliding

import "testing"

func TestMinSubArrayLen(t *testing.T) {
	tests := []struct {
		target   int
		nums     []int
		expected int
	}{
		{7, []int{2, 3, 1, 2, 4, 3}, 2},
		{4, []int{1, 4, 4}, 1},
		{11, []int{1, 1, 1, 1, 1, 1, 1, 1}, 0},
		{4, []int{2, 1, 2, 1}, 3},
		{15, []int{5, 1, 3, 5, 10, 7, 4, 9, 2, 8}, 2},
	}

	for i, values := range tests {
		if MinSubArrayLen(values.target, values.nums) != values.expected {
			t.Errorf("failed %v", i)
		}
	}
}
