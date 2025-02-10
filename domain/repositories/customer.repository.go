package repositories

import (
	"store/domain/entities"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// CustomerRepository is an interface that defines the methods that must
// be implemented by any data store that wants to interact with the customers
// table in the database.
//
// It provides methods for creating a new customer, getting a customer by its ID,
// getting all customers, updating a customer, deleting a customer, and getting a
// customer with its orders or contact.
type CustomerRepository interface {
	Create(ctx *gin.Context, customer *entities.Customer) error                   // Create a new customer
	GetByID(ctx *gin.Context, id uint) (*entities.Customer, error)                // Get a customer by ID
	GetAll(ctx *gin.Context) ([]*entities.Customer, error)                        // Get all customers
	Update(ctx *gin.Context, customer *entities.Customer) error                   // Update a customer
	Delete(ctx *gin.Context, id uint) error                                       // Delete a customer
	DeleteAll(ctx *gin.Context, ids []uint) error                                 // Delete multiple customers	GetCustomerWithOrders(ctx *gin.Context, id uint) (*entities.Customer, error)  // Get a customer with orders
	GetCustomerWithContact(ctx *gin.Context, id uint) (*entities.Customer, error) // Get a customer with contact
}

// customerRepository is a struct that contains a pointer to a gorm DB instance and
// implements the CustomerRepository.
//
// The struct contains a pointer to a gorm DB instance which is used to interact
// with the customers table in the database.
type customerRepository struct {
	db *gorm.DB
}

// NewCustomerRepository creates a new instance of contactRepository with the provided
// database instance and returns it as a CustomerRepository.
// This function is used to initialize a new product repository that can perform
// CRUD operations and other queries on the products table.
func NewCustomerRepository(db *gorm.DB) CustomerRepository {
	return &customerRepository{db: db}
}

// Creates a new customer in the database.
//
// This method is a part of the CustomerRepository and is used to create a
// new customer in the database.
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
func (r *customerRepository) Create(ctx *gin.Context, customer *entities.Customer) error {
	return r.db.WithContext(ctx).Create(customer).Error
}

// Retrieves a customer by its ID.
//
// This method is a part of the CustomerRepository and is used to get a
// customer by its ID.
//
// The method takes a pointer to a *gin.Context and a uint as parameters. It
// returns a pointer to a entities.Customer and an error. If something goes wrong,
// the method returns nil and an error.
//
// The method gets a customer by its ID from the database using the given ID.
// The method returns a pointer to a entities.Customer and an error. If the
// customer is found, the method returns the customer and nil. If the customer is
// not found, the method returns nil and an error.
func (r *customerRepository) GetByID(ctx *gin.Context, id uint) (*entities.Customer, error) {
	var customer entities.Customer
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&customer).Error
	return &customer, err
}

// Retrieves all customers from the database.
//
// This method is a part of the CustomerRepository and is used to get all
// customers from the database.
//
// The method takes a pointer to a *gin.Context as a parameter. It returns a
// slice of pointers to entities.Customer and an error. If something goes wrong,
// the method returns an empty slice and an error.
//
// The method gets all customers from the database. The method returns a slice of
// pointers to entities.Customer and an error. If the customers are found, the
// method returns the customers and nil. If no customers are found, the method
// returns an empty slice and an error.
func (r *customerRepository) GetAll(ctx *gin.Context) ([]*entities.Customer, error) {
	var customers []*entities.Customer
	err := r.db.WithContext(ctx).Find(&customers).Error
	return customers, err
}

// Updates a customer in the database.
//
// This method is a part of the CustomerRepository and is used to update a
// customer in the database.
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
func (r *customerRepository) Update(ctx *gin.Context, customer *entities.Customer) error {
	return r.db.WithContext(ctx).Save(customer).Error
}

// Deletes a customer from the database.
//
// This method is a part of the CustomerRepository and is used to delete a
// customer from the database.
//
// The method takes a pointer to a *gin.Context and a uint as parameters. It
// returns an error if something goes wrong.
//
// The method deletes a customer from the database using the given ID.
// The method returns an error if something goes wrong. If the customer is deleted
// successfully, the method returns nil.
func (r *customerRepository) Delete(ctx *gin.Context, id uint) error {
	return r.db.WithContext(ctx).Where("id = ?", id).Delete(&entities.Customer{}).Error
}

// Deletes multiple customers from the database by their IDs.
//
// The method takes a pointer to a *gin.Context and a slice of uints as parameters.
// It returns an error if something goes wrong.
//
// The method deletes multiple customers from the database using the given slice
// of IDs. The method returns an error if something goes wrong. If the customers
// are deleted successfully, the method returns nil.
func (r *customerRepository) DeleteAll(ctx *gin.Context, ids []uint) error {
	return r.db.WithContext(ctx).Where("id in (?)", ids).Delete(&entities.Customer{}).Error
}

// Retrieves a customer by its ID with its orders.
//
// The method takes a pointer to a *gin.Context and a uint as parameters. It
// returns a pointer to a entities.Customer and an error. If something goes wrong,
// the method returns nil and an error.
//
// The method gets a customer by its ID from the database using the given ID, and
// preloads the Orders field. The method returns a pointer to a entities.Customer
// and an error. If the customer is found, the method returns the customer and nil.
// If the customer is not found, the method returns nil and an error.
func (r *customerRepository) GetCustomerWithOrders(ctx *gin.Context, id uint) (*entities.Customer, error) {
	var customer entities.Customer
	err := r.db.WithContext(ctx).
		Preload("Orders").
		Where("id = ?", id).
		First(&customer).
		Error
	return &customer, err
}

// Retrieves a customer by its ID with its contact information.
//
// The method takes a pointer to a *gin.Context and a uint as parameters. It
// returns a pointer to a entities.Customer and an error. If something goes wrong,
// the method returns nil and an error.
//
// The method gets a customer by its ID from the database using the given ID, and
// preloads the Contact field. The method returns a pointer to a entities.Customer
// and an error. If the customer is found, the method returns the customer and nil.
// If the customer is not found, the method returns nil and an error.
func (r *customerRepository) GetCustomerWithContact(ctx *gin.Context, id uint) (*entities.Customer, error) {
	var customer entities.Customer
	err := r.db.WithContext(ctx).
		Preload("Contact").
		Where("id = ?", id).
		First(&customer).
		Error
	return &customer, err
}
