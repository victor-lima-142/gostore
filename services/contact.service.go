package services

import (
	"store/domain/entities"
	"store/domain/repositories"

	"github.com/gin-gonic/gin"
)

// ContactService is an interface that defines methods for interacting with contacts.
//
// It provides methods to create, retrieve, update, and delete contacts in the database.
//
// The interface is used to define the methods that a contact service should provide.
type ContactService interface {
	Create(ctx *gin.Context, contact *entities.Contact) error
	GetByID(ctx *gin.Context, id uint) (*entities.Contact, error)
	GetAll(ctx *gin.Context) ([]*entities.Contact, error)
	Update(ctx *gin.Context, contact *entities.Contact) error
	Delete(ctx *gin.Context, id uint) error
	DeleteAll(ctx *gin.Context, ids []uint) error
	GetAllByCustomerID(ctx *gin.Context, customerID uint) ([]*entities.Contact, error)
	GetAllBySupplierID(ctx *gin.Context, supplierID uint) ([]*entities.Contact, error)
}

// contactService is a struct that contains a pointer to a contactRepository and
// implements the ContactService. The struct contains a pointer to a
// contactRepository which is used to interact with the contacts table in the
// database.
type contactService struct {
	contactRepository repositories.ContactRepository
}

// NewContactService creates a new instance of contactService with the provided
// contactRepository and returns it as a ContactService. This function is used to
// initialize a new contact service that can perform CRUD operations and other
// queries on the contacts table.
func NewContactService(contactRepository repositories.ContactRepository) ContactService {
	return &contactService{contactRepository: contactRepository}
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
func (s *contactService) Create(ctx *gin.Context, contact *entities.Contact) error {
	return s.contactRepository.Create(ctx, contact)
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
func (s *contactService) GetByID(ctx *gin.Context, id uint) (*entities.Contact, error) {
	return s.contactRepository.GetByID(ctx, id)
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
func (s *contactService) GetAll(ctx *gin.Context) ([]*entities.Contact, error) {
	return s.contactRepository.GetAll(ctx)
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
func (s *contactService) Update(ctx *gin.Context, contact *entities.Contact) error {
	return s.contactRepository.Update(ctx, contact)
}

// Deletes a contact from the database.
//
// The method takes a pointer to a *gin.Context and a uint as parameters. It
// returns an error if something goes wrong.
//
// The method deletes a contact from the database using the given ID.
// If the contact is deleted successfully, the method returns nil. If the contact
// is not found or an error occurs, the method returns an error.
func (s *contactService) Delete(ctx *gin.Context, id uint) error {
	return s.contactRepository.Delete(ctx, id)
}

// Deletes multiple contacts from the database by their IDs.
//
// The method takes a pointer to a *gin.Context and a slice of uints as parameters.
// It returns an error if something goes wrong.
//
// The method deletes multiple contacts from the database using the given slice
// of IDs. The method returns an error if something goes wrong. If the contacts are
// deleted successfully, the method returns nil.
func (s *contactService) DeleteAll(ctx *gin.Context, ids []uint) error {
	return s.contactRepository.DeleteAll(ctx, ids)
}

// GetAllByCustomerID gets all contacts from the database that belong to the
// given customer ID.
//
// The method takes a pointer to a *gin.Context and a uint as parameters. It
// returns a slice of pointers to entities.Contact and an error. If something goes
// wrong, the method returns an empty slice and an error.
//
// The method returns a slice of pointers to entities.Contact and an error. If the
// contacts are found, the method returns the contacts and nil. If no contacts are
// found or an error occurs, the method returns an empty slice and an error.
func (s *contactService) GetAllByCustomerID(ctx *gin.Context, customerID uint) ([]*entities.Contact, error) {
	return s.contactRepository.GetAllByCustomerID(ctx, customerID)
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
func (s *contactService) GetAllBySupplierID(ctx *gin.Context, supplierID uint) ([]*entities.Contact, error) {
	return s.contactRepository.GetAllBySupplierID(ctx, supplierID)
}
