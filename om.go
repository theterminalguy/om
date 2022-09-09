// Package om implements an ordered map datastructure
// with an extensilve list of APIS borrowed from other languages
package om

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

var ErrKeyNotFound error = errors.New("key not found")

type omap struct {
	container       map[string]interface{}
	keys, rkeys     []string
	values, rvalues []interface{}
}

// New creates a new ordered map
func New() *omap {
	return &omap{
		container: make(map[string]interface{}),
	}
}

// Add adds a key value pair to the map. If the key already exist, it's value is updated
func (m *omap) Add(key string, value interface{}) {
	if _, err := m.Get(key); err == nil {
		m.Put(key, value)
		return
	}
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

// Put upate a key's value. If the key is not found, returns an error
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

// Get returns the value of the passed key. If the key is not found an error is returned
func (m *omap) Get(key string) (interface{}, error) {
	if _, ok := m.container[key]; ok {
		return m.container[key], nil
	}
	return nil, ErrKeyNotFound
}

// Fetch returns the value of the passed key. If the key is not found the defaultValue is returned
func (m *omap) Fetch(key string, defaultValue interface{}) interface{} {
	if v, err := m.Get(key); err == nil {
		return v
	}
	return defaultValue
}

// Index return the key/value pair stored at the given index. If the index is out of bound,
// returns an empty string and nil
func (m *omap) Index(pos int) (string, interface{}) {
	if pos > 0 && pos < m.Size() {
		return m.keys[pos], m.values[pos]
	}
	return "", nil
}

// GetKeyIndex returns the numeric index for the given key. If the key is not found, returns -1
func (m *omap) GetKeyIndex(key string) int {
	for i, k := range m.Keys() {
		if key == k {
			return i
		}
	}
	return -1
}

// ValuesAt Returns a new slice containing values for the given keys
func (m *omap) ValuesAt(keys ...string) []interface{} {
	var values []interface{}
	for _, k := range keys {
		if v, err := m.Get(k); err == nil {
			values = append(values, v)
			continue
		}
		values = append(values, nil)
	}
	return values
}

// HasKey retruns true if the key is contained in the map
func (m *omap) HasKey(key string) bool {
	_, err := m.Get(key)
	return err == nil
}

// HasAny Returns true if any element satisfies a given criterion; false otherwise.
func (m *omap) HasAny(cb func(key string, value interface{}) bool) bool {
	for _, k := range m.Keys() {
		if cb(k, m.container[k]) {
			return true
		}
	}
	return false
}

// Delete removes the entry for the given key and returns its associated value.
func (m *omap) Delete(key string) (interface{}, error) {
	// TODO: adjust rkeys, lkeys rvalues and lvalues
	v, err := m.Get(key)
	if err != nil {
		return nil, err
	}
	// remove values & rvalues
	// remove keys & rkeys
	delete(m.container, key)
	return v, nil
}

// DeleteIF calls the callback with each key/value pair. Deletes each entry for which the callback
// returns true
func (m *omap) DeleteIF(cb func(key string, value interface{}) bool) *omap {
	for _, k := range m.Keys() {
		if cb(k, m.container[k]) {
			m.Delete(k)
		}
	}
	return m
}

// KeepIF calls the callback with each key/value pair. Keeps each entry for which the callback
// returns true, otherwise deletes the entry from the map
func (m *omap) KeepIF(cb func(key string, value interface{}) bool) *omap {
	for _, k := range m.Keys() {
		if !cb(k, m.container[k]) {
			m.Delete(k)
		}
	}
	return m
}

// Filter returns a new map whose entries are those for which the callback returns true
func (m *omap) Filter(cb func(key string, value interface{}) bool) *omap {
	nm := New()
	for _, k := range m.Keys() {
		if cb(k, m.container[k]) {
			nm.Add(k, m.container[k])
		}
	}
	return nm
}

// Filter_ modifies the original map, keeping entires for which the callback returns true
func (m *omap) Filter_(cb func(key string, value interface{}) bool) *omap {
	for _, k := range m.Keys() {
		if !cb(k, m.container[k]) {
			m.Delete(k)
		}
	}
	return m
}

// Slice returns a new map containing the entries for the given keys
func (m *omap) Slice(keys ...string) *omap {
	nm := New()
	for _, k := range keys {
		if v, err := m.Get(k); err == nil {
			nm.Add(k, v)
		}
	}
	return nm
}

// Compact returns a new map with all nil entries removed
func (m *omap) Compact() *omap {
	nm := New()
	for _, k := range m.Keys() {
		if v, _ := m.Get(k); v != nil {
			nm.Add(k, v)
		}
	}
	return nm
}

// Compact modifies the original map with all nil entries removed
func (m *omap) Compact_() *omap {
	for _, k := range m.Keys() {
		if v, _ := m.Get(k); v == nil {
			m.Delete(k)
		}
	}
	return m
}

// Except returns a new map excluding entries for the given keys
func (m *omap) Except(keys ...string) *omap {
	nm := New()
	for _, k := range keys {
		nm.Add(k, m.container[k])
	}
	for _, k := range m.Keys() {
		if _, err := nm.Get(k); err == nil {
			nm.Delete(k)
			continue
		}
		if v, err := m.Get(k); err == nil {
			nm.Add(k, v)
		}
	}
	return nm
}

// Merge returns a new map formed by merging the other map into a copy of self
func (m1 *omap) Merge(m2 *omap) *omap {
	nm := New()
	for _, k := range m1.Keys() {
		nm.Add(k, m1.container[k])
	}
	for _, k := range m2.Keys() {
		nm.Add(k, m2.container[k])
	}
	return nm
}

// Merge_ modifies the original by merging the new map into a copy of self
func (m1 *omap) Merge_(m2 *omap) *omap {
	for _, k := range m2.Keys() {
		if v, err := m2.Get(k); err != nil {
			m1.Add(k, v)
		}
	}
	return m1
}

// Clear, removes all map entries and returns a pointer to the same map
func (m *omap) Clear() *omap {
	m.keys, m.rkeys = []string{}, []string{}
	m.values, m.rvalues = []interface{}{}, []interface{}{}
	m.container = map[string]interface{}{}
	return m
}

// Keys Returns all keys in the map. Keys are returned in the order in which they are added
func (m *omap) Keys() []string {
	return m.keys
}

// RKeys Returns all keys in the map. Keys are returned in reverse order
func (m *omap) RKeys() []string {
	return m.rkeys
}

// Values Returns all keys in the map. Values are returned in the order in which they are added
func (m *omap) Values() []interface{} {
	return m.values
}

// RValues returns all keys in the map. Keys are returned in reverse order
func (m *omap) RValues() []interface{} {
	return m.rvalues
}

// EQ compares two map for equality. A map is equal if they both have the same keys and values
func (m1 *omap) EQ(m2 *omap) bool {
	if eq := m1.EQKey(m2); !eq {
		return !eq
	}
	for _, k := range m1.keys {
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

// Each iterates through the map, yielding each key and value to the callback function.
// key/value pairs are yielded in the order in which they where added
func (m *omap) Each(cb func(key string, value interface{})) {
	for _, k := range m.keys {
		cb(k, m.container[k])
	}
}

// REach iterates through the map, yielding each key and value to the callback function.
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

// OM returns the original golang map which is the underlying data structure used to store the key/value pairs
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

// Join glues all items in the map by key value in the order in which they where added
//
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

// Join glues all items in the map by key value in reverse order
//
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

// String returns a new String containing the map entries
func (m *omap) String() string {
	var b strings.Builder
	for k, v := range m.container {
		fmt.Fprintf(&b, "%v=%v ", k, v)
	}
	return strings.TrimSpace(b.String())
}
