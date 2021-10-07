package om
type Map struct {

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
func (*Map) EQ(m *Map) bool  {
	return false
}

// Iterates through the map,
// Yielding each key and value to the callback function.
// 
// KeyValue pair are yielded in the order in which they where added 
func (*Map) Each(cb func(key string, value interface{})) {
}