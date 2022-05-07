// Package singleton is a demonstration of singleton pattern
package singleton

import "fmt"

// a boiler should be unique, since we only have one boiler to boil milk
var uniqueInstance *boiler

type boiler struct {
	IsEmpty bool
}

// Fill fills up the boiler with milk
func (b *boiler) Fill() {
	if b.IsEmpty {
		b.IsEmpty = false
	} else {
		fmt.Println("Boiler is already filled!")
	}

}

// GetInstance is the key function of singleton pattern
// It returns a boiler instance and makes sure that there is only one existing instance.
func GetInstance() *boiler {
	if uniqueInstance == nil {
		uniqueInstance = &boiler{true}
	}
	return uniqueInstance
}
