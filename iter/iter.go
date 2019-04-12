// Package iter which defines iterator
package iter

// Iterator interface
type Iterator interface {
	HasNext() bool

	Next() interface{}
}
