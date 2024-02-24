package arrayops

import (
	"fmt"
	"reflect"
)

const ErrInvalidType = "\n>>> Error: The expected data structure is slice or an array, but got:\n---> %v.\n"

type Array interface{}

type ArrayInstance struct {
	Array
}

func New(array Array) *ArrayInstance {
	return &ArrayInstance{Array: array}
}

func (instance *ArrayInstance) _isValid() reflect.Value {
	arrayValue := reflect.ValueOf(instance.Array)

	if arrayValue.Kind() != reflect.Slice && arrayValue.Kind() != reflect.Array {
		panic(fmt.Sprintf(
			ErrInvalidType,
			reflect.TypeOf(instance.Array).Kind(),
		))
	}

	return arrayValue
}

func (instance *ArrayInstance) IndexOf(value interface{}) int {
	arrayValue := instance._isValid()
	valueValue := reflect.ValueOf(value)

	for index := 0; index < arrayValue.Len(); index++ {
		if arrayValue.Index(index).Interface() == valueValue.Interface() {
			return index
		}
	}

	return -1
}

func (instance *ArrayInstance) Includes(value interface{}) bool {
	return instance.IndexOf(value) != -1
}

func (instance *ArrayInstance) FindIndex(predicate func(int) bool) int {
	arrayValue := instance._isValid()

	for index := 0; index < arrayValue.Len(); index++ {
		if predicate(index) {
			return index
		}
	}

	return -1
}

func (instance *ArrayInstance) Find(predicate func(interface{}) bool) interface{} {
	arrayValue := instance._isValid()

	for index := 0; index < arrayValue.Len(); index++ {
		val := arrayValue.Index(index).Interface()
		if predicate(val) {
			return val
		}
	}

	return nil
}

func (instance *ArrayInstance) All(predicate func(interface{}) bool) bool {
	arrayValue := instance._isValid()

	for index := 0; index < arrayValue.Len(); index++ {
		if !predicate(arrayValue.Index(index).Interface()) {
			return false
		}
	}

	return true
}
