package trie

// Service trie service contract
type Service interface {
	Add(key string, value interface{}) error
	Contains(key string) bool
	Get(key string) (interface{}, bool)
	Delete(key string) error
	Size() int
}

// Node struct definition
type Node struct {
	val      rune
	term     bool
	depth    int
	meta     interface{}
	mask     uint64
	parent   *Node
	children map[rune]*Node
}

// NewChild new instance of Node with given details
func (n *Node) NewChild(val rune, bitmask uint64, meta interface{}, term bool) *Node {
	node := &Node{
		val:      val,
		mask:     bitmask,
		term:     term,
		meta:     meta,
		parent:   n,
		children: make(map[rune]*Node),
		depth:    n.depth + 1,
	}
	n.children[val] = node
	n.mask |= bitmask
	return node
}

// Meta gives you meta data stored in the node
func (n *Node) Meta() interface{} {
	return n.meta
}

// Parent return's parent node of the current node
func (n *Node) Parent() *Node {
	return n.parent
}

// Children provides childrens of the node
func (n *Node) Children() map[rune]*Node {
	return n.children
}

// Value rune key
func (n *Node) Value() interface{} {
	return n.val
}

// IsEnd last child node or not
func (n *Node) IsEnd() bool {
	return n.term
}

// Depth gives depth of the node
func (n *Node) Depth() int {
	return n.depth
}
