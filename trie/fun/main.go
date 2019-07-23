package main

import (
	"fmt"
	"trie"
)

func main() {
	fmt.Println("Hello")

	obj := trie.NewTrie()

	obj.Add("John", "John is cool")
	obj.Add("Java", "Created by James Gosling in 90's")
	obj.Add("python", "Python is trending because of AI/ML")
	obj.Add("go", "go is awesome, Created by Rob Pike, Ken Thompson, Robert Griesmer")

	fmt.Println(obj.Size())

	item, found := obj.Get("go")

	fmt.Println(item, " ", found)

}
