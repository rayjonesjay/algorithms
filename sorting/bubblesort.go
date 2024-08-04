// This file contains all various bubble sort implementations
package sort

func bubbleSort(list []int) (list []int) {
	for i := 0; i < len(list)-1; i++{
		for j := i+1; j < len(list); j++{
			if list[j] < list[i] {
			list[i] , list[j] = list[j] , list[i]
			}
		}		
	}
	return list
}
