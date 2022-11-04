package collection

func SliceLiteral() []uint8 {
	slice := []uint8{}

	return slice
}

func SliceByMake() [10]int {
	var sliceByMake [10]int = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	return sliceByMake
}

func StudentFullNameSlice(studentList []string) []string {
	var students = studentList[:]
	return students
}
