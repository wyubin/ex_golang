package easyother

func generate(numRows int) [][]int {
	res := make([][]int, numRows)
	// init lv
	res[0] = []int{1}
	// generate lv
	for lv := 1; lv < numRows; lv++ {
		preLv := lv - 1
		res[lv] = make([]int, lv+1)
		res[lv][0], res[lv][lv] = 1, 1
		for j := 1; j < lv; j++ {
			res[lv][j] = res[preLv][j-1] + res[preLv][j]
		}
	}
	return res
}
