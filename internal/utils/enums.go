package utils

import (
	"fmt"
	"reflect"
)

// Enum is an interface that represents an enumeration
type Enum interface {
	String() string
	EnumIndex() int
}

// EnumValue is a struct that implements the Enum interface
type EnumValue struct {
	index int
	name  string
}

func (e EnumValue) String() string {
	return e.name
}

func (e EnumValue) EnumIndex() int {
	return e.index
}

// NewEnum creates a new enum type with the given names
func NewEnum(names ...string) []EnumValue {
	values := make([]EnumValue, len(names))
	for i, name := range names {
		values[i] = EnumValue{index: i + 1, name: name}
	}
	return values
}

// EnumFromString returns the enum value for a given string
func EnumFromString(enumType interface{}, s string) (Enum, error) {
	v := reflect.ValueOf(enumType)
	for i := 0; i < v.NumField(); i++ {
		if v.Field(i).Interface().(Enum).String() == s {
			return v.Field(i).Interface().(Enum), nil
		}
	}
	return nil, fmt.Errorf("invalid enum value: %s", s)
}
