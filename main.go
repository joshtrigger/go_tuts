package main

import (
	"fmt"
	"github.com/joshtrigger/go_tuts/src"
)



func main() {
	age := 90 // shorthand  variable declaration
	const str int = 23
	intArray := [...]int32{12, 43, 39}
	intSlice1 := []int32{10, 20, 30}

	var myMap = map[string]int{"Joshua": 24, "Mike": 32}
	var myCat src.Cat = src.Cat{Name: "Lulu", Age: 4, Owner: src.Owner{OwnerName: "Joshua", OwnerAge: 44}}
	var myDog src.Dog = src.Dog{Name: "Mikey", Age: 7, Owner: src.Owner{OwnerName: "Lugada", OwnerAge: 26}}

	ans, rem, err := src.Divide(age, str)

	fmt.Println("Hello World")
	fmt.Println(src.Add(age, str))
	fmt.Println(len("§")) // len returns number of bytes not number of chars, § returns 2 not 1

	src.MyStrings()

	src.Decide(age, str, ans, rem, err)

	fmt.Println((intArray))
	fmt.Printf("Length is %v, capacity is %v\n", len(intSlice1), cap(intSlice1))
	intSlice1 = append(intSlice1, 50)
	fmt.Printf("Length is %v, capacity is %v\n", len(intSlice1), cap(intSlice1))

	var num, ok = myMap["Mike"] // maps return a second value which is a bool indicating if the key exists
	fmt.Println(num, ok)
	fmt.Println(myCat.Name, myDog.Name)
	fmt.Printf("%v is owned by %v\n", myCat.Name, myCat.GetOwnerName())
	src.PrintName(myDog)
}
