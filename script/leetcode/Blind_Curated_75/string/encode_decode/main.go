package main

import (
	"fmt"
	"strconv"
	"strings"
)

func encode(inStrs []string) string {
	res := ""
	for _, inStr := range inStrs {
		res = fmt.Sprintf("%s%d/%s", res, len(inStr), inStr)
	}
	return res
}

func decode(dumpStr string) []string {
	res := []string{}
	for len(dumpStr) > 0 {
		idxSplit := strings.Index(dumpStr, "/")
		if idxSplit == -1 {
			panic("idxSplit need exist")
		}
		lenInt, err := strconv.Atoi(dumpStr[:idxSplit])
		if err != nil {
			panic(err)
		}
		idxNext := idxSplit + lenInt + 1
		res = append(res, dumpStr[idxSplit+1:idxNext])
		dumpStr = dumpStr[idxNext:]
	}
	return res
}

func main() {
	s := []string{"I", "am", "yubin"}
	strDump := encode(s)
	fmt.Printf("strDump: %+v\n", strDump)
	k := decode(strDump)
	fmt.Printf("k: %+v\n", k)
}
