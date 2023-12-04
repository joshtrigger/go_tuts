package src

import (
	"fmt"
	"strings"
)

func MyStrings() {
	var myString = "résumé" // underlying representation is a bytearray, this is where runes come in
	var myString2 = []rune("résumé")
	var myRune = 'b' // we can declare runes using single quotes
	var sample = []string{"l", "u", "g", "a", "d", "a"}
	var strBuilder strings.Builder

	fmt.Println(myString, myRune)
	fmt.Println(myString[0])                         // prints out unicode code char number not the char
	fmt.Printf("%v, %T\n", myString[1], myString[1]) // print value and type
	fmt.Printf("%v, %T\n", myString2[1], myString2[1])
	fmt.Printf("Length of String1: %v, Length of String2: %v\n", len(myString), len(myString2)) // string1 is 8 while string2 is 6

	for i := range sample {
		strBuilder.WriteString(sample[i])
	}

	var concat = strBuilder.String() // use string builder to concatenate strings
	fmt.Println(concat)
}
