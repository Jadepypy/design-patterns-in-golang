package main

import "fmt"

type PizzaInterface interface {
	bake()
	box()
}

type Pizza struct {
	name string
}

func (p Pizza) bake() {
	fmt.Printf("%s Pizza is baking!\n", p.name)
}

func (p Pizza) box() {
	fmt.Printf("%s Pizza is boxing!\n", p.name)
}

type NewYorkPizza struct {
	Pizza
}

// which one is a better approach
func newNewYorkPizza() PizzaInterface {
	return &Pizza{
		name: "New York",
	}
}

type ChicagoPizza struct {
	Pizza
}

func newChicagoPizza() PizzaInterface {
	return &ChicagoPizza{
		Pizza: Pizza{
			name: "Chicago",
		},
	}
}
