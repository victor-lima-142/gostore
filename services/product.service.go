package services

import (
	"store/domain/entities"
	"store/domain/repositories"

	"github.com/gin-gonic/gin"
)

// ProductService defines the methods that a service must implement to manage
// products in the application. It provides methods to create, retrieve, update,
// and delete product entities.
type ProductService interface {
	Create(ctx *gin.Context, supplier *entities.Product) error    // Create a new product
	GetByID(ctx *gin.Context, id uint) (*entities.Product, error) // Get a product by ID
	GetAll(ctx *gin.Context) ([]*entities.Product, error)         // Get all products
	Update(ctx *gin.Context, supplier *entities.Product) error    // Update a product
	Delete(ctx *gin.Context, id uint) error                       // Delete a product
	DeleteAll(ctx *gin.Context, ids []uint) error                 // Delete a product
}

// productService is a struct that contains a pointer to a repositories.ProductRepository
// and implements the ProductService interface.
// productRepository which is used to interact with the customers table in the
// database.
type productService struct {
	productRepository repositories.ProductRepository
}

// NewProductService creates a new ProductService with the given productRepository.
// The ProductService is an interface that defines methods for creating, retrieving,
// updating, and deleting products in the application.
func NewProductService(productRepository repositories.ProductRepository) ProductService {
	return &productService{productRepository: productRepository}
}

// Creates a new product to the database.
//
// The method takes a pointer to a *gin.Context for handling request-specific data
// and a pointer to an entities.Product representing the product to be added.
// It returns an error if the creation process encounters any issues.
//
// This method ensures that the product is created in the database with the provided
// attributes. If successful, it returns nil; otherwise, it returns the encountered error.
func (s *productService) Create(ctx *gin.Context, product *entities.Product) error {
	return s.productRepository.Create(ctx, product)
}

// Retrieves a product from the database by its ID.
//
// The method takes a pointer to a *gin.Context for handling request-specific data
// and a uint representing the ID of the product to be retrieved. It returns a
// pointer to an entities.Product if the product is found, and an error if the
// product is not found or if something goes wrong during the retrieval process.
//
// This method ensures that the product is retrieved from the database with the
// provided ID. If the product is found, it returns the product and nil; otherwise,
// it returns nil and an error.
func (s *productService) GetByID(ctx *gin.Context, id uint) (*entities.Product, error) {
	return s.productRepository.GetByID(ctx, id)
}

// Retrieves all products from the database.
//
// The method takes a pointer to a *gin.Context as a parameter. It returns a slice
// of pointers to entities.Product and an error. If something goes wrong, the
// method returns an empty slice and an error.
//
// The method fetches all products from the database. If the products are found,
// the method returns the products and nil. If no products are found or an error
// occurs, the method returns an empty slice and an error.
func (s *productService) GetAll(ctx *gin.Context) ([]*entities.Product, error) {
	return s.productRepository.GetAll(ctx)
}

// Updates an existing product in the database.
//
// The method takes a pointer to a *gin.Context for handling request-specific data
// and a pointer to an entities.Product representing the product to be updated.
// It returns an error if the update process encounters any issues.
//
// This method ensures that the product is updated in the database with the provided
// attributes. If successful, it returns nil; otherwise, it returns the encountered error.
func (s *productService) Update(ctx *gin.Context, product *entities.Product) error {
	return s.productRepository.Update(ctx, product)
}

// Deletes a product from the database by its ID.
//
// The method takes a pointer to a *gin.Context for handling request-specific data
// and a uint representing the ID of the product to be deleted. It returns an error
// if the deletion process encounters any issues.
//
// The method deletes a product from the database using the given ID. If the
// product is deleted successfully, the method returns nil. If the product is not
// found or an error occurs, the method returns an error.
func (s *productService) Delete(ctx *gin.Context, id uint) error {
	return s.productRepository.Delete(ctx, id)
}

// Deletes multiple products from the database by their IDs.
//
// The method takes a pointer to a *gin.Context and a slice of uints as parameters.
// It returns an error if something goes wrong.
//
// The method deletes multiple products from the database using the given slice
// of IDs. The method returns an error if something goes wrong. If the products are
// deleted successfully, the method returns nil.
func (s *productService) DeleteAll(ctx *gin.Context, ids []uint) error {
	return s.productRepository.DeleteAll(ctx, ids)
}
