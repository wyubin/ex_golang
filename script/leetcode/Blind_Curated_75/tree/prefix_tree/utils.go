package main

type Trie struct {
	child  map[rune]*Trie
	isWord bool
}

func (s *Trie) Insert(inWord string) {
	currTrie := s
	for _, tmpRune := range inWord {
		nextTrie, found := currTrie.child[tmpRune]
		if !found {
			nextTrie = NewTrie()
			currTrie.child[tmpRune] = nextTrie
		}
		currTrie = nextTrie
	}
	currTrie.isWord = true
}

// search trie for complete word
func (s *Trie) HasWord(searchWord string) bool {
	return s.checkPrefix(searchWord, true)
}

func (s *Trie) checkPrefix(inString string, isWord bool) bool {
	currNode := s
	for _, tmpRune := range inString {
		nextNode, found := currNode.child[tmpRune]
		if !found {
			return false
		}
		currNode = nextNode
	}
	return currNode.isWord == isWord
}

func (s *Trie) HasPrefix(searchWord string) bool {
	return s.checkPrefix(searchWord, false)
}

func NewTrie() *Trie {
	res := Trie{child: map[rune]*Trie{}}
	return &res
}
