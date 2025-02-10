package repositories

import (
	"store/domain/entities"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// ProductRepository is an interface that defines the methods that must
// be implemented by any data store that wants to interact with the products
// table in the database.
//
// It provides methods for creating a new product, getting a product by its ID, getting all products,
// updating a product, and deleting a product.
type ProductRepository interface {
	Create(ctx *gin.Context, product *entities.Product) error     // Create a new product
	GetByID(ctx *gin.Context, id uint) (*entities.Product, error) // Get a product by its ID
	GetAll(ctx *gin.Context) ([]*entities.Product, error)         // Get all products
	Update(ctx *gin.Context, product *entities.Product) error     // Update a product
	Delete(ctx *gin.Context, id uint) error                       // Delete a product
	DeleteAll(ctx *gin.Context, ids []uint) error                 // Delete multiple products
}

// productRepository is a struct that contains a pointer to a gorm DB instance and
// implements the ProductRepository.
//
// The struct contains a pointer to a gorm DB instance which is used to interact
// with the products table in the database.
type productRepository struct {
	db *gorm.DB
}

// NewProductRepository creates a new instance of productRepository with the provided
// database instance and returns it as a ProductRepository.
// This function is used to initialize a new product repository that can perform
// CRUD operations and other queries on the products table.
func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{db: db}
}

// Creates a new product in the database.
//
// The method takes a pointer to a *gin.Context and a pointer to a entities.Product
// as parameters. It returns an error if something goes wrong.
//
// The method creates a new product in the database using the given product object.
// The product object is passed as a pointer and the method is responsible for creating
// a new product in the database with the given attributes.
//
// The method returns an error if something goes wrong. If the product is created
// successfully, the method returns nil.
func (r *productRepository) Create(ctx *gin.Context, product *entities.Product) error {
	return r.db.WithContext(ctx).Create(product).Error
}

// Retrieves a product by its ID from the database.
//
// The method takes a pointer to a *gin.Context and a uint as parameters. It
// returns a pointer to a entities.Product and an error. If something goes wrong,
// the method returns nil and an error.
//
// The method gets a product by its ID from the database using the given ID.
// The method returns a pointer to a entities.Product and an error. If the
// product is found, the method returns the product and nil. If the product is not
// found, the method returns nil and an error.
func (r *productRepository) GetByID(ctx *gin.Context, id uint) (*entities.Product, error) {
	var product entities.Product
	err := r.db.WithContext(ctx).First(&product, id).Error
	return &product, err
}

// GetAll gets all products from the database.
//
// The method takes a pointer to a *gin.Context as a parameter. It returns a slice
// of pointers to entities.Product and an error. If something goes wrong, the
// method returns an empty slice and an error.
//
// The method fetches all products from the database. If the products are
// successfully retrieved, the method returns the products and nil. If no
// products are found or an error occurs, the method returns an empty slice and
// an error.
func (r *productRepository) GetAll(ctx *gin.Context) ([]*entities.Product, error) {
	var products []*entities.Product
	err := r.db.WithContext(ctx).Find(&products).Error
	return products, err
}

// Updates a product in the database.
//
// The method takes a pointer to a *gin.Context and a pointer to an entities.Product
// as parameters. It returns an error if something goes wrong.
//
// The method updates a product in the database using the given product object.
// The product object is passed as a pointer and the method is responsible for updating
// a product in the database with the given attributes.
//
// The method returns an error if something goes wrong. If the product is updated
// successfully, the method returns nil.
func (r *productRepository) Update(ctx *gin.Context, product *entities.Product) error {
	return r.db.WithContext(ctx).Save(product).Error
}

// Deletes a product from the database by its ID.
//
// The method takes a pointer to a *gin.Context and a uint as parameters. It
// returns an error if something goes wrong.
//
// The method deletes a product from the database using the given ID. If the
// product is deleted successfully, the method returns nil. If the product is not
// found or an error occurs, the method returns an error.
func (r *productRepository) Delete(ctx *gin.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&entities.Product{}, id).Error
}

// Deletes multiple products from the database by their IDs.
//
// The method takes a pointer to a *gin.Context and a slice of uints as parameters.
// It returns an error if something goes wrong.
//
// The method deletes multiple products from the database using the given slice
// of IDs. The method returns an error if something goes wrong. If the products are
// deleted successfully, the method returns nil.
func (r *productRepository) DeleteAll(ctx *gin.Context, ids []uint) error {
	return r.db.WithContext(ctx).Delete(&entities.Product{}, ids).Error
}
