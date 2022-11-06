package structType

type University struct {
	Id 		   int32
	SchoolName string
	Address    string
}

type Student struct {
	FirstName 	string
	LastName  	string
	Age 		byte
	University 	University
}