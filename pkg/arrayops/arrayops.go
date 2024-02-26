package arrayops

import (
	"fmt"
	"reflect"
)

type Array interface{}

type ArrayInstance struct {
	Array
}

func New(array Array) *ArrayInstance {
	return &ArrayInstance{Array: array}
}

func (instance *ArrayInstance) _isValid() reflect.Value {
	arrayValue := reflect.ValueOf(instance.Array)

	invalidTypeErrMsg := "\n\n>>> Error: The expected data structure is " +
		"slice or an array, but got:\n---> %v.\n\n"

	if arrayValue.Kind() != reflect.Slice && arrayValue.Kind() != reflect.Array {
		panic(fmt.Sprintf(
			invalidTypeErrMsg,
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

func (instance *ArrayInstance) Map(transform func(interface{}) interface{}) *ArrayInstance {
	arrayValue := instance._isValid()
	result := reflect.MakeSlice(
		arrayValue.Type(),
		arrayValue.Len(),
		arrayValue.Cap(),
	)

	for i := 0; i < arrayValue.Len(); i++ {
		result.Index(i).Set(
			reflect.ValueOf(
				transform(arrayValue.Index(i).Interface()),
			),
		)
	}

	return New(result.Interface())
}

func (instance *ArrayInstance) Filter(predicate func(interface{}) bool) *ArrayInstance {
	arrayValue := instance._isValid()
	result := reflect.MakeSlice(
		arrayValue.Type(),
		0,
		arrayValue.Cap(),
	)

	for i := 0; i < arrayValue.Len(); i++ {
		val := arrayValue.Index(i).Interface()
		if predicate(val) {
			result = reflect.Append(
				result,
				reflect.ValueOf(val),
			)
		}
	}

	return New(result.Interface())
}

func (instance *ArrayInstance) Reduce(
	accumulator func(interface{}, interface{}) interface{},
	initialValue interface{},
) interface{} {
	arrayValue := instance._isValid()
	result := initialValue

	for i := 0; i < arrayValue.Len(); i++ {
		result = accumulator(
			result,
			arrayValue.Index(i).Interface(),
		)
	}

	return result
}
