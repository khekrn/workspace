package singlelist

import "errors"

type node struct {
	item interface{}
	next *node
}

// SingleLinkedList data definition
type SingleLinkedList struct {
	head *node
	tail *node
	size int
}

// New creates new SingleLinkedList instance
func New() *SingleLinkedList {
	return &SingleLinkedList{head: nil, tail: nil, size: 0}
}

// Size return's the length of the SingleLinkedList
func (slist *SingleLinkedList) Size() int {
	return slist.size
}

// Clear removes the contents of SingleLinkedList
func (slist *SingleLinkedList) Clear() {
	slist.head, slist.tail, slist.size = nil, nil, 0
}

// Add inserts new node at the end of the list
func (slist *SingleLinkedList) Add(item interface{}) {
	if slist.head == nil {
		slist.head = &node{item: item, next: nil}
		slist.tail = slist.head
	} else {
		newTail := &node{item: item, next: nil}
		slist.tail.next = newTail
		slist.tail = newTail
	}
	slist.size++
}

// Remove delete the given element from SingleLinkedList if exist
func (slist *SingleLinkedList) Remove(item interface{}) bool {
	removeSuccess := false

	if slist.size > 0 {
		previousNode, currentNode := slist.head, slist.head
		for currentNode != nil {
			if currentNode.item == item {
				removeSuccess = true
				break
			}
			previousNode = currentNode
			currentNode = currentNode.next
		}

		if currentNode == slist.head {
			newHead := slist.head.next
			slist.head = newHead
			slist.size--
		} else if currentNode == slist.tail {
			slist.tail = previousNode
			slist.size--
		} else {
			previousNode.next = currentNode.next
			slist.size--
		}
	}

	return removeSuccess
}

// Get fetches the item from SingleLinkedList for the given index
func (slist *SingleLinkedList) Get(index int) (interface{}, error) {
	var item interface{}
	var err error
	if index >= slist.size {
		err = errors.New("invalid index, index out of range")
	} else {
		if index == 0 {
			item = slist.head.item
		} else if index == slist.size-1 {
			item = slist.tail.item
		} else {
			i := 0
			currentNode := slist.head
			for i != index {
				currentNode = currentNode.next
				i++
			}
			item = currentNode.item
		}
	}

	return item, err
}
