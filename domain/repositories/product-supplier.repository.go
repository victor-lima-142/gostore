package repositories

import (
	"store/domain/entities"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// ProductSupplierRepository is an interface that defines the methods that must
// be implemented by any data store that wants to interact with the product_suppliers
// table in the database.
//
// It provides methods for creating a new productSupplier, getting a productSupplier by its ID, getting all productSuppliers,
// updating a productSupplier, and deleting a productSupplier.
type ProductSupplierRepository interface {
	Create(ctx *gin.Context, productSupplier *entities.ProductSupplier) error // Create a new productSupplier
	GetByID(ctx *gin.Context, id uint) (*entities.ProductSupplier, error)     // Get a productSupplier by ID
	GetAll(ctx *gin.Context) ([]*entities.ProductSupplier, error)             // Get all productSuppliers
	Update(ctx *gin.Context, productSupplier *entities.ProductSupplier) error // Update a productSupplier
	Delete(ctx *gin.Context, id uint) error                                   // Delete a productSupplier
	DeleteAll(ctx *gin.Context, ids []uint) error                             // Delete multiple productSuppliers
}

// productSupplierRepository is a struct that contains a pointer to a gorm DB instance and
// implements the ProductSupplierRepository.
// This function is used to initialize a new product repository that can perform
// CRUD operations and other queries on the products table.
type productSupplierRepository struct {
	db *gorm.DB
}

// NewProductSupplierRepository creates a new instance of productSupplierRepository with the provided
// database instance and returns it as a ProductSupplierRepository.
// This function is used to initialize a new productSupplier repository that can perform
// CRUD operations and other queries on the product_suppliers table.
func NewProductSupplierRepository(db *gorm.DB) ProductSupplierRepository {
	return &productSupplierRepository{db: db}
}

// Creates a new productSupplier in the database.
//
// The method takes a pointer to a *gin.Context and a pointer to a entities.ProductSupplier
// as parameters. It returns an error if something goes wrong.
//
// The method creates a new productSupplier in the database using the given productSupplier object.
// The productSupplier object is passed as a pointer and the method is responsible for creating
// a new productSupplier in the database with the given attributes.
//
// The method returns an error if something goes wrong. If the productSupplier is created
// successfully, the method returns nil.
func (r *productSupplierRepository) Create(ctx *gin.Context, productSupplier *entities.ProductSupplier) error {
	return r.db.WithContext(ctx).Create(productSupplier).Error
}

// Retrieves a productSupplier by its ID from the database.
//
// The method takes a pointer to a *gin.Context and a uint as parameters. It
// returns a pointer to a entities.ProductSupplier and an error. If something goes wrong,
// the method returns nil and an error.
//
// The method gets a productSupplier by its ID from the database using the given ID.
// The method returns a pointer to a entities.ProductSupplier and an error. If the
// productSupplier is found, the method returns the productSupplier and nil. If the
// productSupplier is not found, the method returns nil and an error.
func (r *productSupplierRepository) GetByID(ctx *gin.Context, id uint) (*entities.ProductSupplier, error) {
	var productSupplier entities.ProductSupplier
	err := r.db.WithContext(ctx).First(&productSupplier, id).Error
	return &productSupplier, err
}

// Retrieves all productSuppliers from the database.
//
// The method takes a pointer to a *gin.Context as a parameter. It returns a slice
// of pointers to entities.ProductSupplier and an error. If something goes wrong,
// the method returns an empty slice and an error.
//
// The method fetches all productSuppliers from the database. If the productSuppliers
// are found, the method returns the productSuppliers and nil. If no productSuppliers
// are found or an error occurs, the method returns an empty slice and an error.
func (r *productSupplierRepository) GetAll(ctx *gin.Context) ([]*entities.ProductSupplier, error) {
	var productSuppliers []*entities.ProductSupplier
	err := r.db.WithContext(ctx).Find(&productSuppliers).Error
	return productSuppliers, err
}

// Updates a productSupplier in the database.
//
// The method takes a pointer to a *gin.Context and a pointer to a entities.ProductSupplier
// as parameters. It returns an error if something goes wrong.
//
// The method updates a productSupplier in the database using the given productSupplier object.
// The productSupplier object is passed as a pointer and the method is responsible for updating
// a productSupplier in the database with the given attributes.
//
// The method returns an error if something goes wrong. If the productSupplier is updated
// successfully, the method returns nil.
func (r *productSupplierRepository) Update(ctx *gin.Context, productSupplier *entities.ProductSupplier) error {
	return r.db.WithContext(ctx).Save(productSupplier).Error
}

// Deletes a productSupplier by its ID from the database.
//
// The method takes a pointer to a *gin.Context and a uint as parameters. It
// returns an error if something goes wrong.
//
// The method deletes a productSupplier from the database using the given ID.
// If the productSupplier is deleted successfully, the method returns nil. If the
// productSupplier is not found or an error occurs, the method returns an error.
func (r *productSupplierRepository) Delete(ctx *gin.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&entities.ProductSupplier{}, id).Error
}

// Deletes multiple productSuppliers from the database by their IDs.
//
// The method takes a pointer to a *gin.Context and a slice of uints as parameters.
// It returns an error if something goes wrong.
//
// The method deletes multiple productSuppliers from the database using the given slice
// of IDs. The method returns an error if something goes wrong. If the productSuppliers
// are deleted successfully, the method returns nil.
func (r *productSupplierRepository) DeleteAll(ctx *gin.Context, ids []uint) error {
	return r.db.WithContext(ctx).Delete(&entities.ProductSupplier{}, ids).Error
}
