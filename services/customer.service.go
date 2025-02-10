package services

import (
	"store/domain/entities"
	"store/domain/repositories"

	"github.com/gin-gonic/gin"
)

// CustomerService is an interface that defines methods for interacting with customers.
//
// It provides methods to create, retrieve, update, and delete customers in the database.
//
// The interface is used to define the methods that a customer service should provide.
type CustomerService interface {
	Create(ctx *gin.Context, customer *entities.Customer) error
	GetByID(ctx *gin.Context, id uint) (*entities.Customer, error)
	GetAll(ctx *gin.Context) ([]*entities.Customer, error)
	Update(ctx *gin.Context, customer *entities.Customer) error
	Delete(ctx *gin.Context, id uint) error
	DeleteAll(ctx *gin.Context, ids []uint) error
}

// customerService is a struct that contains a pointer to a customerRepository and
// implements the CustomerService. The struct contains a pointer to a
// customerRepository which is used to interact with the customers table in the
// database.
type customerService struct {
	customerRepository repositories.CustomerRepository
}

// NewCustomerService creates a new instance of customerService with the provided
// customerRepository and returns it as a CustomerService. This function is used to
// initialize a new customer service that can perform CRUD operations and other
// queries on the customers table.
func NewCustomerService(customerRepository repositories.CustomerRepository) CustomerService {
	return &customerService{customerRepository: customerRepository}
}

// Creates a new customer in the database.
//
// The method takes a pointer to a *gin.Context and a pointer to a entities.Customer
// as parameters. It returns an error if something goes wrong.
//
// The method creates a new customer in the database using the given customer object.
// The customer object is passed as a pointer and the method is responsible for creating
// a new customer in the database with the given attributes.
//
// The method returns an error if something goes wrong. If the customer is created
// successfully, the method returns nil.
func (s *customerService) Create(ctx *gin.Context, customer *entities.Customer) error {
	return s.customerRepository.Create(ctx, customer)
}

// Retrieves a customer by its ID from the database.
//
// The method takes a pointer to a *gin.Context and a uint as parameters. It
// returns a pointer to a entities.Customer and an error. If something goes wrong,
// the method returns nil and an error.
//
// The method gets a customer by its ID from the database using the given ID.
// The method returns a pointer to a entities.Customer and an error. If the customer
// is found, the method returns the customer and nil. If the customer is not found,
// the method returns nil and an error.
func (s *customerService) GetByID(ctx *gin.Context, id uint) (*entities.Customer, error) {
	return s.customerRepository.GetByID(ctx, id)
}

// Retrieves all customers from the database.
//
// The method takes a pointer to a *gin.Context as a parameter. It returns a slice
// of pointers to entities.Customer and an error. If something goes wrong, the
// method returns an empty slice and an error.
//
// The method retrieves all customers from the database. If the customers are found,
// the method returns the customers and nil. If no customers are found or an error
// occurs, the method returns an empty slice and an error.
func (s *customerService) GetAll(ctx *gin.Context) ([]*entities.Customer, error) {
	return s.customerRepository.GetAll(ctx)
}

// Updates a customer in the database.
//
// The method takes a pointer to a *gin.Context and a pointer to a entities.Customer
// as parameters. It returns an error if something goes wrong.
//
// The method updates a customer in the database using the given customer object.
// The customer object is passed as a pointer and the method is responsible for updating
// a customer in the database with the given attributes.
//
// The method returns an error if something goes wrong. If the customer is updated
// successfully, the method returns nil.
func (s *customerService) Update(ctx *gin.Context, customer *entities.Customer) error {
	return s.customerRepository.Update(ctx, customer)
}

// Deletes a customer from the database.
//
// The method takes a pointer to a *gin.Context and a uint as parameters. It
// returns an error if something goes wrong.
//
// The method deletes a customer from the database using the given ID.
// If the customer is deleted successfully, the method returns nil. If the customer
// is not found or an error occurs, the method returns an error.
func (s *customerService) Delete(ctx *gin.Context, id uint) error {
	return s.customerRepository.Delete(ctx, id)
}

// Deletes multiple customers from the database by their IDs.
//
// The method takes a pointer to a *gin.Context and a slice of uints as parameters.
// It returns an error if something goes wrong.
//
// The method deletes multiple customers from the database using the given slice
// of IDs. The method returns an error if something goes wrong. If the customers are
// deleted successfully, the method returns nil.
func (s *customerService) DeleteAll(ctx *gin.Context, ids []uint) error {
	return s.customerRepository.DeleteAll(ctx, ids)
}
