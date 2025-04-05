package main

// BubbleSort sorts a slice of integers in ascending order using
// the Bubble Sort algorithm.
//
// This function repeatedly steps through the list, compares adjacent elements,
// and swaps them if they are in the wrong order. The pass through the list
// is repeated until the list is sorted.
//
// Time complexity: O(n²) in worst case.
//
// Parameters:
//   - arr: slice of integers to be sorted
func BubbleSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}
