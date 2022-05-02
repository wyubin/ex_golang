package easyother

func isHappy(n int) bool {
	if n == 0 {
		return false
	}
	if n == 1 {
		return true
	}
	arrRes := map[int]struct{}{}
	resTemp := n
	for resTemp != 1 {
		_, exists := arrRes[resTemp]
		if exists {
			return false
		}
		arrRes[resTemp] = struct{}{}
		resTemp = sumDigitSquare(resTemp)
	}
	return true
}

func sumDigitSquare(n int) int {
	ss := 0
	digit10 := n
	var residue int
	for digit10 != 0 {
		residue = digit10 % 10
		digit10 = digit10 / 10
		ss += residue * residue
	}
	return ss
}
