package om

import (
	"container/list"
	"encoding/json"
)

type omap struct {
	container map[string]interface{}
	keys      []string
	values    []interface{}
	rkeys     *list.List
	rvalues   *list.List
}

// NewMap creates a new map
func NewMap() *omap {
	m := make(map[string]interface{})
	return &omap{
		rkeys:     list.New(),
		rvalues:   list.New(),
		container: m,
	}
}

// Add adds a key value pair to the map
func (m *omap) Add(k string, v interface{}) {
	m.keys = append(m.keys, k)
	m.values = append(m.values, v)
	m.rkeys.PushFront(k)
	m.rvalues.PushFront(v)
	m.container[k] = v
}

// Remove removes an item from the map by key
func (m *omap) Remove(k string) error {
	return nil
}

// Keys Returns all keys in the map.
// Keys are returned in the order in which they are added
func (m *omap) Keys() []string {
	return m.keys
}

// RKeys Returns all keys in the map.
// Keys are returned in reverse order
func (m *omap) RKeys() *list.List {
	return m.rkeys
}

// Values Returns all keys in the map.
// Values are returned in the order in which they are added
func (m *omap) Values() []interface{} {
	return m.values
}

// RValues returns all keys in the map.
// Keys are returned in reverse order
func (m *omap) RValues() *list.List {
	return m.rvalues
}

// EQ compares two map for equality
func (*omap) EQ(m *omap) bool {
	return false
}

// Each iterates through the map,
// Yielding each key and value to the callback function.
//
// key/value pair are yielded in the order in which they where added
func (m *omap) Each(cb func(key string, value interface{})) {
	for _, k := range m.keys {
		cb(k, m.container[k])
	}
}

// REach iterates through the map,
// Yielding each key and value to the callback function.
//
// key/value pair are yielded in reverse order
func (*omap) REach(cb func(key string, value interface{})) {
}

// Size returns the total number of items in the map
func (m *omap) Size() int {
	return len(m.keys)
}

// OMap returns the original map which is the
// underlying data structure used to store the key value pairs
func (m *omap) OMap() map[string]interface{} {
	return m.container
}

// JSON returns a json representation of the map
func (m *omap) JSON() string {
	b, err := json.Marshal(m.container)
	if err != nil {
		return ""
	}
	return string(b)
}
