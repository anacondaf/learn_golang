package main

import (
	"fmt"

	"github.com/kainguyen/beginner/calc"
	"github.com/kainguyen/beginner/collection"
	_interface "github.com/kainguyen/beginner/interfaceType"
	"github.com/kainguyen/beginner/iteration"
	"github.com/kainguyen/beginner/rune"
	"github.com/kainguyen/beginner/scope"
	"github.com/kainguyen/beginner/structType"
)

func main() {
	var chuoi string = "10"
	var chuoi1 = "Hello"

	fmt.Printf("Value of chuoi is: %v \n , %v \n", chuoi, chuoi1)

	fmt.Printf("sum = %v \n", calc.Sum2Number(1, 2))
	fmt.Printf("divide = %v \n", calc.Divide(9, 10))

	var funnyString string = "Hello World"
	rune.GetRuneOfString(funnyString)

	//===================COLLECTION=====================//
	fmt.Printf("Length of an array: %v\n", len(collection.Array()))

	slice := collection.SliceLiteral()
	slice = append(slice, 1)
	slice = append(slice, 9)

	for j := 0; j < len(slice); j++ {
		fmt.Printf("Element %v has value %v \n", j, slice[j])
	}

	for i, v := range slice {
		fmt.Printf("Element %v has value %v \n", i, v)
	}

	sliceByMake := collection.SliceByMake()
	fmt.Printf("Len %v:\nCap %v:\n", len(sliceByMake), cap(sliceByMake))

	var mySlice = sliceByMake[:5]

	for _, v := range mySlice {
		fmt.Printf("%v\n", v)
	}

	for _, firstName := range iteration.SeparateFullName() {
		fmt.Printf("%v\n", firstName)
	}

	//========================SCOPE===========================
	scope.Scope()

	//=======================STRUCT===========================//
	var school = structType.University{
		Id:         19,
		SchoolName: "UIT",
		Address:    "123",
	}

	var student1 = structType.Student{
		FirstName:  "Kai",
		LastName:   "Nguyen",
		University: school,
	}

	fmt.Printf("Student %v is attending %v\n", student1.FirstName, student1.University.SchoolName)

	//=======================INTERFACE==========================//
	s1 := _interface.Student{Name: "Nguyen Duc Khai"}
	t1 := _interface.Teacher{Name: "Cao Thị Nhạn"}

	_interface.GetNameOnStruct[_interface.Student](s1)
	_interface.GetNameOnStruct[_interface.Teacher](t1)
}
