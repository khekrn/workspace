package hashmap

import (
	"fmt"
	"strings"

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

func (h *HashMap) String() string {
	var sb strings.Builder
	sb.WriteString("{ ")
	for _, entry := range h.entryList {
		tempEntry := entry
		for tempEntry != nil {
			sb.WriteString(tempEntry.key)
			sb.WriteString(":")
			sb.WriteString(fmt.Sprintf("%v", tempEntry.value))
			sb.WriteString(", ")
			tempEntry = tempEntry.next
		}
	}
	sb.WriteString(" }")
	return sb.String()
}
