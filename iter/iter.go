// Package iter which defines iterator
package iter

// Iterator interface
type Iterator interface {
	HasNext() bool

	Next() interface{}
}

// IteratorWithKey for Key/Value Data Structures
type IteratorWithKey interface {
	Key() string

	Value() interface{}

	HasNext() bool
}
