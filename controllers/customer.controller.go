package controllers

import (
	"net/http"
	"store/domain/entities"
	"store/services"
	"store/utils"

	"github.com/gin-gonic/gin"
)

// CustomerController is an interface that defines the methods for the customer controller.
//
// The methods in this interface are used to create, retrieve, update, and delete
// the customers in the database.
type CustomerController interface {
	CreateCustomer(ctx *gin.Context)     // Create a new customer
	GetAllCustomers(ctx *gin.Context)    // Get all customers
	GetCustomerByID(ctx *gin.Context)    // Get a customer by ID
	UpdateCustomer(ctx *gin.Context)     // Update a customer
	DeleteCustomer(ctx *gin.Context)     // Delete a customer
	DeleteAllCustomers(ctx *gin.Context) // Delete multiple customers
}

// customerController is a struct that contains a pointer to a customerService and
// implements the CustomerController.
//
// The struct contains a pointer to a customerService which is used to interact with
// the customers table in the database.
type customerController struct {
	customerService services.CustomerService
}

// NewCustomerController creates a new instance of customerController with the provided
// customerService and returns it as a CustomerController. This function is used to
// initialize a new customer controller that can handle HTTP requests for CRUD operations
// on customers by utilizing the customer service.
func NewCustomerController(customerService services.CustomerService) CustomerController {
	return &customerController{customerService: customerService}
}

// Handles the HTTP request for creating a new customer.
//
// This method takes a pointer to a *gin.Context as a parameter and binds the JSON
// request body to a customer entity. If the request body is not valid JSON, it returns
// a 400 error response. It then calls the Create method of the customer service to
// create the customer in the database. If the creation fails, it returns a 500 error
// response. On success, it returns a 201 status code along with the created customer
// in the response body.
func (c *customerController) CreateCustomer(ctx *gin.Context) {
	var customer entities.Customer

	if err := ctx.ShouldBindJSON(&customer); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.customerService.Create(ctx, &customer); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, customer)
}

// Handles the HTTP request for getting all customers from the database.
//
// This method takes a pointer to a *gin.Context as a parameter. It calls the GetAll
// method of the customer service to get all customers from the database. If the
// retrieval fails, it returns a 500 error response. On success, it returns a 200 status
// code along with the customers in the response body.
func (c *customerController) GetAllCustomers(ctx *gin.Context) {
	customers, err := c.customerService.GetAll(ctx)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, customers)
}

// Handles the HTTP request for retrieving a customer by its ID.
//
// This method takes a pointer to a *gin.Context as a parameter and extracts the
// customer ID from the URL parameters. It then calls the GetByID method of the
// customer service to retrieve the customer from the database. If the customer is
// found, it returns a 200 status code with the customer in the response body.
// If an error occurs during the retrieval, it returns a 500 error response.
func (c *customerController) GetCustomerByID(ctx *gin.Context) {
	id := ctx.Param("id")

	contact, err := c.customerService.GetByID(ctx, utils.StringToUint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, contact)
}

// Handles the HTTP request for updating a customer.
//
// This method takes a pointer to a *gin.Context as a parameter and binds the JSON
// request body to a customer entity. If the request body is not valid JSON, it returns
// a 400 error response. It then calls the Update method of the customer service to
// update the customer in the database. If the creation fails, it returns a 500 error
// response. On success, it returns a 200 status code along with the updated customer
// in the response body.
func (c *customerController) UpdateCustomer(ctx *gin.Context) {
	var customer entities.Customer

	if err := ctx.ShouldBindJSON(&customer); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.customerService.Update(ctx, &customer); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, customer)
}

// Handles the HTTP request for deleting a customer by its ID.
//
// This method takes a pointer to a *gin.Context as a parameter and extracts the
// customer ID from the URL parameters. It then calls the Delete method of the
// customer service to delete the customer from the database. If the customer is
// found, it returns a 200 status code with a message in the response body. If an
// error occurs during the deletion, it returns a 500 error response.
func (c *customerController) DeleteCustomer(ctx *gin.Context) {
	id := ctx.Param("id")

	err := c.customerService.Delete(ctx, utils.StringToUint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Customer deleted successfully"})
}

// Handles the HTTP request for deleting multiple customers by their IDs.
//
// This method takes a pointer to a *gin.Context as a parameter and extracts the
// customer IDs from the URL query parameters. It then calls the DeleteAll method
// of the customer service to delete the customers from the database. If the
// deletion is successful, it returns a 200 status code with a message in the
// response body. If an error occurs during the deletion, it returns a 500 error
// response.
func (c *customerController) DeleteAllCustomers(ctx *gin.Context) {
	ids := ctx.QueryArray("ids")

	err := c.customerService.DeleteAll(ctx, utils.StringArrToUintArr(ids))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "All customers deleted successfully"})
}
