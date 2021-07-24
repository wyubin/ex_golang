package main

import "fmt"

func firstUniqChar(s string) int {
	indFirst := map[byte]int{}
	indDup := map[byte]struct{}{}
	arrByte := []byte(s)
	lenKeep := 0
	lenRm := 0
	for lenKeep != len(arrByte) {
		chrTarget := arrByte[lenKeep]
		_, exFirst := indFirst[chrTarget]
		_, exDup := indDup[chrTarget]
		if exFirst || exDup {
			arrByte = append(arrByte[:lenKeep], arrByte[(lenKeep+1):]...)
			lenRm++
			if exFirst {
				delete(indFirst, chrTarget)
				indDup[chrTarget] = struct{}{}
			}
			continue
		}
		indFirst[chrTarget] = lenKeep + lenRm
		lenKeep++
	}
	for _, val := range arrByte {
		ind, exists := indFirst[val]
		if exists {
			return ind
		}
	}
	return -1
}

func main() {
	nums := "aadadaad"
	k := firstUniqChar(nums)
	fmt.Printf("k:%d\n", k)
}
