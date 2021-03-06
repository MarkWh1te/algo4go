package sort

import (
	"github.com/MarkWh1te/algo4go/util"
)

// Partition moves the small number to head, big number to end, same to middle.
func Partition(arr *[]int, l int, r int, n int) (int, int) {
	start := l
	current := l
	end := r
	pivot := (*arr)[n]
	for current <= end {
		if (*arr)[current] < pivot {
			util.Swap(&(*arr)[current], &(*arr)[start])
			current++
			start++
		} else if (*arr)[current] > pivot {
			util.Swap(&(*arr)[current], &(*arr)[end])
			end--
		} else if (*arr)[current] == pivot {
			current++
		}
	}
	return start, end
}

// QuickSort is a util.Randomize sorting algorithms
func QuickSort(arr *[]int, left int, right int) {
	if right > left {
		// util.Randomize pivot
		start, end := Partition(arr, left, right, util.Random(left, right))
		QuickSort(arr, left, start-1)
		QuickSort(arr, end+1, right)
	}
}
