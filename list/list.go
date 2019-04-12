/*

Package List high level methods

*/

package list

import "github.com/khekrn/workspace/iter"

// List interface
type List interface {
	Size() int

	Clear()

	Add(item interface{})

	Remove(item interface{}) bool

	Get(index int) (interface{}, error)

	Iter() iter.Iterator
}
