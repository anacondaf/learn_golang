package _interface

import (
	"fmt"
)

type Student struct {
	Name string
	Id   string
}

type Teacher struct {
	Name string
	Id   string
}

type Getter[T any] interface {
	Get() T
}

// Golang method
func (s Student) Get() Student {
	return s
}

func (t Teacher) Get() Teacher {
	return t
}

func GetNameOnStruct[T any](sgs Getter[T]) {
	fmt.Printf("%+v\n", sgs.Get())
}
