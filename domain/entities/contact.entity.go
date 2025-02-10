package entities

import "gorm.io/gorm"

// Contact represents contact information for a customer or supplier.
//
// Table name: contacts
type Contact struct {
	gorm.Model
	ID             uint   `gorm:"primaryKey;autoIncrement" json:"id"` // primary key
	Phone          string `gorm:"not null" json:"phone"`              // phone number of the contact
	SecondaryPhone string `json:"secondary_phone"`                    // secondary phone number of the contact
	PostalCode     string `gorm:"not null" json:"postal_code"`        // postal code of the contact for address
	Area           string `gorm:"not null" json:"area"`               // area of the contact for address
	District       string `gorm:"not null" json:"district"`           // district of the contact for address
	AddressNumber  string `gorm:"not null" json:"address_number"`     // address number of the contact for address
	City           string `gorm:"not null" json:"city"`               // city of the contact for address
	State          string `gorm:"not null" json:"state"`              // state of the contact for address
	Country        string `gorm:"not null" json:"country"`            // country of the contact for address
	Email          string `json:"email"`                              // email address of the contact
	CustomerID     uint   `json:"customer_id"`                        // Foreign key for Customer (one-to-one)
	SupplierID     uint   `json:"supplier_id"`                        // Foreign key for Supplier (one-to-one)
}

// TableName overrides the table name used by Contact to `sales.contacts`.
func (Contact) TableName() string {
	return "sales.contacts"
}
