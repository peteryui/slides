package qsort

import (
	"math/rand"
	"testing"
)

func generateRandomSlice(n int) []int {

	slice := make([]int, n, n)
	for i := 0; i < n; i++ {
		slice[i] = rand.Int() % (n * 100)
	}
	return slice
}

func isAscSorted(slice []int) bool {

	for i := 1; i < len(slice); i++ {
		if slice[i-1] > slice[i] {
			return false
		}
	}
	return true
}

func TestQsortBad(t *testing.T) {
	array := generateRandomSlice(100)

	qsortBad(array)

	if isAscSorted(array) {
		t.Error("the sorting is buggy", array)
	}
}
