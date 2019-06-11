package counter

import (
	"strconv"
	"sync/atomic"
)

// Counter struct definition
type Counter struct {
	value int32
}

// Increment function
func (c *Counter) Increment(amount int32) {
	atomic.AddInt32(&c.value, amount)
}

// Count function
func (c *Counter) Count() int {
	return int(atomic.LoadInt32(&c.value))
}

func (c *Counter) String() string {
	return strconv.Itoa(c.Count())
}
