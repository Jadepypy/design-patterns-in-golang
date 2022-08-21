package main

import (
	"errors"
)

func createPizza(flavor string) (PizzaInterface, error) {
	switch flavor {
	case "New York":
		return newNewYorkPizza(), nil
	case "Chicago":
		return newChicagoPizza(), nil
	default:
		return nil, errors.New("No such flavor")
	}
}
