package util

import (
	"fmt"
	"reflect"
)

type Animal interface {
	MakeSound() string
}

type Dog struct{}

func (d *Dog) MakeSound() string {
	return "Bark"
}

type Cat struct {
}

func (c Cat) MakeSound() string {
	return "Meow"
}

func IsNil(i interface{}) bool {
	fmt.Println(i, i == nil, reflect.ValueOf(i).IsNil())
	return i == nil || reflect.ValueOf(i).IsNil()
}

func IsNilFixed(i interface{}) bool {
	if i == nil {
		return true
	}
	switch reflect.TypeOf(i).Kind() {
	case reflect.Ptr, reflect.Map, reflect.Slice, reflect.Chan, reflect.Array:
		return reflect.ValueOf(i).IsNil()
	}
	return false
}

func IsNilBetter(i Animal) bool {
	var ret bool
	switch i.(type) {
	case *Dog:
		v := i.(*Dog)
		ret = v == nil
	case Cat:
		ret = false
	}
	return ret
}
