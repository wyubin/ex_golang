package main

import (
	"fmt"
	"slices"
)

type LRUCache struct {
	data     map[int]int
	order    []int
	capasity int
}

func (s *LRUCache) Get(key int) int {
	val, found := s.data[key]
	if !found {
		return -1
	}
	idx := slices.Index(s.order, key)
	s.order = append([]int{key}, append(s.order[:idx], s.order[idx+1:]...)...)
	return val
}

func (s *LRUCache) Put(key, val int) {
	if len(s.order) == s.capasity {
		var keyRemove int
		keyRemove, s.order = s.order[s.capasity-1], s.order[:(s.capasity-1)]
		delete(s.data, keyRemove)
	}
	s.order = append([]int{key}, s.order...)
	s.data[key] = val
}

func NewLRUCache(capasity int) *LRUCache {
	lru := LRUCache{
		data:     map[int]int{},
		order:    []int{},
		capasity: capasity,
	}
	return &lru
}

func main() {
	cache := NewLRUCache(2)
	cache.Put(1, 2)
	cache.Put(2, 3)
	fmt.Printf("Get[1]: %+v\n", cache.Get(1))
	cache.Put(3, 4)
	fmt.Printf("Get[2]: %+v\n", cache.Get(2))
}
