package services

import (
	"store/domain/entities"
	"store/domain/repositories"

	"github.com/gin-gonic/gin"
)

// SupplierService defines the methods that a service must implement to manage
// suppliers in the application. It provides methods to create, retrieve, update,
// and delete supplier entities.
type SupplierService interface {
	Create(ctx *gin.Context, supplier *entities.Supplier) error    // Create a new supplier
	GetByID(ctx *gin.Context, id uint) (*entities.Supplier, error) // Get a supplier by ID
	GetAll(ctx *gin.Context) ([]*entities.Supplier, error)         // Get all suppliers
	Update(ctx *gin.Context, supplier *entities.Supplier) error    // Update a supplier
	Delete(ctx *gin.Context, id uint) error                        // Delete a supplier
	DeleteAll(ctx *gin.Context, ids []uint) error                  // Delete multiple suppliers
}

// supplierService is a struct that implements the SupplierService interface.
// It contains a pointer to a SupplierRepository which is used to interact
// with the suppliers table in the database.
type supplierService struct {
	supplierRepository repositories.SupplierRepository
}

// NewSupplierService creates a new SupplierService with the given supplierRepository.
// The SupplierService is an interface that defines methods for creating, retrieving,
// updating, and deleting suppliers in the application.
func NewSupplierService(supplierRepository repositories.SupplierRepository) SupplierService {
	return &supplierService{supplierRepository: supplierRepository}
}

// Create a new supplier in the database.
//
// The method takes a pointer to a *gin.Context and a pointer to a entities.Supplier
// as parameters. It returns an error if something goes wrong.
//
// The method creates a new supplier in the database using the given supplier object.
// The supplier object is passed as a pointer and the method is responsible for creating
// a new supplier in the database with the given attributes.
//
// The method returns an error if something goes wrong. If the supplier is created
// successfully, the method returns nil.
func (s *supplierService) Create(ctx *gin.Context, supplier *entities.Supplier) error {
	return s.supplierRepository.Create(ctx, supplier)
}

// Retrieves a supplier by its ID from the database.
//
// The method takes a pointer to a *gin.Context and a uint as parameters.
// It returns a pointer to an entities.Supplier and an error. If something
// goes wrong, the method returns nil and an error.
//
// The method retrieves a supplier by its ID from the database using the
// given ID. If the supplier is found, the method returns the supplier
// and nil. If the supplier is not found, the method returns nil and an
// error.
func (s *supplierService) GetByID(ctx *gin.Context, id uint) (*entities.Supplier, error) {
	return s.supplierRepository.GetByID(ctx, id)
}

// Retrieves all suppliers from the database.
//
// The method takes a pointer to a *gin.Context as a parameter.
// It returns a slice of pointers to entities.Supplier and an error.
// If something goes wrong, the method returns an empty slice and an error.
//
// The method retrieves all suppliers from the database using the supplierRepository.
// If suppliers are found, it returns the suppliers and nil. If no suppliers
// are found or an error occurs, it returns an empty slice and an error.
func (s *supplierService) GetAll(ctx *gin.Context) ([]*entities.Supplier, error) {
	return s.supplierRepository.GetAll(ctx)
}

// Updates a supplier in the database.
//
// The method takes a pointer to a *gin.Context and a pointer to an entities.Supplier
// as parameters. It returns an error if something goes wrong.
//
// The method updates a supplier in the database using the given supplier object.
// If the supplier is updated successfully, the method returns nil. If an error
// occurs, the method returns an error.
func (s *supplierService) Update(ctx *gin.Context, supplier *entities.Supplier) error {
	return s.supplierRepository.Update(ctx, supplier)
}

// Deletes a supplier by its ID from the database.
//
// The method takes a pointer to a *gin.Context and a uint as parameters.
// It returns an error if something goes wrong.
//
// The method deletes a supplier by its ID from the database using the
// given ID. If the supplier is successfully deleted, the method returns
// nil. If the supplier is not found or an error occurs, the method returns
// an error.
func (s *supplierService) Delete(ctx *gin.Context, id uint) error {
	return s.supplierRepository.Delete(ctx, id)
}

// Deletes multiple suppliers by their IDs from the database.
//
// The method takes a pointer to a *gin.Context and a slice of uints as parameters.
// It returns an error if something goes wrong.
//
// The method deletes multiple suppliers from the database using the given slice
// of IDs. If the suppliers are deleted successfully, the method returns nil.
// If an error occurs, the method returns an error.
func (s *supplierService) DeleteAll(ctx *gin.Context, ids []uint) error {
	return s.supplierRepository.DeleteAll(ctx, ids)
}
