package main

import "fmt"

type PizzaStore struct {
}

func (ps *PizzaStore) orderPizza(flavor string) {
	var pizza PizzaInterface
	var err error
	if pizza, err = createPizza(flavor); err == nil {
		pizza.bake()
		pizza.box()
	} else {
		fmt.Println(err.Error())
	}

	return
}
