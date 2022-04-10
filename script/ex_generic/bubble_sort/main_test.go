package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBubbleSort(t *testing.T) {
	val := []int{2, 1, 6, 4, 9}
	valSorted := []int{1, 2, 4, 6, 9}
	bubbleSortGeneric(val)
	assert.Equal(t, val, valSorted)
}
