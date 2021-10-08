// Package om implements an ordered map datastructure.
// It also provides lots of usesful APIs for manipulating
// this structure with concepts borrowed from Ruby and Java.
package om

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

type omap struct {
	container map[string]interface{}
	keys      []string
	rkeys     []string
	values    []interface{}
	rvalues   []interface{}
}

// NewMap creates a new map
func New() *omap {
	return &omap{
		container: make(map[string]interface{}),
	}
}

// Add adds a key value pair to the map
func (m *omap) Add(k string, v interface{}) {
	m.keys = append(m.keys, k)
	m.values = append(m.values, v)
	m.rkeys = append([]string{k}, m.rkeys...)
	m.rvalues = append([]interface{}{v}, m.rvalues...)
	m.container[k] = v
}

// Get returns the value of the passed key
// If the value is not found an error is returned
func (m *omap) Get(key string) (interface{}, error) {
	if _, ok := m.container[key]; ok {
		return m.container[key], nil
	}
	return nil, errors.New("key not found!")
}

// Fetch returns the value of the passed key
// If the key is not found the defaultValue is returned
func (m *omap) Fetch(key string, defaultValue interface{}) interface{} {
	if v, err := m.Get(key); err == nil {
		return v
	}
	return defaultValue
}

func (m *omap) HasKey(key string) bool {
	_, err := m.Get(key)
	return err == nil
}

// Remove removes an item from the map by it's key
func (m *omap) Remove(k string) error {
	return nil
}

// Keys Returns all keys in the map
// Keys are returned in the order in which they are added
func (m *omap) Keys() []string {
	return m.keys
}

// RKeys Returns all keys in the map
// Keys are returned in reverse order
func (m *omap) RKeys() []string {
	return m.rkeys
}

// Values Returns all keys in the map
// Values are returned in the order in which they are added
func (m *omap) Values() []interface{} {
	return m.values
}

// RValues returns all keys in the map
// Keys are returned in reverse order
func (m *omap) RValues() []interface{} {
	return m.rvalues
}

// EQ compares two map for equality
// A map is equal if they both have the same keys and values
func (m1 *omap) EQ(m2 *omap) bool {
	if eq := m1.EQKey(m2); !eq {
		return !eq
	}
	for _, k := range m1.keys {
		// TODO, can you compare interfaces?
		v1, _ := m1.Get(k)
		v2, _ := m2.Get(k)
		if v1 != v2 {
			return false
		}
	}
	return true
}

// EQKey checks if both map has the same keys regardless of order
func (m1 *omap) EQKey(m2 *omap) bool {
	for _, k := range m1.keys {
		if _, err := m2.Get(k); err != nil {
			return false
		}
	}
	return true
}

// Each iterates through the map,
// Yielding each key and value to the callback function.
//
// key/value pairs are yielded in the order in which they where added
func (m *omap) Each(cb func(key string, value interface{})) {
	for _, k := range m.keys {
		cb(k, m.container[k])
	}
}

// REach iterates through the map,
// Yielding each key and value to the callback function.
//
// key/value pairs are yielded in reverse order
func (m *omap) REach(cb func(key string, value interface{})) {
	for _, k := range m.rkeys {
		cb(k, m.container[k])
	}
}

// Size returns the total number of items in the map
func (m *omap) Size() int {
	return len(m.container)
}

// OM returns the original golang map which is the
// underlying data structure used to store the key/value pairs
func (m *omap) OM() map[string]interface{} {
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

// Join glues all items in the map by key value
// in the order in which they where added
// glue is the string used to join key=value
// lpad is a text to pad on the left (key=value
// rpad is a text to pad on th right key=value)
func (m *omap) Join(glue, lpad, rpad string) string {
	var b strings.Builder
	m.Each(func(key string, value interface{}) {
		fmt.Fprintf(&b, "%v%v%v%v%v", lpad, key, glue, value, rpad)
	})
	return b.String()
}

// Join glues all items in the map by key value
// in reverse order
// glue is the string used to join key=value
// lpad is a text to pad on the left (key=value
// rpad is a text to pad on th right key=value)
func (m *omap) RJoin(glue, lpad, rpad string) string {
	var b strings.Builder
	m.REach(func(key string, value interface{}) {
		fmt.Fprintf(&b, "%v%v%v%v%v", lpad, key, glue, value, rpad)
	})
	return b.String()
}
