package services

import (
	"store/domain/entities"
	"store/domain/repositories"

	"github.com/gin-gonic/gin"
)

// OrderService defines the methods that a service must implement to manage
// orders in the application. It provides methods to create, retrieve, update,
// and delete order entities.
type OrderService interface {
	Create(ctx *gin.Context, order *entities.Order) error       // Create a new order
	GetByID(ctx *gin.Context, id uint) (*entities.Order, error) // Get an order by ID
	GetAll(ctx *gin.Context) ([]*entities.Order, error)         // Get all orders
	Update(ctx *gin.Context, order *entities.Order) error       // Update an order
	Delete(ctx *gin.Context, id uint) error                     // Delete an order
	DeleteAll(ctx *gin.Context, ids []uint) error               // Delete multiple orders
}

// orderService is a struct that contains a pointer to an OrderRepository
// and implements the OrderService interface.
// It is used to manage orders in the application.
type orderService struct {
	orderRepository repositories.OrderRepository
}

// NewOrderService creates a new OrderService with the given OrderRepository.
// It returns an instance of orderService that implements the OrderService interface,
// allowing for the management of orders in the application.
func NewOrderService(orderRepository repositories.OrderRepository) OrderService {
	return &orderService{orderRepository: orderRepository}
}

// Create a new order in the database.
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
func (s *orderService) Create(ctx *gin.Context, order *entities.Order) error {
	return s.orderRepository.Create(ctx, order)
}

// Retrieves an order from the database by its ID.
//
// The method takes a pointer to a *gin.Context for handling request-specific data
// and a uint representing the ID of the order to be retrieved. It returns a
// pointer to an entities.Order if the order is found, and an error if the
// order is not found or if something goes wrong during the retrieval process.
//
// This method ensures that the order is retrieved from the database with the
// provided ID. If the order is found, it returns the order and nil; otherwise,
// it returns nil and an error.
func (s *orderService) GetByID(ctx *gin.Context, id uint) (*entities.Order, error) {
	return s.orderRepository.GetByID(ctx, id)
}

// Retrieves all orders from the database.
//
// The method takes a pointer to a *gin.Context as a parameter. It returns a slice
// of pointers to entities.Order and an error. If something goes wrong, the
// method returns an empty slice and an error.
//
// The method fetches all orders from the database. If the orders are found,
// the method returns the orders and nil. If no orders are found or an error
// occurs, the method returns an empty slice and an error.
func (s *orderService) GetAll(ctx *gin.Context) ([]*entities.Order, error) {
	return s.orderRepository.GetAll(ctx)
}

// Updates an order in the database.
//
// The method takes a pointer to a *gin.Context and a pointer to an entities.Order
// as parameters. It returns an error if something goes wrong.
//
// The method updates an order in the database using the given order object.
// The order object is passed as a pointer and the method is responsible for updating
// an order in the database with the given attributes.
//
// The method returns an error if something goes wrong. If the order is updated
// successfully, the method returns nil.
func (s *orderService) Update(ctx *gin.Context, order *entities.Order) error {
	return s.orderRepository.Update(ctx, order)
}

// Deletes an order by its ID from the database.
//
// The method takes a pointer to a *gin.Context and a uint as parameters. It
// returns an error if something goes wrong.
//
// The method deletes an order by its ID from the database using the given ID.
// If the order is deleted successfully, the method returns nil. If the order is
// not found or an error occurs, the method returns an error.
func (s *orderService) Delete(ctx *gin.Context, id uint) error {
	return s.orderRepository.Delete(ctx, id)
}

// Deletes multiple orders by their IDs from the database.
//
// The method takes a pointer to a *gin.Context and a slice of uints as parameters.
// It returns an error if something goes wrong.
//
// The method deletes multiple orders from the database using the given slice
// of IDs. The method returns an error if something goes wrong. If the orders are
// deleted successfully, the method returns nil.
func (s *orderService) DeleteAll(ctx *gin.Context, ids []uint) error {
	return s.orderRepository.DeleteAll(ctx, ids)
}
