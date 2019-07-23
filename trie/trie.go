package trie

import "errors"

var (
	// ErrInvalidKey invalid key error message
	ErrInvalidKey = errors.New("invalid key found")
)

// Empty rune value
const Empty = 0x0

// Trie struct definition
type Trie struct {
	root *Node

	size int
}

// NewTrie new instance of Trie
func NewTrie() *Trie {
	return &Trie{
		root: &Node{children: make(map[rune]*Node), depth: 0},
		size: 0,
	}
}

// Root trie root node
func (t *Trie) Root() *Node {
	return t.root
}

// Add stores/updates given key and value in trie
func (t *Trie) Add(key string, value interface{}) *Node {
	if len(key) == 0 {
		return nil
	}

	t.size++

	runes := []rune(key)

	bitmask := t.maskRuneSlices(runes)
	node := t.root

	node.mask |= bitmask

	for index := range runes {
		startChar := runes[index]
		bitmask = t.maskRuneSlices(runes[index:])
		if n, ok := node.children[startChar]; ok {
			node = n
			node.mask |= bitmask
		} else {
			node = node.NewChild(startChar, bitmask, nil, false)
		}
	}

	node = node.NewChild(Empty, 0, value, true)

	return node
}

// Get retrieves given key if exist
func (t *Trie) Get(key string) (interface{}, bool) {
	if len(key) == 0 {
		return nil, false
	}

	current := t.root

	if current == nil {
		return nil, false
	}

	runes := []rune(key)

	for _, item := range runes {
		node, exist := current.children[item]
		if !exist {
			return nil, false
		}
		current = node
	}

	if n, ok := current.Children()[Empty]; ok {
		current = n
	}

	if current.term {
		return current.meta, true
	}

	return current.meta, false

}

// Size return's Trie item count
func (t *Trie) Size() int {
	return t.size
}

func (t *Trie) maskRuneSlices(rs []rune) uint64 {
	var mask uint64
	for _, item := range rs {
		mask |= uint64(1) << uint64(item-'a') // 'a' is 97
	}
	return mask
}
