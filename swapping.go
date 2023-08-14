package main

import "fmt"

func swappingSliceElements(v []int) []int {
	for i, j := 0, len(v)-1; i < j; i, j = i+1, j-1 {
		v[i], v[j] = v[j], v[i]
	}
	return v
}

func main() {
	testSlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(swappingSliceElements(testSlice))
}
