package easyother

var pairEnd = map[byte]byte{'(': ')', '{': '}', '[': ']'}

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
	for _, val := range []byte(s) {
		// check first byte
		_, exists := pairEnd[val]
		if exists {
			byteStack.append(val)
		} else if byteStack.len() == 0 {
			// stack 已經消化完
			return false
		} else {
			validStart := byteStack.pop()
			if val != pairEnd[validStart] {
				return false
			}
		}
	}
	return byteStack.len() == 0
}
