package easyother

func mySqrt(x int) int {
	if x == 0 {
		return 0
	}
	if x == 1 {
		return 1
	}
	var res int
	// 如果1以上，最小是1，最大不會超過x/2
	low := 1
	high := x / 2

	for low <= high {
		mid := low + (high-low)/2

		if mid == x/mid {
			return mid
		}
		if mid < x/mid {
			// 因為要不大於結果的整數
			res = mid
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return res
}
