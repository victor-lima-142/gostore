package services

import (
	"store/domain/entities"
	"store/domain/repositories"

	"github.com/gin-gonic/gin"
)

// ProductSupplierService is an interface that defines the methods that a service must
// implement to manage product suppliers in the application. It provides methods to
// create, retrieve, update, and delete productSupplier entities.
type ProductSupplierService interface {
	Create(ctx *gin.Context, productSupplier *entities.ProductSupplier) error // Creates a productSupplier
	GetByID(ctx *gin.Context, id uint) (*entities.ProductSupplier, error)     // Retrieves a productSupplier
	GetAll(ctx *gin.Context) ([]*entities.ProductSupplier, error)             // Retrieves all productSuppliers
	Update(ctx *gin.Context, productSupplier *entities.ProductSupplier) error // Updates a productSupplier
	Delete(ctx *gin.Context, id uint) error                                   // Deletes a productSupplier
	DeleteAll(ctx *gin.Context, ids []uint) error                             // Deletes multiple productSuppliers
}

// productSupplierService is a struct that implements the ProductSupplierService interface.
// It contains a pointer to a ProductSupplierRepository which is used to interact
// with the product_suppliers table in the database.
type productSupplierService struct {
	productSupplierRepository repositories.ProductSupplierRepository
}

// NewProductSupplierService creates a new ProductSupplierService with the given productSupplierRepository.
// The ProductSupplierService is an interface that defines methods for creating, retrieving,
// updating, and deleting productSupplier entities in the application.
// It returns an instance of productSupplierService that implements the ProductSupplierService interface.
func NewProductSupplierService(productSupplierRepository repositories.ProductSupplierRepository) ProductSupplierService {
	return &productSupplierService{productSupplierRepository: productSupplierRepository}
}

// Creates a new productSupplier in the database.
//
// The method takes a context and an entities.ProductSupplier as parameters.
// It delegates the creation of the productSupplier to the productSupplierRepository and
// returns an error if the creation process fails. If successful, it returns nil.
func (s *productSupplierService) Create(ctx *gin.Context, productSupplier *entities.ProductSupplier) error {
	return s.productSupplierRepository.Create(ctx, productSupplier)
}

// Retrieves a productSupplier by its ID from the database.
//
// The method takes a context and the ID of the productSupplier as parameters.
// It delegates the retrieval of the productSupplier to the productSupplierRepository and
// returns an error if the retrieval process fails. If successful, it returns
// the productSupplier and nil.
func (s *productSupplierService) GetByID(ctx *gin.Context, id uint) (*entities.ProductSupplier, error) {
	return s.productSupplierRepository.GetByID(ctx, id)
}

// Retrieves all productSuppliers from the database.
//
// The method takes a context as a parameter. It delegates the retrieval of the
// productSuppliers to the productSupplierRepository and returns an error if the retrieval
// process fails. If successful, it returns a slice of productSuppliers and
// nil.
func (s *productSupplierService) GetAll(ctx *gin.Context) ([]*entities.ProductSupplier, error) {
	return s.productSupplierRepository.GetAll(ctx)
}

// Updates a productSupplier in the database.
//
// The method takes a context and an entities.ProductSupplier as parameters.
// It delegates the update of the productSupplier to the productSupplierRepository and
// returns an error if the update process fails. If successful, it returns nil.
func (s *productSupplierService) Update(ctx *gin.Context, productSupplier *entities.ProductSupplier) error {
	return s.productSupplierRepository.Update(ctx, productSupplier)
}

// Deletes a productSupplier by its ID from the database.
//
// The method takes a context and the ID of the productSupplier as parameters.
// It delegates the deletion of the productSupplier to the productSupplierRepository and
// returns an error if the deletion process fails. If successful, it returns nil.
func (s *productSupplierService) Delete(ctx *gin.Context, id uint) error {
	return s.productSupplierRepository.Delete(ctx, id)
}

// Deletes multiple productSuppliers by their IDs from the database.
//
// The method takes a context and a slice of uints as parameters.
// It delegates the deletion of the productSuppliers to the productSupplierRepository and
// returns an error if the deletion process fails. If successful, it returns nil.

func (s *productSupplierService) DeleteAll(ctx *gin.Context, ids []uint) error {
	return s.productSupplierRepository.DeleteAll(ctx, ids)
}
