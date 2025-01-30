// Package backedEnum provides a generic enum implementation with backing values.
// It allows creating type-safe enums with string or integer backing values
// while maintaining compile-time type safety and runtime validation.
package backedEnum

import (
	"fmt"
	"maps"
)

// backedEnum represents a type-safe enumeration with backing values.
//
// Type Parameters:
//   - MapValue: The type of the backing value (~string or ~int)
//   - Map: The map type storing enum key-value pairs
//   - StringOrNumber: The allowed input types (string or int)
type backedEnum[
	MapValue interface{ ~string | ~int },
	Map ~map[string]MapValue,
	StringOrNumber interface{ string | int },
] struct {
	structure Map
}

// NewBackedEnum - creates a backed enum
//
// Type Parameters:
//   - MapValue: The type of the backing value (~string or ~int)
//   - Map: The map type storing enum key-value pairs
//   - StringOrNumber: The allowed input types (string or int)
//
// Example:
//
// ```go
//
//	type status map[string]string
//	var StatusEnum = NewBackedEnum[string, status, string](
//	    structure: Status{
//	        "OK": "ok",
//	        "ERROR": "error",
//	    },
//	}
//
// )
// ```
func NewBackedEnum[
	MapValue interface{ ~string | ~int },
	Map ~map[string]MapValue,
	StringOrNumber interface{ string | int },
](structure Map) backedEnum[MapValue, Map, StringOrNumber] {

	return backedEnum[MapValue, Map, StringOrNumber]{structure}

}

// Structure returns a deep copy of the underlying enum structure.
// This prevents direct modification of the enum's internal state.
func (e backedEnum[MapValue, Map, StringOrNumber]) Structure() Map {

	return maps.Clone(e.structure)

}

// Validate checks if the provided input is a valid enum value.
// It returns true if the input matches any of the enum's backing values.
func (e backedEnum[MapValue, Map, StringOrNumber]) Validate(input StringOrNumber) bool {

	for _, value := range e.Values() {

		if value == any(input) {

			return true
		}

	}

	return false

}

// Parse validates the input value against the enum's valid values.
// It returns nil if the input is valid, or an error describing the invalid input.
func (e backedEnum[MapValue, Map, StringOrNumber]) Parse(input StringOrNumber) error {

	for _, value := range e.Values() {

		if value == any(input) {

			return nil
		}

	}

	return fmt.Errorf("invalid enum value %v", input)

}

// Values returns a slice containing all backing values of the enum.
func (e backedEnum[MapValue, Map, StringOrNumber]) Values() []MapValue {

	slice := []MapValue{}

	structValues := maps.Values(e.structure)

	for value := range structValues {

		slice = append(slice, value)

	}

	return slice

}

// loadStatus represents a mapping of load state names to their string values.
// It provides type-safe access to predefined load states through methods.
type loadStatus map[string]string

// IDLE returns the string value representing the idle state.
func (self loadStatus) IDLE() string {

	return self["IDLE"]

}

// LOADING returns the string value representing the loading state.
func (self loadStatus) LOADING() string {

	return self["LOADING"]

}

// ERROR returns the string value representing the error state.
func (self loadStatus) ERROR() string {

	return self["ERROR"]

}

// SUCCESS returns the string value representing the success state.
func (self loadStatus) SUCCESS() string {

	return self["SUCCESS"]

}

// LoadStatus is a predefined enum representing various states of a loading operation.
// It defines four states: IDLE, LOADING, ERROR, and SUCCESS.
var LoadStatus = backedEnum[string, loadStatus, string]{
	structure: loadStatus{
		"ERROR":   "error",
		"SUCCESS": "success",
		"LOADING": "loading",
		"IDLE":    "idle",
	},
}
