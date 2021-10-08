package om

import "testing"

func TestGet_KeyExist(t *testing.T) {
	m := New()
	m.Add("key", 2)

	want := 2
	got, _ := m.Get("key")
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestGet_KeyNotExist(t *testing.T) {
	m := New()
	_, err := m.Get("nokey")
	if err == nil {
		t.Error("key not found!")
	}
}

func TestFetch_KeyExist(t *testing.T) {
	m := New()
	m.Add("key", 2)

	want := 2
	got := m.Fetch("key", nil)
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestFetch_KeyNotExist(t *testing.T) {
	m := New()

	want := "rabbit"
	got := m.Fetch("nokey", "rabbit")
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
