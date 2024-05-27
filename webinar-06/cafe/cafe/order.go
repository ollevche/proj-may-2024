package cafe

import "math/rand"

type ItemType string

const (
	ItemTypeDrink   ItemType = "drink"
	ItemTypeDish    ItemType = "dish"
	ItemTypeDessert ItemType = "dessert"
)

type Item struct {
	Type ItemType
	Name string
}

type Order struct {
	ID   int
	Item Item
}

func NewOrder(i Item) Order {
	return Order{
		ID:   rand.Intn(100),
		Item: i,
	}
}

type Product string

type KitchenWorker interface {
	CanCook() ItemType
	Cook(Item) Product
}

type Kitchen map[ItemType]KitchenWorker

func NewKitchen() Kitchen {
	return make(Kitchen)
}

func (k Kitchen) AddWorker(w KitchenWorker) {
	k[w.CanCook()] = w
}

type Waiter struct {
	Menu    map[string]Item
	Kitchen Kitchen
}

func (w Waiter) ProcessOrder(itemName string) (Product, bool) {
	item, ok := w.Menu[itemName]
	if !ok {
		return "", false
	}

	order := NewOrder(item)

	worker, ok := w.Kitchen[order.Item.Type]
	if !ok {
		return "", false
	}

	product := worker.Cook(order.Item)

	return product, true
}
