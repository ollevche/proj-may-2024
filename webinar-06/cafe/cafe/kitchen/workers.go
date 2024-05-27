package kitchen

import "cafe/cafe"

type Bartender struct{}

func (b Bartender) CanCook() cafe.ItemType {
	return cafe.ItemTypeDrink
}

func (b Bartender) Cook(item cafe.Item) cafe.Product {
	if item.Type != cafe.ItemTypeDrink {
		return "Something awful"
	}

	return cafe.Product("Delicious " + item.Name)
}
