package util

import "testing"

func TestIsNil(t *testing.T) {
	var c Cat
	var a Animal = c
	t.Log(IsNil(a))
}

func TestIsNilFixed(t *testing.T) {
	t.Log(IsNilFixed(1))
}

func TestIsNilBetter(t *testing.T) {
	var d *Dog
	//var a Animal = c
	t.Log(IsNilBetter(d))
}
