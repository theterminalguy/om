package main

import (
	"fmt"

	"github.com/theterminalguy/om"
)

func main() {
	m := om.New()
	m.Add("first_name", "Simon Peter")
	m.Add("last_name", "Damian")
	m.Add("age", 27)
	m.Add("sex", "male")
	m.Add("married", false)

	// print json
	fmt.Println("JSON: ", m.JSON())

	// print values
	fmt.Println("Values: ", m.Values())
	fmt.Println("RValues: ", m.RValues())

	// print keys
	fmt.Println("Keys: ", m.Keys())
	fmt.Println("RKeys: ", m.RKeys())

	// print size
	fmt.Println("Size: ", m.Size())

	// print key value in order
	m.Each(func(key string, value interface{}) {
		fmt.Printf("key=%v, value=%v ", key, value)
	})
	fmt.Println("")
	// print key value reverse order
	m.REach(func(key string, value interface{}) {
		fmt.Printf("key=%v, value=%v ", key, value)
	})
	fmt.Println("")

	// Join the map
	fmt.Println(m.Join("=", "", " "))

	fmt.Println(m.HasKey("age"))

	// Check if map has any key called age
	d := m.HasAny(func(_ string, v interface{}) bool {
		return v == 27
	})
	fmt.Println("HasAny() => ", d)
}
