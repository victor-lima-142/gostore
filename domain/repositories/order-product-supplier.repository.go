package repositories

import (
	"store/domain/entities"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// OrderProductSupplierRepository is an interface that defines the methods that must
// be implemented by any data store that wants to interact with the order_product_suppliers
// table in the database.
//
// It provides methods for creating a new orderProductSupplier, getting a orderProductSupplier by its ID,
// getting all orderProductSuppliers, updating a orderProductSupplier, and deleting a orderProductSupplier.
type OrderProductSupplierRepository interface {
	Create(ctx *gin.Context, orderProductSupplier *entities.OrderProductSupplier) error // Create a new orderProductSupplier
	GetByID(ctx *gin.Context, id uint) (*entities.OrderProductSupplier, error)          // Get a orderProductSupplier by ID
	GetAll(ctx *gin.Context) ([]*entities.OrderProductSupplier, error)                  // Get all orderProductSuppliers
	Update(ctx *gin.Context, orderProductSupplier *entities.OrderProductSupplier) error // Update a orderProductSupplier
	Delete(ctx *gin.Context, id uint) error                                             // Delete a orderProductSupplier
	DeleteAll(ctx *gin.Context, ids []uint) error                                       // Delete multiple orderProductSuppliers
}

// orderProductSupplierRepository is a struct that contains a pointer to a gorm DB instance and
// implements the OrderProductSupplierRepository.
//
// The struct contains a pointer to a gorm DB instance which is used to interact
// with the order_product_suppliers table in the database.
type orderProductSupplierRepository struct {
	db *gorm.DB
}

// NewOrderProductSupplierRepository creates a new instance of orderProductSupplierRepository with the provided
// database instance and returns it as an OrderProductSupplierRepository.
// This function is used to initialize a new orderProductSupplier repository that can perform
// CRUD operations and other queries on the order_product_suppliers table.
func NewOrderProductSupplierRepository(db *gorm.DB) OrderProductSupplierRepository {
	return &orderProductSupplierRepository{db: db}
}

// Creates a new orderProductSupplier in the database.
//
// The method takes a pointer to a *gin.Context and a pointer to a entities.OrderProductSupplier
// as parameters. It returns an error if something goes wrong.
//
// The method creates a new orderProductSupplier in the database using the given orderProductSupplier object.
// The orderProductSupplier object is passed as a pointer and the method is responsible for creating
// a new orderProductSupplier in the database with the given attributes.
//
// The method returns an error if something goes wrong. If the orderProductSupplier is created
// successfully, the method returns nil.
func (r *orderProductSupplierRepository) Create(ctx *gin.Context, orderProductSupplier *entities.OrderProductSupplier) error {
	return r.db.WithContext(ctx).Create(orderProductSupplier).Error
}

// Retrieves an orderProductSupplier by its ID from the database.
//
// The method takes a pointer to a *gin.Context and a uint as parameters. It
// returns a pointer to a entities.OrderProductSupplier and an error. If something goes wrong,
// the method returns nil and an error.
//
// The method gets an orderProductSupplier by its ID from the database using the given ID.
// The method returns a pointer to a entities.OrderProductSupplier and an error. If the
// orderProductSupplier is found, the method returns the orderProductSupplier and nil. If the
// orderProductSupplier is not found, the method returns nil and an error.
func (r *orderProductSupplierRepository) GetByID(ctx *gin.Context, id uint) (*entities.OrderProductSupplier, error) {
	var orderProductSupplier entities.OrderProductSupplier
	err := r.db.WithContext(ctx).First(&orderProductSupplier, id).Error
	return &orderProductSupplier, err
}

// Retrieves all orderProductSuppliers from the database.
//
// The method takes a pointer to a *gin.Context as a parameter. It returns a slice
// of pointers to entities.OrderProductSupplier and an error. If something goes wrong,
// the method returns an empty slice and an error.
//
// The method fetches all orderProductSuppliers from the database. If the
// orderProductSuppliers are found, the method returns the orderProductSuppliers and
// nil. If no orderProductSuppliers are found or an error occurs, the method returns
// an empty slice and an error.
func (r *orderProductSupplierRepository) GetAll(ctx *gin.Context) ([]*entities.OrderProductSupplier, error) {
	var orderProductSuppliers []*entities.OrderProductSupplier
	err := r.db.WithContext(ctx).Find(&orderProductSuppliers).Error
	return orderProductSuppliers, err
}

// Updates an orderProductSupplier in the database.
//
// The method takes a pointer to a *gin.Context and a pointer to an entities.OrderProductSupplier
// as parameters. It returns an error if something goes wrong.
//
// The method updates an orderProductSupplier in the database using the given orderProductSupplier object.
// The orderProductSupplier object is passed as a pointer and the method is responsible for updating
// an orderProductSupplier in the database with the given attributes.
//
// The method returns an error if something goes wrong. If the orderProductSupplier is updated
// successfully, the method returns nil.
func (r *orderProductSupplierRepository) Update(ctx *gin.Context, orderProductSupplier *entities.OrderProductSupplier) error {
	return r.db.WithContext(ctx).Save(orderProductSupplier).Error
}

// Deletes an orderProductSupplier by its ID from the database.
//
// The method takes a pointer to a *gin.Context and a uint as parameters. It
// returns an error if something goes wrong.
//
// The method deletes an orderProductSupplier by its ID from the database using the
// given ID. If the orderProductSupplier is deleted successfully, the method returns
// nil. If the orderProductSupplier is not found or an error occurs, the method returns
// an error.
func (r *orderProductSupplierRepository) Delete(ctx *gin.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&entities.OrderProductSupplier{}, id).Error
}

// Deletes multiple orderProductSuppliers from the database by their IDs.
//
// The method takes a pointer to a *gin.Context and a slice of uints as parameters.
// It returns an error if something goes wrong.
//
// The method deletes multiple orderProductSuppliers from the database using the given slice
// of IDs. The method returns an error if something goes wrong. If the orderProductSuppliers are
// deleted successfully, the method returns nil.
func (r *orderProductSupplierRepository) DeleteAll(ctx *gin.Context, ids []uint) error {
	return r.db.WithContext(ctx).Delete(&entities.OrderProductSupplier{}, ids).Error
}
