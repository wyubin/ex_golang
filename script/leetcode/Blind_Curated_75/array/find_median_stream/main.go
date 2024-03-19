package main

import (
	"fmt"
	"slices"
)

type MedianFinder struct {
	data []int
}

func (s *MedianFinder) AddNum(inData int) {
	s.data = append(s.data, inData)
}

func (s *MedianFinder) FindMedian() float32 {
	tmpData := s.data[:]
	slices.Sort(tmpData)
	idxHalf, hasRemider := len(tmpData)/2, len(tmpData)%2 == 1
	if hasRemider {
		return float32(tmpData[idxHalf])
	}
	return float32(tmpData[idxHalf-1]+tmpData[idxHalf]) / 2
}

func NewMedianFinder() *MedianFinder {
	res := MedianFinder{data: []int{}}
	return &res
}

func main() {
	mf := NewMedianFinder()
	mf.AddNum(1)
	mf.AddNum(2)
	fmt.Printf("median(1): %.1f\n", mf.FindMedian())
	mf.AddNum(3)
	fmt.Printf("median(2): %.1f\n", mf.FindMedian())
}
