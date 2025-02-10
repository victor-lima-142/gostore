package entities

import "gorm.io/gorm"

// Product represents a product sold by a supplier.
//
// Table name: products
type Product struct {
	gorm.Model
	ID          uint              `gorm:"primaryKey;autoIncrement" json:"id"`    // primary key
	Name        string            `gorm:"not null" json:"name"`                  // general name of the product
	Code        string            `gorm:"not null" json:"code"`                  // general code of the product
	Sales       int               `gorm:"not null;default:0" json:"sales"`       // total sales of the product
	MarketValue float32           `json:"market_value"`                          // default market value of product for current market (EMC)
	Suppliers   []ProductSupplier `gorm:"foreignKey:ProductID" json:"suppliers"` // many-to-many relationship with Supplier
}

// TableName overrides the table name used by Product to `sales.products`.
func (Product) TableName() string {
	return "sales.products"
}
