package gcounter

import (
	"log"
	"sync"

	uuid "github.com/satori/go.uuid"
)

// conflict free replicated data type

// GCounter represent a G-counter in CRDT, which is a state-based grow-only counter that only supports increments.
type GCounter struct {
	mtx     sync.RWMutex
	id      string // Unique id for each replica
	counter map[string]int
}

// NewGCounter new instance
func NewGCounter() *GCounter {
	uid, err := uuid.NewV4()
	if err != nil {
		log.Fatal("cannot able to initiate uuid V4")
		return nil
	}
	return &GCounter{
		id:      uid.String(),
		counter: map[string]int{},
	}
}

func (gc *GCounter) Increment() {

}
