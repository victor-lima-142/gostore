package controllers

import (
	"net/http"
	"store/domain/entities"
	"store/services"
	"store/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// SupplierController is an interface that defines the methods for the supplier controller.
//
// The methods in this interface are used to create, retrieve, update, and delete
// suppliers in the database.
type SupplierController interface {
	CreateSupplier(ctx *gin.Context)
	GetAllSuppliers(ctx *gin.Context)
	GetSupplierByID(ctx *gin.Context)
	UpdateSupplier(ctx *gin.Context)
	DeleteSupplier(ctx *gin.Context)
	DeleteAllSuppliers(ctx *gin.Context)
}

// supplierController is a struct that contains a pointer to a supplierService and
// implements the SupplierController.
//
// The struct contains a pointer to a supplierService which is used to interact with
// the suppliers table in the database.
type supplierController struct {
	supplierService services.SupplierService
}

// NewSupplierController creates a new instance of supplierController with the provided
// supplierService and returns it as a SupplierController. This function is used to
// initialize a new supplier controller that can handle HTTP requests for CRUD operations
// on suppliers by utilizing the supplier service.
func NewSupplierController(supplierService services.SupplierService) SupplierController {
	return &supplierController{supplierService: supplierService}
}

// Handles the HTTP POST request to create a new supplier.
//
// The method expects a JSON payload representing the supplier details
// in the request body. It binds the JSON payload to a Supplier entity and
// invokes the supplier service to create a new supplier in the database.
//
// If the JSON binding fails, it responds with a 400 status code and an error message.
// If the creation fails due to a server error, it responds with a 500 status code
// and an error message. Upon successful creation, it responds with a 201 status code
// and the created supplier entity as JSON.
func (c *supplierController) CreateSupplier(ctx *gin.Context) {
	var supplier entities.Supplier

	if err := ctx.ShouldBindJSON(&supplier); err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	if err := c.supplierService.Create(ctx, &supplier); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, supplier)
}

// Handles the HTTP GET request to retrieve all suppliers from the database.
//
// The method invokes the supplier service to retrieve all suppliers from the
// database. If the retrieval fails due to a server error, it responds with a
// 500 status code and an error message. Otherwise, it responds with a 200 status
// code and the list of suppliers as JSON.
func (c *supplierController) GetAllSuppliers(ctx *gin.Context) {
	suppliers, err := c.supplierService.GetAll(ctx)

	if err == gorm.ErrRecordNotFound {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "No suppliers found"})
		return
	}
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, suppliers)
}

// Handles the HTTP GET request to retrieve a supplier by its ID.
//
// This method extracts the supplier ID from the URL parameters and
// calls the supplier service's GetByID method to retrieve the supplier from the database.
// If the supplier is found, it responds with a 200 status code and the supplier entity
// as JSON. If an error occurs during the retrieval, it responds with a 500 status code
// and an error message.
func (c *supplierController) GetSupplierByID(ctx *gin.Context) {
	id := ctx.Param("id")

	supplier, err := c.supplierService.GetByID(ctx, utils.StringToUint(id))
	if err == gorm.ErrRecordNotFound {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "Supplier not found"})
		return
	}
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, supplier)
}

// Handles the HTTP PUT request to update a supplier in the database.
//
// The method expects a JSON payload representing the supplier details
// in the request body. It binds the JSON payload to a Supplier entity and
// invokes the supplier service to update the supplier in the database.
//
// If the JSON binding fails, it responds with a 400 status code and an error message.
// If the update fails due to a server error, it responds with a 500 status code
// and an error message. Upon successful update, it responds with a 200 status code
// and the updated supplier entity as JSON.
func (c *supplierController) UpdateSupplier(ctx *gin.Context) {
	var supplier entities.Supplier

	if err := ctx.ShouldBindJSON(&supplier); err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	if err := c.supplierService.Update(ctx, &supplier); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, supplier)
}

// Handles the HTTP DELETE request for removing a supplier by its ID.
//
// This method extracts the supplier ID from the URL parameters and
// calls the Delete method of the supplier service to remove the supplier
// from the database. If the supplier is successfully deleted, it responds
// with a 200 status code and a success message. If an error occurs during
// the deletion, it responds with a 500 status code and an error message.
func (c *supplierController) DeleteSupplier(ctx *gin.Context) {
	id := ctx.Param("id")

	err := c.supplierService.Delete(ctx, utils.StringToUint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Supplier deleted successfully"})
}

// Handles the HTTP DELETE request for removing multiple suppliers by their IDs.
//
// This method extracts the supplier IDs from the URL query parameters and
// calls the DeleteAll method of the supplier service to remove the suppliers
// from the database. If the suppliers are successfully deleted, it responds
// with a 200 status code and a success message. If an error occurs during
// the deletion, it responds with a 500 status code and an error message.
func (c *supplierController) DeleteAllSuppliers(ctx *gin.Context) {
	ids := ctx.QueryArray("ids")

	err := c.supplierService.DeleteAll(ctx, utils.StringArrToUintArr(ids))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "All suppliers deleted successfully"})
}
