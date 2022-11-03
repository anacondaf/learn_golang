package rune

import "fmt"

func GetRuneOfString(myString string) {
	var myRune = []rune(myString)

	fmt.Printf("%c = %v\n", myRune[0], myRune)
}

//Access modifier: Public: PascalCase; Private: camelCase
