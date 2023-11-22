package main

import (
	"errors"
	"fmt"
	"strings"
)

func main() {
	age := 90 // shorthand  variable declaration
	const str int = 23
	intArray := [...]int32{12, 43, 39}
	intSlice1 := []int32{10, 20, 30}

	var myMap = map[string]int{"Joshua": 24, "Mike": 32}

	ans, rem, err := divide(age, str)

	fmt.Println("Hello World")
	fmt.Println(add(age, str))
	fmt.Println(len("§")) // len returns number of bytes in a string not number of chars

	myStrings()

	if err != nil {
		fmt.Println(err.Error())
	} else if rem == 0 {
		fmt.Printf("%v divided by %v is %v\n", age, str, ans)
	} else {
		fmt.Printf("%v divided by %v is %v remainder %v\n", age, str, ans, rem) // string formatting (does not add new line)
	}

	// switch statements do not need break keyword and can be conditional
	switch {
	case err != nil:
		fmt.Println(err.Error())
	case rem == 0:
		fmt.Printf("%v divided by %v is %v\n", age, str, ans)
	default:
		fmt.Printf("%v divided by %v is %v remainder %v\n", age, str, ans, rem) // string formatting (does not add new line)
	}

	switch rem {
	case 0:
		fmt.Println("Division has no remainder")
	default:
		fmt.Printf("Division has a remainder of %v\n", rem)
	}

	fmt.Println((intArray))
	fmt.Printf("Length is %v, capacity is %v\n", len(intSlice1), cap(intSlice1))
	intSlice1 = append(intSlice1, 50)
	fmt.Printf("Length is %v, capacity is %v\n", len(intSlice1), cap(intSlice1))

	var num, ok = myMap["Mike"] // maps return a second which is a bool indicating if the key exists
	fmt.Println(num, ok)
}

func add(x int, y int) int {
	return x + y
}

func divide(num1 int, num2 int) (int, int, error) {
	// function that returns more than one value
	var err error
	if num2 == 0 {
		err = errors.New("cannot divide by zero")
		return 0, 0, err
	}

	number := num1 / num2
	remainder := num1 % num2

	return number, remainder, err
}

func myStrings() {
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

	var concat = strBuilder.String() // use string builder concatenate strings
	fmt.Println(concat)
}
