package main

import (
	"errors"
	"fmt"
)

func main() {
	age := 90 // shorthand  variable declaration
	const str int = 23

	ans, rem, err := divide(age, str)

	fmt.Println("Hello World")
	fmt.Println(add(age, str))
	fmt.Println(len("§")) // len returns number of bytes in a string not number of chars

	if err != nil {
		fmt.Println(err.Error())
	} else if rem == 0 {
		fmt.Printf("%v divided by %v is %v\n", age, str, ans)
	} else {
		fmt.Printf("%v divided by %v is %v remainder %v\n", age, str, ans, rem) // string formatting (does not add new line)
	}

	// switch statements do not need break keyword
	switch {
	case err != nil:
		fmt.Println(err.Error())
	case rem == 0:
		fmt.Printf("%v divided by %v is %v\n", age, str, ans)
	default:
		fmt.Printf("%v divided by %v is %v remainder %v\n", age, str, ans, rem) // string formatting (does not add new line)
	}
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
