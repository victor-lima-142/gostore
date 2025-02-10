package entities

import (
	"time"

	"gorm.io/gorm"
)

// Order represents an order placed by a customer.
//
// Table name: orders
type Order struct {
	gorm.Model
	ID            uint                   `gorm:"primaryKey;autoIncrement" json:"id"`       // primary key
	CustomerID    uint                   `gorm:"not null" json:"customer_id"`              // foreign key for Customer
	OrderDate     time.Time              `gorm:"not null" json:"order_date"`               // order date for the order
	DeliveryDate  time.Time              `gorm:"not null" json:"delivery_date"`            // delivery date for the order
	DeliveryOrder bool                   `gorm:"not null" json:"delivery_order"`           // delivery order for the order
	Discount      float32                `gorm:"not null;default:0" json:"discount"`       // discount for the order
	UKOrderNumber string                 `gorm:"not null" json:"uk_order_number"`          // uk order number for the order
	OrderProducts []OrderProductSupplier `gorm:"foreignKey:OrderID" json:"order_products"` // one-to-many relationship with OrderProductSupplier
}

// TableName overrides the table name used by Order to `sales.orders`.
func (Order) TableName() string {
	return "sales.orders"
}
