package main

import (
	"fmt"
	"slices"
)

func main() {
	// binary search works on sorted data
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}

	// let's look for the number 12 since there are 12 months in a year
	target := 12

	// recursiveBS takes an array and a target value and returns the position that the target was
	// found and a bool value to mean it was found or not. if not found position will be -1 and found will be false
	position, found := recursiveBS(arr,target,0, len(arr)-1)
	if found {
		fmt.Printf("%d was found at %d th index\n",target, position)
	}else{
		fmt.Printf("%d was not found\n",target)
	}
}


// recursiveBS attempts to compute binary search using recursion instead of iteration
func recursiveBS(arr []int, target, low, high int) (position int, found bool) {

	if !slices.Contains(arr, target) {
		return -1,false
	}

	left := low
	right := high
	mid := (left +right)/2
	if arr[mid] == target{
		return mid,true
	}

	if arr[mid] > target{
		right = mid - 1
	}else if arr[mid] < target{
		left = mid + 1
	}
	return recursiveBS(arr,target,left,right)
}
