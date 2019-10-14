package main

import (
	"testing"
)

func TestMoveZeros(t *testing.T) {
	nums := []int{1, 2, 3, 0, 9, 0, 4}
	MoveZeroS(nums)
	t.Log(nums)

}

func MoveZeroS(nums []int) {
	slow := 0
	//慢指针和之前的元素都是非0
	//迭代的指针为快指针，快慢指针指点的元素全部为0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 { // 当前元素不为0，慢指针向前移动
			nums[i], nums[slow] = nums[slow], nums[i]
			slow++
		}
	}
}
