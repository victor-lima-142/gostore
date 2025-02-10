package entities

import (
	"time"

	"gorm.io/gorm"
)

// Customer represents a customer in the system.
//
// Table name: customers
type Customer struct {
	gorm.Model
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`          // primary key
	FirstName string    `gorm:"not null" json:"first_name"`                  // first name of customer
	LastName  string    `gorm:"not null" json:"last_name"`                   // last name of customer
	Birthday  time.Time `gorm:"not null" json:"birthday"`                    // birthday of customer
	TaxID     string    `gorm:"not null" json:"tax_id"`                      // tax id of customer
	Orders    []Order   `gorm:"foreignKey:CustomerID" json:"orders"`         // One-to-many relationship with Order
	Contact   *Contact  `gorm:"foreignKey:CustomerID;unique" json:"contact"` // One-to-one relationship with Contact
}

// TableName overrides the table name used by Customer to `sales.customers`.
func (Customer) TableName() string {
	return "sales.customers"
}
