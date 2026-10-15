package collect

import "fmt"

// MustGet returns the value of the first pair with key k. It panics with an
// error that wraps [ErrItemNotFound] and names k if there is none.
// [KeyedCollection.Get] reports the absence instead.
func (c KeyedCollection[K, V]) MustGet(k K) V {
	v, ok := c.Get(k)
	if !ok {
		panic(fmt.Errorf("%w: no key %v", ErrItemNotFound, k))
	}
	return v
}
