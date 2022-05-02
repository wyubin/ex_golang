package easyother

import (
	"math"
)

func titleToNumber(columnTitle string) int {
	listAlpha := []byte(" ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	alpha2Count := map[byte]int{}
	for k, v := range listAlpha {
		alpha2Count[v] = k
	}
	bytesCol := []byte(columnTitle)
	lastIndexCol := len(bytesCol) - 1
	res := 0
	for k, v := range bytesCol {
		res += alpha2Count[v] * int(math.Pow(26, float64(lastIndexCol-k)))
	}
	return res
}
