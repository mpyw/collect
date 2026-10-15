package collect

// Count reads the collection and returns how many values it has.
func (c Collection[T]) Count() int {
	n := 0
	for range c.Seq() {
		n++
	}
	return n
}

// Count reads the collection and returns how many pairs it has. A key that
// comes more than once is counted each time.
func (c KeyedCollection[K, V]) Count() int {
	n := 0
	for range c.Seq() {
		n++
	}
	return n
}
