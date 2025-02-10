package repositories

import (
	"store/domain/entities"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// OrderRepository is an interface that defines the methods that must
// be implemented by any data store that wants to interact with the orders
// table in the database.
//
// It provides methods for creating a new order, getting an order by its ID, getting all orders,
// updating an order, deleting an order, and getting an order with its order products.
type OrderRepository interface {
	Create(ctx *gin.Context, order *entities.Order) error                         // Create a new order
	GetByID(ctx *gin.Context, id uint) (*entities.Order, error)                   // Get an order by ID
	GetAll(ctx *gin.Context) ([]*entities.Order, error)                           // Get all orders
	Update(ctx *gin.Context, order *entities.Order) error                         // Update an order
	Delete(ctx *gin.Context, id uint) error                                       // Delete an order
	DeleteAll(ctx *gin.Context, ids []uint) error                                 // Delete multiple orders
	GetOrderWithOrderProducts(ctx *gin.Context, id uint) (*entities.Order, error) // Get an order with its order products
}

// orderRepository is a struct that contains a pointer to a gorm DB instance
// and implements the OrderRepository.
//
// The struct contains a pointer to a gorm DB instance which is used to interact
// with the orders table in the database.
type orderRepository struct {
	db *gorm.DB
}

// NewOrderRepository creates a new instance of orderRepository with the provided
// database instance and returns it as an OrderRepository.
// This function is used to initialize a new order repository that can perform
// CRUD operations and other queries on the orders table.
func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &orderRepository{db: db}
}

// Creates a new order in the database.
//
// This method is part of the OrderRepository and is used to create a
// new order in the database.
//
// The method takes a pointer to a *gin.Context and a pointer to an entities.Order
// as parameters. It returns an error if something goes wrong.
//
// The method creates a new order in the database using the given order object.
// The order object is passed as a pointer and the method is responsible for creating
// a new order in the database with the given attributes.
//
// The method returns an error if something goes wrong. If the order is created
// successfully, the method returns nil.
func (r *orderRepository) Create(ctx *gin.Context, order *entities.Order) error {
	return r.db.WithContext(ctx).Create(order).Error
}

// Retrieves an order by its ID from the database.
//
// This method is part of the OrderRepository and is used to get an order
// by its ID from the database.
//
// The method takes a pointer to a *gin.Context and a uint as parameters. It
// returns a pointer to an entities.Order and an error. If something goes wrong,
// the method returns nil and an error.
//
// The method gets an order by its ID from the database using the given ID.
// The method returns a pointer to an entities.Order and an error. If the order is
// found, the method returns the order and nil. If the order is not found, the
// method returns nil and an error.
func (r *orderRepository) GetByID(ctx *gin.Context, id uint) (*entities.Order, error) {
	var order entities.Order
	err := r.db.WithContext(ctx).First(&order, id).Error
	return &order, err
}

// Retrieves all orders from the database.
//
// This method is part of the OrderRepository and is used to retrieve all
// orders from the database.
//
// The method takes a pointer to a *gin.Context as a parameter. It returns a slice
// of pointers to entities.Order and an error. If something goes wrong, the method
// returns an empty slice and an error.
//
// The method fetches all orders from the database. If the orders are successfully
// retrieved, the method returns the orders and nil. If no orders are found or
// an error occurs, the method returns an empty slice and an error.
func (r *orderRepository) GetAll(ctx *gin.Context) ([]*entities.Order, error) {
	var orders []*entities.Order
	err := r.db.WithContext(ctx).Find(&orders).Error
	return orders, err
}

// Updates an order in the database.
//
// This method is part of the OrderRepository and is used to update an
// order in the database.
//
// The method takes a pointer to a *gin.Context and a pointer to an entities.Order
// as parameters. It returns an error if something goes wrong.
//
// The method updates an order in the database using the given order object. The
// order object is passed as a pointer and the method is responsible for updating
// an order in the database with the given attributes.
//
// The method returns an error if something goes wrong. If the order is updated
// successfully, the method returns nil.
func (r *orderRepository) Update(ctx *gin.Context, order *entities.Order) error {
	return r.db.WithContext(ctx).Save(order).Error
}

// Deletes an order by its ID from the database.
//
// This method takes a pointer to a *gin.Context and a uint as parameters. It
// returns an error if something goes wrong.
//
// The method deletes an order by its ID from the database using the given ID.
// If the order is deleted successfully, the method returns nil. If the order is
// not found or an error occurs, the method returns an error.
func (r *orderRepository) Delete(ctx *gin.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&entities.Order{}, id).Error
}

// Deletes multiple orders from the database by their IDs.
//
// The method takes a pointer to a *gin.Context and a slice of uints as parameters.
// It returns an error if something goes wrong.
//
// The method deletes multiple orders from the database using the given slice
// of IDs. The method returns an error if something goes wrong. If the orders are
// deleted successfully, the method returns nil.
func (r *orderRepository) DeleteAll(ctx *gin.Context, ids []uint) error {
	return r.db.WithContext(ctx).Delete(&entities.Order{}, ids).Error
}

// Retrieves an order by its ID from the database, including its order products.
//
// The method takes a pointer to a *gin.Context and a uint as parameters. It
// returns a pointer to an entities.Order and an error. If something goes wrong,
// the method returns nil and an error.
//
// The method gets an order by its ID from the database using the given ID, and
// preloads the OrderProducts field. The method returns a pointer to an
// entities.Order and an error. If the order is found, the method returns the
func (r *orderRepository) GetOrderWithOrderProducts(ctx *gin.Context, id uint) (*entities.Order, error) {
	var order entities.Order
	err := r.db.WithContext(ctx).Preload("OrderProducts").First(&order, id).Error
	return &order, err
}
