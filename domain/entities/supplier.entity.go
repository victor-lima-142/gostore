package entities

import "gorm.io/gorm"

// Supplier represents a supplier of products.
//
// Table name: suppliers
type Supplier struct {
	gorm.Model
	ID            uint              `gorm:"primaryKey;autoIncrement" json:"id"`          // primary key
	Name          string            `gorm:"not null" json:"name"`                        // supplier name
	TaxID         string            `gorm:"not null" json:"tax_id"`                      // supplier tax ID
	FantasyName   string            `json:"fantasy_name"`                                // supplier fantasy name
	Sales         int               `gorm:"not null;default:0" json:"sales"`             // supplier quantity of sales
	QuantityStock int               `gorm:"not null;default:0" json:"quantity_stock"`    // supplier quantity stock
	Products      []ProductSupplier `gorm:"foreignKey:SupplierID" json:"products"`       // One-to-many relationship with ProductSupplier
	Contact       *Contact          `gorm:"foreignKey:CustomerID;unique" json:"contact"` // One-to-one relationship with Contact
}

// TableName overrides the table name used by Supplier to `sales.suppliers`.
func (Supplier) TableName() string {
	return "sales.suppliers"
}
