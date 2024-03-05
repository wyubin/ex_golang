package main

import "fmt"

type Stack struct {
	StringByte []byte
}

func (s *Stack) append(b byte) {
	s.StringByte = append(s.StringByte, b)
}

func (s *Stack) pop() byte {
	lastIndexByte := s.len() - 1
	lastByte := s.StringByte[lastIndexByte]
	s.StringByte = s.StringByte[0:lastIndexByte]
	return lastByte
}

func (s *Stack) len() int {
	return len(s.StringByte)
}

func isValid(s string) bool {
	byteStack := Stack{}
	front2End := map[byte]byte{'(': ')', '{': '}', '[': ']'}
	for _, val := range []byte(s) {
		// check current byte
		end, exists := front2End[val]
		if exists {
			byteStack.append(end)
		} else if byteStack.len() == 0 || byteStack.pop() != val {
			// stack 已經消化完 / 沒有對應在stack 的括號
			return false
		}
	}
	return byteStack.len() == 0
}

func main() {
	s := "()[]{}"
	k := isValid(s)
	fmt.Printf("k:%+v\n", k)
}
