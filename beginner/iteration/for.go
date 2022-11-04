package iteration

import (
	"strings"

	"github.com/kainguyen/beginner/collection"
)

func SeparateFullName() []string {
	var studentList = []string{"Khai Nguyen", "Doni Tran", "Minh Tran"}

	var firstNameList = []string{}

	studentList = collection.StudentFullNameSlice(studentList)

	for _, student := range studentList {
		firstNameList = append(firstNameList, strings.Fields(student)[0])
	}

	return firstNameList
}
