package main

import "fmt"

func main() {
	trie := NewTrie()
	trie.Insert("abcde")
	fmt.Printf("check has abcde: %+v\n", trie.HasWord("abcde"))
	fmt.Printf("check has abc: %+v\n", trie.HasWord("abc"))
	fmt.Printf("check has abc prefix: %+v\n", trie.HasPrefix("abc"))
	fmt.Printf("check has abb prefix: %+v\n", trie.HasPrefix("abb"))
}
