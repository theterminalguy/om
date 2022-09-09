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

