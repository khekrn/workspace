package singlelist

// Iterator definition
type Iterator struct {
	list        *SingleLinkedList
	index       int
	currentNode *node
}

// Iter method implementation for SingleLinkedList
func (slist *SingleLinkedList) Iter() Iterator {
	return Iterator{list: slist, index: 0, currentNode: nil}
}

// HasNext checks whether next element exist in list while iterating
func (iter *Iterator) HasNext() bool {
	hasNext := false
	if iter.index < iter.list.size {
		iter.index++
		hasNext = true
	}
	if iter.index == 1 {
		iter.currentNode = iter.list.head
	} else {
		iter.currentNode = iter.currentNode.next
	}
	return hasNext
}

// Next returns the next element from the list
func (iter *Iterator) Next() interface{} {
	return iter.currentNode.item
}
