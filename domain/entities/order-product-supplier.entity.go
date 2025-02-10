package entities

import "gorm.io/gorm"

// OrderProductSupplier represents the association between orders, products, and suppliers.
//
// Table name: order_product_suppliers
type OrderProductSupplier struct {
	gorm.Model                // Adds ID, CreatedAt, UpdatedAt, DeletedAt
	ID                uint    `gorm:"primaryKey;autoIncrement" json:"id"`  // primary key
	OrderID           uint    `gorm:"not null" json:"order_id"`            // foreign key for Order
	ProductSupplierID uint    `gorm:"not null" json:"product_supplier_id"` // foreign key for ProductSupplier
	Value             float32 `gorm:"not null" json:"value"`               // value of the product of a supplier for this specific order
	Discount          float32 `gorm:"not null;default:0" json:"discount"`  // discount of the product of a supplier for this specific order
}

// TableName overrides the table name used by OrderProductSupplier to `sales.order_product_suppliers`.
func (OrderProductSupplier) TableName() string {
	return "sales.order_product_suppliers"
}
