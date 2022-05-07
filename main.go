package main

import (
	"design-patterns-in-golang/singleton"
	"fmt"
)

func main() {
	singletonInstance := singleton.GetInstance()
	fmt.Println("is empty", singletonInstance.IsEmpty)
	singletonInstance.Fill()
	fmt.Println("is empty", singletonInstance.IsEmpty)
	if singleton.GetInstance() == nil {
		return
	}
	return
}
