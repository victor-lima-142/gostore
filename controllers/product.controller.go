package controllers

import (
	"net/http"
	"store/domain/entities"
	"store/services"
	"store/utils"

	"github.com/gin-gonic/gin"
)

// ProductController is an interface that defines the methods for handling HTTP requests related to product operations.
//
// The methods in this interface are utilized to create, retrieve, update, and delete
// products in the database.
type ProductController interface {
	CreateProduct(ctx *gin.Context)     // Create a new product
	GetAllProducts(ctx *gin.Context)    // Get all products
	GetProductByID(ctx *gin.Context)    // Get a product by ID
	UpdateProduct(ctx *gin.Context)     // Update a product
	DeleteProduct(ctx *gin.Context)     // Delete a product
	DeleteAllProducts(ctx *gin.Context) // Delete all products
}

// productController is a struct that contains a pointer to a productService
// and implements the ProductController interface.
//
// The struct contains a pointer to a productService which is used to interact
// with the products table in the database.
type productController struct {
	productService services.ProductService
}

// NewProductController creates a new instance of productController with the provided
// productService and returns it as a ProductController. This function is used to
// initialize a new product controller that can handle HTTP requests for CRUD operations
// on products by utilizing the product service.
func NewProductController(productService services.ProductService) ProductController {
	return &productController{productService: productService}
}

func (c *productController) CreateProduct(ctx *gin.Context) {
	product := &entities.Product{}

	if err := ctx.ShouldBindJSON(product); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.productService.Create(ctx, product); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, product)
}

// Handles the HTTP request for retrieving all products from the database.
//
// This method takes a pointer to a *gin.Context as a parameter. It calls the GetAll
// method of the product service to retrieve all products from the database. If the
// retrieval fails, it returns a 500 error response. On success, it returns a 200 status
// code along with the products in the response body.
func (c *productController) GetAllProducts(ctx *gin.Context) {
	products, err := c.productService.GetAll(ctx)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, products)
}

// Handles the HTTP request for retrieving a product by its ID.
//
// This method takes a pointer to a *gin.Context as a parameter and extracts the
// product ID from the URL parameters. It then calls the GetByID method of the
// product service to retrieve the product from the database. If the product is
// found, it returns a 200 status code with the product in the response body.
// If an error occurs during the retrieval, it returns a 500 error response.
func (c *productController) GetProductByID(ctx *gin.Context) {
	id := ctx.Param("id")

	product, err := c.productService.GetByID(ctx, utils.StringToUint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, product)
}

// Handles the HTTP request for updating a product.
//
// This method takes a pointer to a *gin.Context as a parameter and binds the JSON
// request body to a product entity. If the request body is not valid JSON, it returns
// a 400 error response. It then calls the Update method of the product service to
// update the product in the database. If the update fails, it returns a 500 error
// response. On success, it returns a 200 status code along with the updated product
// in the response body.

func (c *productController) UpdateProduct(ctx *gin.Context) {
	product := &entities.Product{}

	if err := ctx.ShouldBindJSON(product); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.productService.Update(ctx, product); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, product)
}

// Handles the HTTP request for deleting a product by its ID.
//
// The method takes a pointer to a *gin.Context as a parameter and extracts the
// ID of the product to be deleted from the URL parameters. It then calls the
// Delete method of the product service to delete the product from the database.
// If the product is deleted successfully, the method returns a 200 status code
// with a message in the response body. If an error occurs during the deletion,
// the method returns a 500 error response.
func (c *productController) DeleteProduct(ctx *gin.Context) {
	id := ctx.Param("id")

	if err := c.productService.Delete(ctx, utils.StringToUint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully"})
}

// Handles the HTTP request for deleting all products by their IDs.
//
// The method takes a pointer to a *gin.Context as a parameter and extracts the
// IDs of the products to be deleted from the URL query parameters. It then calls
// the DeleteAll method of the product service to delete the products from the
// database. If the deletion is successful, it returns a 200 status code with a
// success message in the response body. If an error occurs during the deletion,
// it returns a 500 error response.
func (c *productController) DeleteAllProducts(ctx *gin.Context) {
	ids := ctx.QueryArray("ids")

	if err := c.productService.DeleteAll(ctx, utils.StringArrToUintArr(ids)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "All products deleted successfully"})
}
