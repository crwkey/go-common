package pattern

import "fmt"

type WithName struct {
	Name string
}

type Country struct {
	WithName
}

type City struct {
	WithName
}

func (w WithName) PrintStr() {
	fmt.Println(w.Name)
}

type Shape interface {
	Sides() int
	Area() int
}

type Square struct {
	Len int
}

func (s Square) Sides() int {
	return 4
}

func (s Square) Area() int {
	return s.Len
}

func MapUpCase(arr []string, fn func(s string) string) []string {
	var newArray []string

	for _, it := range arr {
		newArray = append(newArray, fn(it))
	}
	return newArray
}
