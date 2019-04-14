package hashmap

import (
	"fmt"
	"strings"

	"github.com/khekrn/workspace/iter"
	"github.com/khekrn/workspace/maps"
)

const (
	loadFactor = 0.75
	intMax     = 0x7FFFFFFF
)

// Entry hashmap entry type
type Entry struct {
	key   string
	value interface{}
	next  *Entry
}

// HashMap definition
type HashMap struct {
	mapSize   int
	hashFunc  func(key string, size int) int
	entryList []*Entry
	entrySize int
}

func hashFunction(key string, size int) int {
	res := 0
	for i := 0; i < len(key); i++ {
		res = 31*res + int(key[i])
	}
	return (res * intMax) % size
}

// New creates new HashMap
func New(defaultCapacity int) *HashMap {
	if defaultCapacity <= 0 {
		defaultCapacity = 10
	}
	return &HashMap{mapSize: defaultCapacity, hashFunc: hashFunction,
		entryList: make([]*Entry, defaultCapacity), entrySize: 0}
}

// Length gives you the hashmap size
func (h *HashMap) Length() int {
	return h.entrySize
}

// Put insert's the given key and value into the hashmap
func (h *HashMap) Put(key string, value interface{}) error {
	if key == "" {
		return maps.ErrInvalidKey
	}
	index := h.hashFunc(key, h.mapSize)
	entry := h.entryList[index]
	if entry == nil {
		h.entryList[index] = &Entry{key: key, value: value}
		h.entrySize++
	} else {
		found := false
		previousEntry := entry
		for entry != nil {
			if entry.key == key {
				entry.value = value
				h.entrySize++
				found = true
				break
			}
			previousEntry = entry
			entry = entry.next
		}
		if !found {
			newEntry := &Entry{key: key, value: value}
			h.entrySize++
			previousEntry.next = newEntry
		}
	}
	return nil
}

// Get returns value for the given key
func (h *HashMap) Get(key string) (interface{}, error) {
	if key == "" {
		return nil, maps.ErrInvalidKey
	}

	keyIndex := h.hashFunc(key, h.mapSize)
	entry := h.entryList[keyIndex]
	for entry != nil {
		if entry.key == key {
			return entry.value, nil
		}
		entry = entry.next
	}
	return nil, maps.ErrKeyNotFound
}

// Delete remove's element from hashmap
func (h *HashMap) Delete(key string) (bool, error) {
	if key == "" {
		return false, maps.ErrInvalidKey
	}

	keyIndex := h.hashFunc(key, h.mapSize)
	entry := h.entryList[keyIndex]

	previousEntry := entry
	for entry != nil {
		if entry.key == key {
			if entry == previousEntry {
				entry = entry.next
			} else {
				previousEntry.next = entry.next
				entry = nil
			}
			h.entrySize--
			return true, nil
		}
		previousEntry = entry
		entry = entry.next
	}
	return false, maps.ErrKeyNotFound
}

// Iter returns iterator struct
func (h *HashMap) Iter() iter.IteratorWithKey {
	return NewIterator(h)
}

func (h *HashMap) String() string {
	var sb strings.Builder
	sb.WriteString("{ ")
	index := 0
	for _, entry := range h.entryList {
		tempEntry := entry
		for tempEntry != nil {
			if index > 0 {
				sb.WriteString(", ")
			}
			sb.WriteString(tempEntry.key)
			sb.WriteString(":")
			sb.WriteString(fmt.Sprintf("%v", tempEntry.value))
			tempEntry = tempEntry.next
			index++
		}
	}
	sb.WriteString(" }")
	return sb.String()
}

func (h *HashMap) resize() {
	if float64(h.entrySize)/float64(h.mapSize) >= loadFactor {
		newSize := h.mapSize * 2
		newEntryList := make([]*Entry, newSize)
		for _, entry := range h.entryList {
			if entry != nil {
				index := h.hashFunc(entry.key, newSize)
				newEntryList[index] = entry
			}
		}
		h.mapSize = newSize
		h.entryList = newEntryList
	}
}
