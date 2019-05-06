package main

import (
<<<<<<< HEAD
	"github.com/khekrn/workspace/read"
)

func main() {
	read.ReadInChunks("dummy.txt")
=======
	"fmt"
	"github.com/khekrn/workspace/sort"
	HashMap "github.com/khekrn/workspace/maps/hashmap"
	"github.com/khekrn/workspace/filters"
)

func main() {
	arr := []int{1, 2, 3, 4, 5}
	res := sort.InsertionSort(arr)
	fmt.Println("Unsorted = ", arr)
	fmt.Println("Sorted = ", res)
	
	hashMap := HashMap.New(5)
	hashMap.Put("Alice", 10)
	hashMap.Put("Bob", 20)
	hashMap.Put("Teena", 18)
	hashMap.Put("John", 34)
	hashMap.Put("Reese", 40)
	hashMap.Put("Root", 37)
	hashMap.Put("Kiran", 36)
	fmt.Println(hashMap)
	fmt.Println("Hash Map Size := ", hashMap.Length())

	mapIter := hashMap.Iter()
	for mapIter.HasNext() {
		key := mapIter.Key()
		value := mapIter.Value()
		fmt.Println(key, " - ", value)
	}

	item, err := hashMap.Get("Alice")
	fmt.Println("For Key Alice = ", item, err)
	item, err = hashMap.Get("Teena")
	fmt.Println("For Key Teena = ", item, err)
	item, err = hashMap.Get("Ta")
	fmt.Println("For Key Ta = ", item, err)
	fmt.Println(hashMap)
	fmt.Println("HashMap Size = ", hashMap.Length())
	mapIter = hashMap.Iter()
	for mapIter.HasNext() {
		key := mapIter.Key()
		value := mapIter.Value()
		fmt.Println(key, " - ", value)
	}

	bloom, _ := filters.NewBloomFilter(1000, 0.1)
	
	v1 := []byte{1, 2, 3, 4, 5}
	bloom.Add((v1))
	fmt.Println("\n\n\n")
	fmt.Println(bloom.Contains(v1))
	fmt.Println(bloom.Contains([]byte{10}))
>>>>>>> 1debb3f4931fabdf050c67ec66742316e49e0c1a
}
