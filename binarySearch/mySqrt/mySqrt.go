package mySqrt

func mySqrt(x int) int {
	var low, high = 1, x
	for low <= high {
		mid := low + ((high - low) >> 1)
		res := mid * mid
		if res == x {
			return mid
		} else if res > x { //中点大于预期
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return low - 1
}
