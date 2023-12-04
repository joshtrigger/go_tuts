package src

import "fmt"

type Dog struct {
	Name string
	Age  uint8
	Owner
}

type Cat struct {
	Name string
	Age  uint8
	Owner
}

type Owner struct {
	OwnerName string
	OwnerAge  uint8
}

type Pet interface {
	GetOwnerName() string
}

func PrintName(entity Pet) {
	fmt.Println(entity.GetOwnerName())
}


func (c Cat) GetOwnerName() string {
	return c.OwnerName
}

func (d Dog) GetOwnerName() string {
	return d.OwnerName
}
