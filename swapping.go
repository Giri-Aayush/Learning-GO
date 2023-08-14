package main

import "fmt"

func swappingSliceElements(v []int) {
	for i, j := 0, len(v)-1; i < j; i, j = i+1, j-1 {
		v[i], v[j] = v[j], v[i]
	}
}

// No need of return since slices are passed by reference

func main() {
	testSlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	swappingSliceElements(testSlice)
	fmt.Println(testSlice)
	//the original slice values change
}
