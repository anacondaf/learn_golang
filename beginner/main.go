package main

import (
	"fmt"

	c "github.com/kainguyen/beginner/calc"
	"github.com/kainguyen/beginner/collection"
	r "github.com/kainguyen/beginner/rune"
)

func main() {
	var chuoi string = "10"
	var chuoi1 = "Hello"

	fmt.Printf("Value of chuoi is: %v \n , %v \n", chuoi, chuoi1)

	fmt.Printf("sum = %v \n", c.Sum2Number(1, 2))
	fmt.Printf("divide = %v \n", c.Divide(9, 10))

	var funnyString string = "Hello World"
	r.GetRuneOfString(funnyString)

	//===================COLLECTION=====================//
	fmt.Printf("Length of an array: %v\n", len(collection.Array()))

	slice := collection.Slice()
	slice = append(slice, 1)
	slice = append(slice, 9)

	for j := 0; j < len(slice); j++ {
		fmt.Printf("Element %v has value %v \n", j, slice[j])
	}

	for i, v := range slice {
		fmt.Printf("Element %v has value %v \n", i, v)
	}

}
