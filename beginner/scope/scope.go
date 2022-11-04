package scope

import "fmt"

func Scope() {
	var number int = 5;
	fmt.Printf("Scope function: %v\n", number)
	fmt.Printf("Scope function: %v\n", greeting)
	fmt.Printf("Scope function: %v\n", Speech)

	{
		number := 10
		fmt.Printf("In room 1: %v\n", number)

		fmt.Printf("In room 1: %v\n", number)
	}
	
	if true {
		fmt.Printf("In If scope: %v\n", number)
	}
}

// func outsideScopeFunc() {
// 	fmt.Print(number)
// }