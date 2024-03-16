package main

type Trie struct {
	child  map[int]*Trie
	isWord bool
}

func (s *Trie) Insert(inWord string) {
	var trieNext *Trie
	var found bool
	for _, tmpRune := range inWord {
		tmpInt := int(tmpRune - 'a')
		trieNext, found = s.child[tmpInt]
		if !found {
			trieNext = &Trie{}
			s.child[tmpInt] = trieNext
		}
	}
	trieNext.isWord = true
}

// search trie for complete word
func (s *Trie) HasWord(searchWord string) bool {
	return false
}

func NewTrie() *Trie {
	res := Trie{}
	return &res
}
