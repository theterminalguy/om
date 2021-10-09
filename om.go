// Package om implements an ordered map datastructure
// with an extensilve list of APIS borrowed from other languages
package om

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

var ErrKeyNotFound error = errors.New("key not found!")
type omap struct {
	container       map[string]interface{}
	keys, rkeys     []string
	values, rvalues []interface{}
}

// NewMap creates a new map
func New() *omap {
	return &omap{
		container: make(map[string]interface{}),
	}
}

// Add adds a key value pair to the map
func (m *omap) Add(key string, value interface{}) {
	m.keys, m.rkeys = append(m.keys, key), append([]string{key}, m.rkeys...)
	m.values, m.rvalues = append(m.values, value), append([]interface{}{value}, m.rvalues...)
	m.container[key] = value
}

// Set updates a key's value using Put. If the key is not found, it calls Add
func (m *omap) Set(k string, v interface{}) {
	if err := m.Put(k, v); err != nil {
		m.Add(k, v)
	}
}

// Put upate a key's value. If the key is not found, returns and error
func (m *omap) Put(key string, value interface{}) error {
	if _, err := m.Get(key); err != nil {
		return err
	}
	lpos := 0
	for i, k := range m.keys {
		if key == k {
			lpos = i
		}
	}
	rpos := (m.Size() - lpos) - 1
	m.values[lpos], m.rvalues[rpos], m.container[key] = value, value, value
	return nil
}

// Get returns the value of the passed key.
// If the value is not found an error is returned
func (m *omap) Get(key string) (interface{}, error) {
	if _, ok := m.container[key]; ok {
		return m.container[key], nil
	}
	return nil, ErrKeyNotFound
}

// Fetch returns the value of the passed key
// If the key is not found the defaultValue is returned
func (m *omap) Fetch(key string, defaultValue interface{}) interface{} {
	if v, err := m.Get(key); err == nil {
		return v
	}
	return defaultValue
}

// Index return the key/value pair stored at the given index.
// If the index is out of bound, returns an empty string and nil
func (m *omap) Index(pos int) (string, interface{}) {
	if pos > 0 && pos < m.Size() {
		return m.keys[pos], m.values[pos]
	}
	return "", nil
}

// GetKeyIndex returns the numeric index for the given key.
// If the key is not found, returns -1
func (m *omap) GetKeyIndex(key string) int {
	for i, k := range m.keys {
		if key == k {
			return i
		}
	}
	return -1
}

// ValuesAt Returns a new slice containing values for the given keys
func (m *omap) ValuesAt() {
	
}

// HasKey retruns true if the key is contained in the map
func (m *omap) HasKey(key string) bool {
	_, err := m.Get(key)
	return err == nil
}

// HasAny Returns true if any element satisfies a given criterion; false otherwise.
func (m *omap) HasAny(cb func(key string, value interface{}) bool) bool {
	for _, k := range m.keys {
		if cb(k, m.container[k]) {
			return true
		}
	}
	return false
}

// Delete removes the entry for the given key and returns its associated value.
// If the key is not found, returns nil
func (m *omap) Delete(key string) (interface{}, error) {
	v, err := m.Get(key)
	if err != nil {
		return nil, err
	}
	delete(m.container, key)
	return v, nil
}

func (m *omap) DeleteIF() {

}

// Clear, removes all map entries and returns a pointer to self
func (m *omap) Clear() *omap {
	m.keys, m.rkeys = []string{}, []string{}
	m.values, m.rvalues = []interface{}{}, []interface{}{}
	m.container = map[string]interface{}{}
	return m
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

// Empty Returns true if there are no map entries, false otherwise:
func (m *omap) Empty() bool {
	return m.Size() == 0
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
