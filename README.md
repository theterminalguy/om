# om
Ordered Map

## Installation

```bash
$ go get github.com/theterminalguy/om
```

## API docs

### `New`
> New creates a new ordered map

```go
m := om.New()
```

### `Add(key string, value interface{})`
> Add adds a key value pair to the map. If the key already exist, it's value is updated

```go
m := om.New()
m.Add("first_name", "Simon Peter")
m.Add("last_name", "Damian")
m.Add("age", 27)
m.Add("sex", "male")
m.Add("married", false)

m.Add("age", 28) // updates the age to 28 because key exists
``` 

### `Set(k string, v interface{})`
> Updates a key's value using Put. If the key is not found, it calls Add

```go
m := om.New() 
m.Set("first_name", "Simon Peter") // calls Add
```

### `Put(key string, value interface{}) error`
> Updates a key's value. If the key is not found, it returns an error

```go
m := om.New()
m.Put("first_name", "Simon Peter") // returns an error

m.Add("first_name", "Simon Peter")

m.Put("first_name", "Simon") // updates the value of first_name to Simon without error
```

### `Get(key string) (interface{}, error)`
> Returns the value of a key. If the key is not found, it returns an error

```go
m := om.New()
m.Add("first_name", "Simon Peter")

v, err := m.Get("first_name")
if err != nil {
    panic(err)
}

fmt.Println(v) // prints Simon Peter
```

### `Fetch(key string, defaultValue interface{}) interface{}`
> Returns the value of a key. If the key is not found, it returnns the default value

```go
m := om.New()
v := m.Fetch("first_name", "Simon Peter")
fmt.Println(v) // prints Simon Peter

m.Add("last_name", "Damian")

v = m.Fetch("last_name", "Doe")
fmt.Println(v) // prints Damian
```

### `Index(pos int) (string, interface{})`
> Index return the key/value pair stored at the given index. If the index is out of bound, returns an empty string and `nil`

```go
m := om.New()
m.Add("first_name", "Simon Peter")  
m.Add("last_name", "Damian")

k, v := m.Index(0)
fmt.Println(k, v) // prints first_name Simon Peter

k, v = m.Index(2)
fmt.Println(k, v) // prints last_name Damian

k, v = m.Index(3)
fmt.Println(k, v) // prints  (empty string) nil
```

### `GetKeyIndex(key string) int`

> Returns the numeric index for the given key. If the key is not found, it returns -1

```go
m := om.New()
m.Add("first_name", "Simon Peter")

i := m.GetKeyIndex("first_name")
fmt.Println(i) // prints 0

i = m.GetKeyIndex("last_name")
fmt.Println(i) // prints -1
```

### `ValuesAt(keys ...string) []interface{}`
> Returns a new slice containing the values of the given keys

```go
m := om.New()
m.Add("first_name", "Simon Peter")
m.Add("last_name", "Damian")

v := m.ValuesAt("first_name", "last_name")
fmt.Println(v) // prints [Simon Peter Damian]
```

### `HasKey(keys ...string) bool`
> Returns true if the key is contained in the map

```go
m := om.New()
m.Add("first_name", "Simon Peter")

fmt.Println(m.HasKey("first_name")) // prints true
fmt.Println(m.HasKey("last_name")) // prints false
```

### `HasAny(cb func(key string, value interface{}) bool) bool`
> Returns true if any of the key/value pair satisfies the given callback

```go
m := om.New()
m.Add("first_name", "Simon Peter")

m.HasAny(func(key string, value interface{}) bool {
    return key == "first_name" && value == "Simon Peter"
}) // prints true

m.HasAny(func(key string, value interface{}) bool {
    return key == "last_name" && value == "Damian"
}) // prints false
```

### `Delete(key string) (interface{}, error)`
> Removes the entry with the given key from the map. If the key is not found, it returns an error

```go
m := om.New()
m.Add("first_name", "Simon Peter")

v, err := m.Delete("first_name")
if err != nil {
    panic(err)
}
fmt.Println(v) // prints Simon Peter
```

### `DeleteIF(cb func(key string, value interface{}) bool) *omap`
> Calls the given callback for each key/value pair in the map. If the callback returns true, the key/value pair is removed from the map

```go
m := om.New()
m.Add("first_name", "Simon Peter")

m.DeleteIF(func(key string, value interface{}) bool {
    return key == "first_name" && value == "Simon Peter"
})

fmt.Println(m.HasKey("first_name")) // prints false
```

### `KeepIF(cb func(key string, value interface{}) bool) *omap`
> Calls the callback with each key/value pair. Keeps each entry for which the callback returns true, otherwise deletes the entry from the map

```go
m := om.New()
m.Add("first_name", "Simon Peter")
m.Add("last_name", "Damian")
m.Add("age", 27)

m.KeepIF(func(key string, value interface{}) bool {
    return key == "first_name" && value == "Simon Peter"
})

fmt.Println(m.HasKey("first_name")) // prints true
fmt.Println(m.HasKey("last_name")) // prints false
fmt.Println(m.HasKey("age")) // prints false
```

### `Filter(cb func(key string, value interface{}) bool) *omap`
> Returns a new map containing the entries for which the callback returns true

```go
m := om.New()
m.Add("fish", "tuna")
m.Add("fruit", "apple")
m.Add("vegetable", "carrot")

n := m.Filter(func(key string, value interface{}) bool {
    return key == "fish" || value == "apple"
})

fmt.Println(n.HasKey("fish")) // prints true
fmt.Println(n.HasKey("fruit")) // prints true
fmt.Println(n.HasKey("vegetable")) // prints false
```

### `Filter_(cb func(key string, value interface{}) bool) *omap`
> Modifies the map to contain only the entries for which the callback returns true

```go
m := om.New()
m.Add("fish", "tuna")
m.Add("fruit", "apple")
m.Add("vegetable", "carrot")

m.Filter_(func(key string, value interface{}) bool {
    return key == "fish" || value == "apple"
})

fmt.Println(m.HasKey("fish")) // prints true
fmt.Println(m.HasKey("fruit")) // prints true
fmt.Println(m.HasKey("vegetable")) // prints false
```

### `Slice(keys ...string) *omap`
> Returns a new map containing the entries for the given keys

```go
m := om.New()
m.Add("fish", "tuna")
m.Add("fruit", "apple")
m.Add("vegetable", "carrot")

n := m.Slice("fish", "fruit")

fmt.Println(n.HasKey("fish")) // prints true
fmt.Println(n.HasKey("fruit")) // prints true
fmt.Println(n.HasKey("vegetable")) // prints false
```

### `Compact() *omap`
> Returns a new map with all nil values removed
```go
m := om.New()
m.Add("fish", "tuna")
m.Add("fruit", "apple")
m.Add("vegetable", nil)
m.Add("meat", nil)

n := m.Compact()

fmt.Println(n.HasKey("fish")) // prints true
fmt.Println(n.HasKey("fruit")) // prints true
fmt.Println(n.HasKey("vegetable")) // prints false
fmt.Println(n.HasKey("meat")) // prints false
```

### `Compact_() *omap`
> Modifies the original map with all nil values removed
```go
m := om.New()
m.Add("fish", "tuna")
m.Add("fruit", "apple")
m.Add("vegetable", nil)
m.Add("meat", nil)

m.Compact_()
fmt.Println(m.HasKey("fish")) // prints true
fmt.Println(m.HasKey("fruit")) // prints true
fmt.Println(m.HasKey("vegetable")) // prints false
fmt.Println(m.HasKey("meat")) // prints false
```
