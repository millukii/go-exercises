package main

func main() {
	array := []int{3, 42, 1, 5}
	quicksort(array, 0, len(array)-1)
}

func quicksort(array []int, left, right int) {

	if left >= right {
		return
	}
	pivot := array[(left+right)/2]
	index := partition(array, left, right, pivot)
	quicksort(array, left, index-1)
	quicksort(array, index, right)
}

func partition(array []int, left, right, pivot int) int {
	for left <= right {
		for array[left] < pivot {
			left++
		}
		for array[right] > pivot {
			right--
		}

		if left <= right {
			swap(array, left, right)
			left++
			right--

		}
	}
	return left
}
