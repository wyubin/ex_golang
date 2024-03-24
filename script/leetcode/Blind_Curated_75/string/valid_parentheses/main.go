package main

import "fmt"

var (
	pre2suf = map[byte]byte{
		'(': ')', '{': '}', '[': ']',
	}
)

func solution(inStr string) bool {
	stack := []byte{}
	for i := 0; i < len(inStr); i++ {
		currByte := inStr[i]
		suf, found := pre2suf[currByte]
		if found {
			stack = append(stack, suf)
		} else {
			lenStack := len(stack)
			if lenStack == 0 || currByte != stack[lenStack-1] {
				return false
			}
			stack = stack[:lenStack-1]
		}
	}
	return len(stack) == 0
}

func main() {
	s := "()[]{}"
	k := solution(s)
	fmt.Printf("k:%+v\n", k)
}
