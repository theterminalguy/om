package om

type KeyValue struct {
	Key   string
	Value interface{}
}

type Map struct {
	container map[string]interface{}
	keys []string
	values []interface{}
}

// Adds a key value pair to the map
func (m *Map) Add(k string, v interface{}) {
	m.keys = append(m.keys, k)
	m.values = append(m.values, v)
	m.container[k] = v
}

// Removes an item from the map by key
func (m *Map) Remove(k string) error {
	return nil
}

// Returns all keys in the map
// Keys are returned in the order in which they are added
// Keys are returned in reverse
func (*Map) Keys() []string {
	return nil
}

// Returns all keys in the map
// Keys are returned in reverse
func (*Map) RKeys() []string {
	return nil
}

// Returns all keys in the map
// Values are returned in the order in which they are added
func (*Map) Values() []interface{} {
	return nil
}

// Returns all keys in the map
// Keys are returned in reverse
func (*Map) RValues() []interface{} {
	return nil
}

// Compares two map for equality
func (*Map) EQ(m *Map) bool {
	return false
}

// Iterates through the map,
// Yielding each key and value to the callback function.
//
// KeyValue pair are yielded in the order in which they where added
func (*Map) Each(cb func(key string, value interface{})) {
}

// Iterates through the map,
// Yielding each key and value to the callback function.
//
// KeyValue pair are yielded in reverse order
func (*Map) REach(cb func(key string, value interface{})) {
}

// Returns the total number of items in the map
func (*Map) Size() int {
	return 0
}
