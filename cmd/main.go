package main

import (
	"fmt"

	"github.com/theterminalguy/om"
)

func main() {
	m := om.NewMap()
	m.Add("first_name", "Simon Peter")
	fmt.Println(m.JSON())
}
