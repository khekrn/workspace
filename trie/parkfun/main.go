package main

import (
	"fmt"
	"trie/parker"
)

func main() {
	obj := parker.New()

	obj.Add("Java", "Java sucks")
	obj.Add("go", "Go is Awesome")

	node, _ := obj.Find("go")
	fmt.Println(node.Meta())

}
