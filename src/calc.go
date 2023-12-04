package src

import "errors"

func Add(x int, y int) int {
	return x + y
}

func Divide(num1 int, num2 int) (int, int, error) {
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
