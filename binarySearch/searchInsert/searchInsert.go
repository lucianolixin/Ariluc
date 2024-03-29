package searchInsert

func searchInsert(nums []int, target int) int {
	numsLen := len(nums)
	high := numsLen - 1
	low := 0

	for low <= high {
		mid := low + ((high - low) >> 2)
		if nums[mid] > target {
			high = mid - 1
		} else if nums[mid] < target {
			low = mid + 1
		} else {
			return mid
		}
	}
	return low
}
