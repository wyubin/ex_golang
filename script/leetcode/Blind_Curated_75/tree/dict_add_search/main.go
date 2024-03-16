package main

import "fmt"

func main() {
	wordDict := NewTrie()
	for _, word := range []string{"bad", "dad", "mad", "bada"} {
		wordDict.Insert(word)
	}
	for _, wordSearch := range []string{"pad", "bad", ".ad", "b..", "b.", "b..."} {
		k := wordDict.Search(wordSearch)
		fmt.Printf("search[%s]: %+v\n", wordSearch, k)
	}
}
