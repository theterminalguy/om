package om

import "testing"

func TestGet_KeyExist(t *testing.T) {
	m := NewMap()
	m.Add("key", 2)

	want := 2
	got, _ := m.Get("key")
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestGet_KeyNotExist(t *testing.T) {
	m := NewMap()
	_, err := m.Get("invalid key")
	if err == nil {
		t.Error("key not found!")
	}
}
