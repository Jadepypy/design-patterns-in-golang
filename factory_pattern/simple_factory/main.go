package main

func main() {
	pizzaStore := &PizzaStore{}
	// order New York pizza
	pizzaStore.orderPizza("New York")
	// order Chicago pizza
	pizzaStore.orderPizza("Chicago")
	// order unknown flavor
	pizzaStore.orderPizza("unknown")
}
