package main

import "fmt"

func isAnagram(s string, t string) bool {
	hash1 := map[byte]int{}
	for _, val := range []byte(s) {
		hash1[val] += 1
	}
	for _, val := range []byte(t) {
		exTime, exists := hash1[val]
		if !exists {
			return false
		}
		if exTime == 0 {
			return false
		} else if exTime == 1 {
			delete(hash1, val)
		} else {
			hash1[val] -= 1
		}
	}
	return len(hash1) == 0
}

func main() {
	s := "anagram"
	t := "nagaram"
	k := isAnagram(s, t)
	fmt.Printf("k:%+v\n", k)
}
