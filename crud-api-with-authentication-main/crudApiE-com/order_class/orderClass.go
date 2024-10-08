package order_class

import (
	"gorm.io/gorm"
	"time"
)

// Order: ID, UserID (Foreign Key), OrderDate, Status, TotalAmount.
// OrderItem: ID, OrderID (Foreign Key), ProductID (Foreign Key), Quantity, Price.

type Order struct {
	OrderId         uint `gorm:"primaryKey"`
	UserId          uint
	ProductId       uint
	Product         string
	OrderDate       time.Time
	OrderQuantity   int
	OrderTotalPrice float64
	OrderCreatedAt  time.Time
	OrderUpdatedAt  time.Time
	OrderDeletedAt  gorm.DeletedAt `gorm:"index"`
}
