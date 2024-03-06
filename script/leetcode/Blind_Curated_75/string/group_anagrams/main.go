package main

import (
	"fmt"
	"strings"

	"slices"
)

func sortStr(s string) string {
	sortedChars := make([]string, len(s))
	for _, char := range s {
		sortedChars = append(sortedChars, string(char))
	}
	slices.Sort(sortedChars)
	return strings.Join(sortedChars, "")
}

func groupAnagrams(strs []string) [][]string {
	res := map[string][]string{}
	for _, strTmp := range strs {
		keyTmp := sortStr(strTmp)
		existAnas, found := res[keyTmp]
		if found {
			existAnas = append(existAnas, strTmp)
		} else {
			existAnas = []string{strTmp}
		}
		res[keyTmp] = existAnas
	}

	resList := [][]string{}
	for _, anas := range res {
		resList = append(resList, anas)
	}
	return resList
}

func main() {

	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Printf("k: %s\n", groupAnagrams(strs))
}
