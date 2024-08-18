package sorting

func SelectionSort(arr []int) []int {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		mini := i

		for j := i + 1; j < n; j++ {
			if arr[j] < arr[mini] {
				mini = j
			}
		}

		if mini != i {
			temp := arr[i]
			arr[i] = arr[mini]
			arr[mini] = temp
		}
	}
	return arr
}
