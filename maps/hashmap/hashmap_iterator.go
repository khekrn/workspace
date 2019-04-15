package hashmap

// Iterator definition
type Iterator struct {
	entriesSize  int
	currentEntry *Entry
	entriesList  []*Entry
	index        int
	entryIndex   int
}

// NewIterator returns new hashmap iterator
func NewIterator(h *HashMap) *Iterator {
	return &Iterator{currentEntry: nil, entriesSize: h.entrySize, entriesList: h.entryList, index: 0, entryIndex: 0}
}

// HasNext returns true if element exist
func (iter *Iterator) HasNext() bool {
	res := false
	if iter.index < iter.entriesSize {
		if iter.currentEntry != nil && iter.currentEntry.next != nil {
			iter.currentEntry = iter.currentEntry.next
		} else {
			iter.currentEntry = nil
		}
		for iter.currentEntry == nil {
			iter.currentEntry = iter.entriesList[iter.entryIndex]
			iter.entryIndex++
		}
		iter.index++
		res = true
	}
	return res
}

// Key returns key
func (iter *Iterator) Key() string {
	return iter.currentEntry.key
}

// Value returns value
func (iter *Iterator) Value() interface{} {
	return iter.currentEntry.value
}
