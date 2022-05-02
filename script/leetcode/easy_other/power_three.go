package easyother

func isPowerOfThree(n int) bool {
	if n == 1 {
		return true
	}
	Final := n
	var residue int
	for Final != 1 {
		residue = Final % 3
		if residue != 0 {
			return false
		}
		Final = Final / 3
		if Final == 0 {
			return false
		}
	}
	return true
}
