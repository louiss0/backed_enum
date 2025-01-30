package backedEnum

import (
	"fmt"
	"maps"
)

type BackedEnum[
	MapValue interface{ ~string | ~int },
	Map ~map[string]MapValue,
	StringOrNumber interface{ string | int },
] struct {
	structure Map
}

func (e BackedEnum[MapValue, Map, StringOrNumber]) Structure() Map {

	return maps.Clone(e.structure)

}

func (e BackedEnum[MapValue, Map, StringOrNumber]) Validate(input StringOrNumber) bool {

	for _, value := range e.Values() {

		if value == any(input) {

			return true
		}

	}

	return false

}

func (e BackedEnum[MapValue, Map, StringOrNumber]) Parse(input StringOrNumber) error {

	for _, value := range e.Values() {

		if value == any(input) {

			return nil
		}

	}

	return fmt.Errorf("invalid enum value %v", input)

}

func (e BackedEnum[MapValue, Map, StringOrNumber]) Values() []MapValue {

	slice := []MapValue{}

	structValues := maps.Values(e.structure)

	for value := range structValues {

		slice = append(slice, value)

	}

	return slice

}

type loadStatus map[string]string

func (self loadStatus) IDLE() string {

	return self["IDLE"]

}

func (self loadStatus) LOADING() string {

	return self["LOADING"]

}

func (self loadStatus) ERROR() string {

	return self["ERROR"]

}

func (self loadStatus) SUCCESS() string {

	return self["SUCCESS"]

}

var LoadStatus = BackedEnum[string, loadStatus, string]{
	structure: loadStatus{
		"ERROR":   "error",
		"SUCCESS": "success",
		"LOADING": "loading",
		"IDLE":    "idle",
	},
}
