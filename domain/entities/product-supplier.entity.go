package entities

import "gorm.io/gorm"

// ProductSupplier represents the association between products and suppliers (offering specific products).
//
// Table name: product_suppliers
type ProductSupplier struct {
	gorm.Model
	ID                  uint                   `gorm:"primaryKey;autoIncrement" json:"id"`
	ProductID           uint                   `gorm:"not null" json:"product_id"`  // Foreign key for Product
	SupplierID          uint                   `gorm:"not null" json:"supplier_id"` // Foreign key for Supplier
	Cost                float32                `gorm:"not null" json:"cost"`
	Value               float32                `gorm:"not null" json:"value"`
	Quantity            int                    `gorm:"not null" json:"quantity"`
	SupplierProductCode string                 `json:"supplier_product_code"`
	SupplierProductName string                 `json:"supplier_product_name"`
	Sales               int                    `gorm:"not null;default:0" json:"sales"`
	OrderProducts       []OrderProductSupplier `gorm:"foreignKey:ProductSupplierID" json:"order_products"` // One-to-many relationship with OrderProductSupplier
}

// TableName overrides the table name used by ProductSupplier to `sales.product_suppliers`.
func (ProductSupplier) TableName() string {
	return "sales.product_suppliers"
}
