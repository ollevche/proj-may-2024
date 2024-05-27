package main

import (
	"cafe/cafe"
	"cafe/cafe/kitchen"
	"fmt"
)

func main() {
	menu := map[string]cafe.Item{
		"coffee": {
			Type: cafe.ItemTypeDrink,
			Name: "coffee",
		},
		"borsch": {
			Type: cafe.ItemTypeDish,
			Name: "borsch",
		},
		"donut": {
			Type: cafe.ItemTypeDessert,
			Name: "donut",
		},
	}

	bartender := kitchen.Bartender{}

	kitchen := cafe.NewKitchen()

	kitchen.AddWorker(bartender)

	w := cafe.Waiter{
		Menu:    menu,
		Kitchen: kitchen,
	}

	for {
		var itemName string
		fmt.Scan(&itemName)

		if product, ok := w.ProcessOrder(itemName); ok {
			fmt.Println("Your", product)
			break
		}

		fmt.Println("Try again!")
		// fmt.Println("Bartender created this:", bartender.Cook(menu[itemName]))
	}
}
