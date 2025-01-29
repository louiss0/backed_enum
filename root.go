package backedEnum

import (
	"fmt"
	"maps"
)

type BackedEnum[
	ValueType interface{ string | int },
	Map ~map[string]ValueType,
] struct {
	structure Map
}

func (e BackedEnum[ValueType, Map]) Structure() Map {

	return maps.Clone(e.structure)

}

func (e BackedEnum[ValueType, Map]) Validate(input ValueType) bool {

	for _, value := range e.Values() {

		if value == input {

			return true
		}

	}

	return false

}

func (e BackedEnum[ValueType, Map]) Parse(input ValueType) error {

	for _, value := range e.Values() {

		if value == input {

			return nil
		}

	}

	return fmt.Errorf("invalid enum value %v", input)

}

func (e BackedEnum[ValueType, Map]) Values() []ValueType {

	slice := []ValueType{}

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

var LoadStatus = BackedEnum[string, loadStatus]{
	structure: loadStatus{
		"ERROR":   "error",
		"SUCCESS": "success",
		"LOADING": "loading",
		"IDLE":    "idle",
	},
}
