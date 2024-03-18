package main

import (
	"fmt"
	"slices"
	"strings"
)

func solution(words []string) string {
	res := ""
	charOrder := make([][]bool, 26)
	for i := 0; i < 26; i++ {
		charOrder[i] = make([]bool, 26)
	}
	visited := map[int]bool{}
	// record all used char
	for _, word := range words {
		for _, runeTmp := range word {
			intTmp := int(runeTmp - 'a')
			charOrder[intTmp][intTmp] = true
		}
	}
	// add ordered char
	for i := 0; i < len(words)-1; i++ {
		minLenWord := slices.Min([]int{len(words[i]), len(words[i+1])})
		var j int
		for j = 0; j < minLenWord; j++ {
			runeCurr, runeNext := rune(words[i][j]-'a'), rune(words[i+1][j]-'a')
			if runeCurr != runeNext {
				charOrder[int(runeCurr)][int(runeNext)] = true
				break
			}
		}
		if j == minLenWord && len(words[i]) > len(words[i+1]) {
			// if words[i] contains words[i+1]
			return ""
		}
	}
	// use dfs to update res
	for i := 0; i < 26; i++ {
		if !dfs(charOrder, i, visited, &res) {
			return ""
		}
	}
	return res
}

func dfs(charOrder [][]bool, intChar int, visited map[int]bool, res *string) bool {
	if !charOrder[intChar][intChar] {
		// early return if no used in dict
		return true
	}
	visited[intChar] = true
	// check prev char, need to pass rules below
	for i := 0; i < 26; i++ {
		if i == intChar || !charOrder[i][intChar] {
			continue
		}
		if visited[i] {
			return false
		}
		// try to retri prev char
		if !dfs(charOrder, i, visited, res) {
			return false
		}
	}
	visited[intChar] = false
	charOrder[intChar][intChar] = false
	*res = strings.Join([]string{*res, string(rune(intChar + 'a'))}, "")
	return true
}

func main() {
	s := []string{"wrt", "wrf", "er", "ett", "rftt"}
	k := solution(s)
	fmt.Printf("k: %+v\n", k)
}
