package order

import "time"

// Order describes basic client orders.
type Order struct {
	ID       int
	Date     time.Time
	Item     string
	Category string
}

// Creator is a wrapper for order create method.
type Creator interface {
	CreateOrder(Date time.Time, Item, Category string) error
}

// Remover is a wrapper for order remove method.
type Remover interface {
	RemoveOrder(ID int) error
}

// Updater is a wrapper for order update method.
type Updater interface {
	UpdateOrder(new, old Order) error
}

// Finder is a wrapper for finding method.
type Finder interface {
	FindOrder(ID int) (Order, error)
}
