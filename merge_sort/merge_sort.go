package main

func main() {
	arr := []int{2, 7, 8, 9, 1, 5, 10}
	n := len(arr) - 1
	mergeSort(arr, 0, n)

	for _, element := range arr {
		print(element, " ")
	}
}

func mergeSort(arr []int, low int, high int) {

	if low < high {
		mid := (low + high) / 2

		mergeSort(arr, low, mid)
		mergeSort(arr, mid+1, high)

		merge(arr, low, mid, high)

	}
	//return arr

}

func merge(a []int, low, mid int, high int) {
	b := make([]int, len(a))
	i := low
	j := mid + 1
	k := low

	for i <= mid && j <= high {
		if a[i] < a[j] {

			b[k] = a[i]
			i++
			k++
		} else {
			b[k] = a[j]
			j++
			k++
		}
	}
	if i > mid {
		for j <= high {
			b[k] = a[j]
			k++
			j++
		}

	} else {
		for i <= mid {
			b[k] = a[i]
			i++
			k++
		}
	}

	for k := low; k < high; k++ {
		a[k] = b[k]
	}

	//return a
}
