package arrayops

import (
	"reflect"
	"strings"
	"testing"
)

func testErrInvalidType(unit *testing.T, array Array) {
	defer func() {
		if r := recover(); r == nil {
			unit.Errorf("\n\n>>> The function did not panic, but it should.")
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
		t.Errorf(
			"\n\n>>> Expected %d but got %d.\n\n",
			1,
			instance.IndexOf(2),
		)
	}

	if instance.IndexOf(4) != -1 {
		t.Errorf(
			"\n\n>>> Expected %d but got %d.\n\n",
			-1,
			instance.IndexOf(4),
		)
	}
}

func TestIncludes(t *testing.T) {
	instance := New([]string{"foo", "bar"})
	if !instance.Includes("foo") {
		t.Errorf(
			"\n\n>>> Expected %t but got %t.\n\n",
			true,
			instance.Includes("foo"),
		)
	}

	if instance.Includes("baz") {
		t.Errorf(
			"\n\n>>> Expected %t but got %t.\n\n",
			false,
			instance.Includes("baz"),
		)
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
		t.Errorf(
			"\n\n>>> Expected %d but got %d.\n\n",
			expected,
			got,
		)
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
		t.Errorf(
			"\n\n>>> Expected %t but got %t.\n\n",
			expected,
			got,
		)
	}

	instance = New([]int{5, 4, 3})
	expected = false
	got = instance.All(predicate)
	if got != expected {
		t.Errorf(
			"\n\n>>> Expected %t but got %t.\n\n",
			expected,
			got,
		)
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
		t.Errorf(
			"\n\n>>> Expected %d but got %d.\n\n",
			2,
			instance.FindIndex(predicate),
		)
	}
}

func TestMap(t *testing.T) {
	instance := New([]int{1, 2, 3, 4})
	double := func(i interface{}) interface{} {
		return i.(int) * 2
	}

	result := instance.Map(double)
	expected := []int{2, 4, 6, 8}
	if !reflect.DeepEqual(result.Array, expected) {
		t.Errorf(
			"\n\n>>> Expected %v but got %v\n\n",
			expected,
			result.Array,
		)
	}

	instance = New([]string{"hello", "world"})
	upper := func(i interface{}) interface{} {
		return strings.ToUpper(i.(string))
	}

	result = instance.Map(upper)
	expectedStr := []string{"HELLO", "WORLD"}
	if !reflect.DeepEqual(result.Array, expectedStr) {
		t.Errorf(
			"\n\n>>> Expected %v but got %v\n\n",
			expectedStr,
			result.Array,
		)
	}
}

func TestFilter(t *testing.T) {
	instance := New([]int{1, 2, 3, 4, 5, 6})
	isEven := func(i interface{}) bool {
		return i.(int)%2 == 0
	}

	result := instance.Filter(isEven)
	expected := []int{2, 4, 6}
	if !reflect.DeepEqual(result.Array, expected) {
		t.Errorf(
			"\n\n>>> Expected %v but got %v\n\n",
			expected,
			result.Array,
		)
	}

	instance = New([]string{
		"apple",
		"banana",
		"cherry",
		"avocado",
	})
	startsWithA := func(i interface{}) bool {
		return strings.HasPrefix(i.(string), "a")
	}

	result = instance.Filter(startsWithA)
	expectedStr := []string{"apple", "avocado"}
	if !reflect.DeepEqual(result.Array, expectedStr) {
		t.Errorf(
			"\n\n>>> Expected %v but got %v\n\n",
			expectedStr,
			result.Array,
		)
	}
}

func TestReduce(t *testing.T) {
	instance := New([]int{1, 2, 3, 4})
	sum := func(acc interface{}, val interface{}) interface{} {
		return acc.(int) + val.(int)
	}

	result := instance.Reduce(sum, 0)
	if result != 10 {
		t.Errorf(
			"\n\n>>> Expected %d but got %d\n\n",
			10,
			result,
		)
	}
}
