# om
Ordered Map

## Installation

```bash
$ go get github.com/theterminalguy/om
```

## API docs

### New 
> New creates a new ordered map

```go
m := om.New()
```

### Add
> Add adds a key value pair to the map. If the key already exist, it's value is updated

```go
m := om.New()
m.Add("first_name", "Simon Peter")
m.Add("last_name", "Damian")
m.Add("age", 27)
m.Add("month", 27)
m.Add("sex", "male")
m.Add("married", false)

m.Add("age", 28) // updates the age to 28 because key exists
``` 

### Set 
> Updates a key's value using Put. If the key is not found, it calls Add

```go
m := om.New() 
m.Set("first_name", "Simon Peter") // calls Add
```

### Put
> Updates a key's value. If the key is not found, it returns an error

```go
m := om.New()
m.Put("first_name", "Simon Peter") // returns an error

m.Add("first_name", "Simon Peter")

m.Put("first_name", "Simon") // updates the value of first_name to Simon without error
```

### Get
> Get returns the value of a key. If the key is not found, it returns an error

```go
m := om.New()
m.Add("first_name", "Simon Peter")

v, err := m.Get("first_name")
if err != nil {
    panic(err)
}

fmt.Println(v) // prints Simon Peter
```
