package src

import "fmt"

func Decide(age int, str int, ans int, rem int, err error){
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
}
