package hashmap

// Iterator definition
type Iterator struct {
	entriesSize int
	entriesList []*Entry
	index       int
	entryIndex  int
}

// NewIterator returns new hashmap iterator
func NewIterator(h *HashMap) *Iterator {
	return &Iterator{entriesSize: h.entrySize, entriesList: h.entryList, index: 0, entryIndex: 0}
}

// HasNext returns true if element exist
func (iter *Iterator) HasNext() bool {
	return iter.index < iter.entriesSize
}

// Key returns key
func (iter *Iterator) Key() string {
	return ""
}

// Value returns value
func (iter *Iterator) Value() interface{} {
	return nil
}
