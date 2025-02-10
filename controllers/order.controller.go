package controllers

import (
	"net/http"
	"store/domain/entities"
	"store/services"
	"store/utils"

	"github.com/gin-gonic/gin"
)

// OrderController is an interface that defines the methods for handling HTTP requests related to order operations.
//
// The methods in this interface are utilized to create, retrieve, update, and delete
// orders in the database.
type OrderController interface {
	CreateOrder(ctx *gin.Context)     // Create a new order
	GetOrderByID(ctx *gin.Context)    // Get an order by id
	UpdateOrder(ctx *gin.Context)     // Update an order
	DeleteOrder(ctx *gin.Context)     // Delete an order
	GetAllOrders(ctx *gin.Context)    // Get all orders
	DeleteAllOrders(ctx *gin.Context) // Delete all orders
}

// orderController is a struct that contains a pointer to an OrderService.
// It is used to manage orders in the application.
//
// The struct contains a pointer to a orderService which is used to interact
// with the orders table in the database.
type orderController struct {
	orderService services.OrderService
}

func NewOrderController(orderService services.OrderService) OrderController {
	return &orderController{orderService: orderService}
}

// Handles the HTTP request for creating a new order.
//
// The method takes a pointer to a *gin.Context as a parameter. It binds the
// request body to a new entities.Order and calls the Create method of the
// order service to create a new order in the database. If the order is created
// successfully, the method returns a 201 status code with the created order in
// the response body. If an error occurs during the creation, the method returns
// a 500 error response.
func (c *orderController) CreateOrder(ctx *gin.Context) {
	var order entities.Order

	if err := ctx.ShouldBindJSON(&order); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.orderService.Create(ctx, &order); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, order)
}

// Handles the HTTP request for retrieving an order by its ID.
//
// The method takes a pointer to a *gin.Context as a parameter and extracts the
// ID of the order to be retrieved from the URL parameters. It then calls the
// GetByID method of the order service to retrieve the order from the database.
// If the order is found, the method returns a 200 status code with the order in
// the response body. If an error occurs during the retrieval, the method
// returns a 500 error response.
func (c *orderController) GetOrderByID(ctx *gin.Context) {
	id := ctx.Param("id")

	order, err := c.orderService.GetByID(ctx, utils.StringToUint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, order)
}

// Handles the HTTP request for updating an order.
//
// The method takes a pointer to a *gin.Context as a parameter and extracts the
// order to be updated from the request body. It then calls the Update method
// of the order service to update the order in the database. If the order is
// updated successfully, the method returns a 200 status code with the updated
// order in the response body. If an error occurs during the update, the method
// returns a 500 error response.
func (c *orderController) UpdateOrder(ctx *gin.Context) {
	var order entities.Order

	if err := ctx.ShouldBindJSON(&order); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.orderService.Update(ctx, &order); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, order)
}

// Handles the HTTP request for deleting an order by its ID.
//
// The method takes a pointer to a *gin.Context as a parameter and extracts the
// ID of the order to be deleted from the URL parameters. It then calls the
// Delete method of the order service to delete the order from the database.
// If the order is deleted successfully, the method returns a 200 status code
// with a message in the response body. If an error occurs during the deletion,
// the method returns a 500 error response.
func (c *orderController) DeleteOrder(ctx *gin.Context) {
	id := ctx.Param("id")

	if err := c.orderService.Delete(ctx, utils.StringToUint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Order deleted successfully"})
}

// Handles the HTTP request for retrieving all orders from the database.
//
// This method takes a pointer to a *gin.Context as a parameter. It calls the GetAll
// method of the order service to retrieve all orders from the database. If the
// retrieval fails, it returns a 500 error response. On success, it returns a 200 status
// code along with the orders in the response body.
func (c *orderController) GetAllOrders(ctx *gin.Context) {
	orders, err := c.orderService.GetAll(ctx)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, orders)
}

// Handles the HTTP request for deleting multiple orders by their IDs.
//
// This method takes a pointer to a *gin.Context as a parameter and extracts the
// order IDs from the URL query parameters. It then calls the DeleteAll method
// of the order service to delete the orders from the database. If the deletion
// is successful, it returns a 200 status code with a message in the response body.
// If an error occurs during the deletion, it returns a 500 error response.
func (c *orderController) DeleteAllOrders(ctx *gin.Context) {
	ids := ctx.QueryArray("ids")

	if err := c.orderService.DeleteAll(ctx, utils.StringArrToUintArr(ids)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "All orders deleted successfully"})
}
