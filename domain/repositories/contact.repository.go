package repositories

import (
	"store/domain/entities"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// ContactRepository is an interface that defines the methods that must
// be implemented by any data store that wants to interact with the contacts
// table in the database.
//
// It provides methods for creating a new contact, getting a contact by its ID,
// getting all customers, updating a contact, deleting a contact, and getting a
// contact with its orders or contact.
type ContactRepository interface {
	Create(ctx *gin.Context, contact *entities.Contact) error                          // Create a new contact
	GetByID(ctx *gin.Context, id uint) (*entities.Contact, error)                      // Get a contact by ID
	GetAll(ctx *gin.Context) ([]*entities.Contact, error)                              // Get all contacts
	Update(ctx *gin.Context, contact *entities.Contact) error                          // Update a contact
	Delete(ctx *gin.Context, id uint) error                                            // Delete a contact
	DeleteAll(ctx *gin.Context, ids []uint) error                                      // Delete multiple contacts
	GetAllByCustomerID(ctx *gin.Context, customerID uint) ([]*entities.Contact, error) // Get all contacts by customer ID
	GetAllBySupplierID(ctx *gin.Context, supplierID uint) ([]*entities.Contact, error) // Get all contacts by supplier ID
}

// contactRepository is a struct that contains a pointer to a gorm DB instance and
// implements the ContactRepository.
//
// The struct contains a pointer to a gorm DB instance which is used to interact
// with the contacts table in the database.
type contactRepository struct {
	db *gorm.DB
}

// NewContactRepository creates a new instance of contactRepository with the provided
// database instance and returns it as a ContactRepository.
// This function is used to initialize a new contact repository that can perform
// CRUD operations and other queries on the contacts table.
func NewContactRepository(db *gorm.DB) ContactRepository {
	return &contactRepository{db: db}
}

// Creates a new contact in the database.
//
// The method takes a pointer to a *gin.Context and a pointer to a entities.Contact
// as parameters. It returns an error if something goes wrong.
//
// The method creates a new contact in the database using the given contact object.
// The contact object is passed as a pointer and the method is responsible for creating
// a new contact in the database with the given attributes.
//
// The method returns an error if something goes wrong. If the contact is created
// successfully, the method returns nil.
func (r *contactRepository) Create(ctx *gin.Context, contact *entities.Contact) error {
	return r.db.WithContext(ctx).Create(contact).Error
}

// Retrieves a contact by its ID from the database.
//
// The method takes a pointer to a *gin.Context and a uint as parameters. It
// returns a pointer to a entities.Contact and an error. If something goes wrong,
// the method returns nil and an error.
//
// The method gets a contact by its ID from the database using the given ID.
// The method returns a pointer to a entities.Contact and an error. If the contact
// is found, the method returns the contact and nil. If the contact is not found,
// the method returns nil and an error.
func (r *contactRepository) GetByID(ctx *gin.Context, id uint) (*entities.Contact, error) {
	var contact entities.Contact
	err := r.db.WithContext(ctx).First(&contact, id).Error
	return &contact, err
}

// Retrieves all contacts from the database.
//
// The method takes a pointer to a *gin.Context as a parameter. It returns a slice
// of pointers to entities.Contact and an error. If something goes wrong, the
// method returns an empty slice and an error.
//
// The method retrieves all contacts from the database. If the contacts are found,
// the method returns the contacts and nil. If no contacts are found or an error
// occurs, the method returns an empty slice and an error.
func (r *contactRepository) GetAll(ctx *gin.Context) ([]*entities.Contact, error) {
	var contacts []*entities.Contact
	err := r.db.WithContext(ctx).Find(&contacts).Error
	return contacts, err
}

// Updates a contact in the database.
//
// The method takes a pointer to a *gin.Context and a pointer to a entities.Contact
// as parameters. It returns an error if something goes wrong.
//
// The method updates a contact in the database using the given contact object.
// The contact object is passed as a pointer and the method is responsible for updating
// a contact in the database with the given attributes.
//
// The method returns an error if something goes wrong. If the contact is updated
// successfully, the method returns nil.
func (r *contactRepository) Update(ctx *gin.Context, contact *entities.Contact) error {
	return r.db.WithContext(ctx).Save(contact).Error
}

// Deletes a contact from the database.
//
// The method takes a pointer to a *gin.Context and a uint as parameters. It
// returns an error if something goes wrong.
//
// The method deletes a contact from the database using the given ID.
// If the contact is deleted successfully, the method returns nil. If the contact
// is not found or an error occurs, the method returns an error.
func (r *contactRepository) Delete(ctx *gin.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&entities.Contact{}, id).Error
}

// Deletes multiple contacts from the database by their IDs.
//
// The method takes a pointer to a *gin.Context and a slice of uints as parameters.
// It returns an error if something goes wrong.
//
// The method deletes multiple contacts from the database using the given slice
// of IDs. The method returns an error if something goes wrong. If the contacts are
// deleted successfully, the method returns nil.
func (r *contactRepository) DeleteAll(ctx *gin.Context, ids []uint) error {
	return r.db.WithContext(ctx).Delete(&entities.Contact{}, ids).Error
}

// Retrieves all contacts from the database that belong to the
// given customer ID.
//
// The method takes a pointer to a *gin.Context and a uint as parameters. It
// returns a slice of pointers to entities.Contact and an error. If something goes
// wrong, the method returns an empty slice and an error.
//
// The method returns a slice of pointers to entities.Contact and an error. If the
// contacts are found, the method returns the contacts and nil. If no contacts are
// found or an error occurs, the method returns an empty slice and an error.
func (r *contactRepository) GetAllByCustomerID(ctx *gin.Context, customerID uint) ([]*entities.Contact, error) {
	var contacts []*entities.Contact
	err := r.db.WithContext(ctx).Where("customer_id = ?", customerID).Find(&contacts).Error
	return contacts, err
}

// Retrieves all contacts from the database that belong to the
// given supplier ID.
//
// The method takes a pointer to a *gin.Context and a uint as parameters. It
// returns a slice of pointers to entities.Contact and an error. If something goes
// wrong, the method returns an empty slice and an error.
//
// The method returns a slice of pointers to entities.Contact and an error. If the
// contacts are found, the method returns the contacts and nil. If no contacts are
// found or an error occurs, the method returns an empty slice and an error.
func (r *contactRepository) GetAllBySupplierID(ctx *gin.Context, supplierID uint) ([]*entities.Contact, error) {
	var contacts []*entities.Contact
	err := r.db.WithContext(ctx).Where("supplier_id = ?", supplierID).Find(&contacts).Error
	return contacts, err
}
