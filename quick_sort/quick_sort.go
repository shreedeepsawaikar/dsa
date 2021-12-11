package main

func main() {
	arr := []int{2, 7, 8, 9, 1, 5, 10}
	n := len(arr) - 1
	quickSort(arr, 0, n)

	for _, element := range arr {
		print(element, " ")
	}
}

func quickSort(arr []int, low int, high int) []int {

	if low < high {

		pi := partition(arr, low, high)

		quickSort(arr, low, pi-1)
		quickSort(arr, pi+1, high)
	}
	return arr
}

func partition(arr []int, low int, high int) int {
	pi := arr[low]
	i := low
	j := high

	for i < j {
		for arr[i] <= pi && i < high {
			i++
		}

		for arr[j] > pi {
			j--
		}

		if i < j {
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[j], arr[low] = arr[low], arr[j]

	return j

}
