package easyother

func romanToInt(s string) int {
	byteS := []byte(s)
	romanLv := map[byte]int{'I': 0, 'V': 1, 'X': 2, 'L': 3, 'C': 4, 'D': 5, 'M': 6}
	romanCount := []int{1, 5, 10, 50, 100, 500, 1000}

	res := 0
	lastIndByte := len(byteS) - 1
	for ind, val := range byteS {
		lvCurr := romanLv[val]
		// check lv
		indPositive := true
		if ind != lastIndByte {
			lvNext := romanLv[byteS[ind+1]]
			if lvCurr < lvNext {
				indPositive = false
			}
		}
		// add
		if indPositive {
			res += romanCount[lvCurr]
		} else {
			res -= romanCount[lvCurr]
		}
	}
	return res
}
