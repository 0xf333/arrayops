package arrayops

import (
	"testing"
)

func testErrInvalidType(unit *testing.T, array Array) {
	defer func() {
		if r := recover(); r == nil {
			unit.Errorf(">>> The function did not panic, but it should.")
		}
	}()

	instance := New(array)
	instance._isValid()
}

func TestErrInvalidType(t *testing.T) {
	testErrInvalidType(t, 123)
	testErrInvalidType(t, "abc")
	testErrInvalidType(t, make(map[string]int))
}

func TestIndexOf(t *testing.T) {
	instance := New([]int{1, 2, 3})
	if instance.IndexOf(2) != 1 {
		t.Errorf(">>> Expected %d but got %d.", 1, instance.IndexOf(2))
	}

	if instance.IndexOf(4) != -1 {
		t.Errorf(">>> Expected %d but got %d.", -1, instance.IndexOf(4))
	}
}

func TestIncludes(t *testing.T) {
	instance := New([]string{"foo", "bar"})
	if !instance.Includes("foo") {
		t.Errorf(">>> Expected %t but got %t.", true, instance.Includes("foo"))
	}

	if instance.Includes("baz") {
		t.Errorf(">>> Expected %t but got %t.", false, instance.Includes("baz"))
	}
}

func TestFind(t *testing.T) {
	instance := New([]int{4, 7, 1, 9})
	predicate := func(i interface{}) bool {
		return i.(int) > 5
	}

	expected := 7
	got := instance.Find(predicate)
	if got != expected {
		t.Errorf(">>> Expected %d but got %d.", expected, got)
	}
}

func TestAll(t *testing.T) {
	instance := New([]int{6, 8, 9})
	predicate := func(i interface{}) bool {
		return i.(int) > 5
	}

	expected := true
	got := instance.All(predicate)
	if got != expected {
		t.Errorf(">>> Expected %t but got %t.", expected, got)
	}

	instance = New([]int{5, 4, 3})
	expected = false
	got = instance.All(predicate)
	if got != expected {
		t.Errorf(">>> Expected %t but got %t.", expected, got)
	}
}

func _isPrime(n int) bool {
	if n <= 2 {
		return n == 2
	}

	if n%2 == 0 {
		return false
	}

	for i := 3; i*i <= n; i += 2 {
		if n%i == 0 {
			return false
		}
	}

	return true
}

func TestFindIndex(t *testing.T) {
	instance := New([]int{4, 6, 7, 10})
	predicate := func(i int) bool {
		return _isPrime(i)
	}

	if instance.FindIndex(predicate) != 2 {
		t.Errorf(">>> Expected %d but got %d.", 2, instance.FindIndex(predicate))
	}
}
