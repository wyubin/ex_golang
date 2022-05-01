package easyother

func strStr(haystack string, needle string) int {
	hByte := []byte(haystack)
	nByte := []byte(needle)
	hLen := len(hByte)
	nLen := len(nByte)

	indStart := 0
	indEnd := nLen - 1
	for indEnd < hLen {
		indCurr := 0
		passCount := 0
		for indTemp, val := range nByte {
			indCurr = indStart + indTemp
			if hByte[indCurr] == val {
				passCount++
			}
		}
		if passCount == nLen {
			return indStart
		}
		// next
		indStart++
		indEnd++
	}

	return -1
}
