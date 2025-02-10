package services

import (
	"store/domain/entities"
	"store/domain/repositories"

	"github.com/gin-gonic/gin"
)

// OrderProductSupplierService is an interface that defines the methods that must
// be implemented by any service that wants to interact with the order_product_suppliers
// table in the database.
//
// It provides methods for creating a new orderProductSupplier, getting a orderProductSupplier by its ID,
// getting all orderProductSuppliers, updating a orderProductSupplier, deleting a orderProductSupplier,
// and deleting multiple orderProductSuppliers.
type OrderProductSupplierService interface {
	Create(ctx *gin.Context, orderProductSupplier *entities.OrderProductSupplier) error // Create a new orderProductSupplier
	GetByID(ctx *gin.Context, id uint) (*entities.OrderProductSupplier, error)          // Get a orderProductSupplier by ID
	GetAll(ctx *gin.Context) ([]*entities.OrderProductSupplier, error)                  // Get all orderProductSuppliers
	Update(ctx *gin.Context, orderProductSupplier *entities.OrderProductSupplier) error // Update a orderProductSupplier
	Delete(ctx *gin.Context, id uint) error                                             // Delete a orderProductSupplier
	DeleteAll(ctx *gin.Context, ids []uint) error                                       // Delete a orderProductSupplier
}

// orderProductSupplierService is a struct that implements the OrderProductSupplierService interface.
// It contains a pointer to a OrderProductSupplierRepository and provides methods for creating a new orderProductSupplier,
// getting a orderProductSupplier by its ID, getting all orderProductSuppliers, updating a orderProductSupplier,
// deleting a orderProductSupplier, and deleting multiple orderProductSuppliers.
type orderProductSupplierService struct {
	orderProductSupplierRepository repositories.OrderProductSupplierRepository
}

// NewOrderProductSupplierService creates a new OrderProductSupplierService with the given
// OrderProductSupplierRepository. The OrderProductSupplierService is an interface that defines
// methods for creating, retrieving, updating, and deleting orderProductSuppliers in the
// application. It returns an instance of orderProductSupplierService that implements the
// OrderProductSupplierService interface.
func NewOrderProductSupplierService(orderProductSupplierRepository repositories.OrderProductSupplierRepository) OrderProductSupplierService {
	return &orderProductSupplierService{orderProductSupplierRepository: orderProductSupplierRepository}
}

// Creates a new orderProductSupplier to the database.
//
// The method takes a context and an orderProductSupplier entity as parameters.
// It delegates the creation of the orderProductSupplier to the repository and
// returns an error if the creation process fails. If successful, it returns nil.
func (s *orderProductSupplierService) Create(ctx *gin.Context, orderProductSupplier *entities.OrderProductSupplier) error {
	return s.orderProductSupplierRepository.Create(ctx, orderProductSupplier)
}

// Retrieves an orderProductSupplier by its ID from the database.
//
// The method takes a context and the ID of the orderProductSupplier as parameters.
// It delegates the retrieval of the orderProductSupplier to the repository and
// returns an error if the retrieval process fails. If successful, it returns
// the orderProductSupplier and nil.
func (s *orderProductSupplierService) GetByID(ctx *gin.Context, id uint) (*entities.OrderProductSupplier, error) {
	return s.orderProductSupplierRepository.GetByID(ctx, id)
}

// Retrieves all orderProductSuppliers from the database.
//
// The method takes a context as a parameter. It delegates the retrieval of the
// orderProductSuppliers to the repository and returns an error if the retrieval
// process fails. If successful, it returns a slice of orderProductSuppliers and
// nil.
func (s *orderProductSupplierService) GetAll(ctx *gin.Context) ([]*entities.OrderProductSupplier, error) {
	return s.orderProductSupplierRepository.GetAll(ctx)
}

// Updates an orderProductSupplier in the database.
//
// The method takes a context and an orderProductSupplier entity as parameters.
// It delegates the update of the orderProductSupplier to the repository and
// returns an error if the update process fails. If successful, it returns nil.
func (s *orderProductSupplierService) Update(ctx *gin.Context, orderProductSupplier *entities.OrderProductSupplier) error {
	return s.orderProductSupplierRepository.Update(ctx, orderProductSupplier)
}

// Deletes an orderProductSupplier by its ID from the database.
//
// The method takes a context and the ID of the orderProductSupplier as parameters.
// It delegates the deletion of the orderProductSupplier to the repository and
// returns an error if the deletion process fails. If successful, it returns nil.
func (s *orderProductSupplierService) Delete(ctx *gin.Context, id uint) error {
	return s.orderProductSupplierRepository.Delete(ctx, id)
}

// Deletes multiple orderProductSuppliers by their IDs from the database.
//
// The method takes a context and a slice of uints as parameters.
// It delegates the deletion of the orderProductSuppliers to the repository and
// returns an error if the deletion process fails. If successful, it returns nil.
func (s *orderProductSupplierService) DeleteAll(ctx *gin.Context, ids []uint) error {
	return s.orderProductSupplierRepository.DeleteAll(ctx, ids)
}
