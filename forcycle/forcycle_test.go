package forcycle

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
	}

	for i, v := range tests {
		if MinSubArrayLen(v.target, v.nums) != v.expected {
			t.Errorf("failed %v", i)
		}
	}
}
