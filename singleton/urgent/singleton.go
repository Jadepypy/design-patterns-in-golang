// Package singleton is a demonstration of singleton pattern
package urgent

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
	return uniqueInstance
}

//init creates a boiler instance at the start.
//This is a simple solution to the race condition which naive(lazy) approach has overlooked.
//If the instance is not necessarily needed, creating an instance without the hint of a caller function could lead to
//unnecessary waste. In this case, we should consider using other approaches, ex using lock.
func init() {
	fmt.Println("Initialzing boiler instance...")
	uniqueInstance = &boiler{true}
}
