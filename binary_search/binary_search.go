package main

func main() {
	arr := []int{1, 3, 5, 7, 9, 13}

	print(binarySearch(arr, 0, len(arr)-1, 11))
}

func binarySearch(arr []int, low int, high int, element int) int {

	if high < low {
		return -1
	}
	mid := (low + high) / 2

	if arr[mid] == element {
		return mid
	} else if arr[mid] > element {
		return binarySearch(arr, low, mid-1, element)
	} else {
		return binarySearch(arr, mid+1, high, element)
	}
}
