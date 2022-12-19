package main

import "fmt"

// productType is a enum type
type productType int

// string enum
const (
	milk productType = iota
	rice
	beer
)

// Unknow state pattern

func (pt productType) String() string {
	switch pt {
	case milk:
		return "milk"
	case rice:
		return "rice"
	case beer:
		return "beer"
	default:
		return "nothing"
	}
}

// func (pt productType) String1() string {
// 	var productsEnumType = [...]string{"milk", "rice", "beer"}
// 	return productsEnumType[pt]
// }

type Store struct {
}

func (s Store) Get() {

}

func (s Store) Set(products []productType) {
	fmt.Println(products)
}

type Human struct {
	weight float32
	height float32
}

type Girl struct {
	Human
}

func (g Girl) goShopping() {

}

type Boy struct {
	Human
}

type myIntInterface interface {
	Get() float32
}

type myInt int

func (_myInt myInt) Get() float32 {
	return float32(_myInt)
}

func main() {
	// fmt.Printf("Value of iotaTest is %v\n", iotaTest)

	// var products = [...]productType{milk, rice, beer}

	// var store = Store{}

	// var sliceProducts = products[0:2] // [0, 2)

	// store.Set(sliceProducts)

	var i = myInt(1)
	fmt.Println(i.Get())
}
