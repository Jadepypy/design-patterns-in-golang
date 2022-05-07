package main

import (
	"design-patterns-in-golang/singleton/urgent"
	"fmt"
)

func main() {
	singletonInstance := urgent.GetInstance()
	fmt.Println("is empty", singletonInstance.IsEmpty)
	singletonInstance.Fill()
	fmt.Println("is empty", singletonInstance.IsEmpty)

	return
}
