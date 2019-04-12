package main

import (
	"algorithms/maps/hashmap"
	"fmt"
)

func main() {
	// arr := sort.GenerateSlice(10)
	// res := sort.MergeSort(arr)
	// fmt.Println("Unsorted = ", arr)
	// fmt.Println("Sorted = ", res)
	hashMap := hashmap.New(10)
	hashMap.Put("Alice", 10)
	hashMap.Put("Bob", 20)
	hashMap.Put("Teena", 18)
	fmt.Println(hashMap)
}
