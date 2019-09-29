package searchInsert

import (
	"testing"
)

func TestSearchInsert(t *testing.T) {
	nums := []int{1, 2, 3, 4, 9}
	t.Log(searchInsert(nums, 3))
}
