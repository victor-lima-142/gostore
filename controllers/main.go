package controllers

import (
	"store/domain/repositories"
	"store/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Sets up the HTTP route handlers for customer-related operations.
//
// It initializes the customer repository, service, and controller, and binds
// the HTTP endpoints to their corresponding handler functions. The following
// routes are registered:
//
// - GET /customers: Retrieve a list of all customers.
//
// - GET /customers/:id: Retrieve a customer by its ID.
//
// - POST /customers: Create a new customer.
//
// - PUT /customers/:id: Update an existing customer by its ID.
//
// - DELETE /customers/:id: Delete a customer by its ID.
func customerRoutes(app *gin.Engine, db *gorm.DB) {
	customerRepository := repositories.NewCustomerRepository(db)
	customerService := services.NewCustomerService(customerRepository)
	controller := NewCustomerController(customerService)

	app.GET("/customers", controller.GetAllCustomers)
	app.GET("/customers/:id", controller.GetCustomerByID)
	app.POST("/customers", controller.CreateCustomer)
	app.PUT("/customers/:id", controller.UpdateCustomer)
	app.DELETE("/customers/:id", controller.DeleteCustomer)
}

// Sets up the HTTP route handlers for supplier-related operations.
//
// It initializes the supplier repository, service, and controller, and binds
// the HTTP endpoints to their corresponding handler functions. The following
// routes are registered:
//
// - GET /suppliers: Retrieve a list of all suppliers.
//
// - GET /suppliers/:id: Retrieve a supplier by its ID.
//
// - POST /suppliers: Create a new supplier.
//
// - PUT /suppliers/:id: Update an existing supplier by its ID.
//
// - DELETE /suppliers/:id: Delete a supplier by its ID.
func supplierRoutes(app *gin.Engine, db *gorm.DB) {
	supplierRepository := repositories.NewSupplierRepository(db)
	supplierService := services.NewSupplierService(supplierRepository)
	controller := NewSupplierController(supplierService)

	app.GET("/suppliers", controller.GetAllSuppliers)
	app.GET("/suppliers/:id", controller.GetSupplierByID)
	app.POST("/suppliers", controller.CreateSupplier)
	app.PUT("/suppliers/:id", controller.UpdateSupplier)
	app.DELETE("/suppliers/:id", controller.DeleteSupplier)
}

// Sets up the HTTP route handlers for product-related operations.
//
// It initializes the product repository, service, and controller, and binds
// the HTTP endpoints to their corresponding handler functions. The following
// routes are registered:
//
// - GET /products: Retrieve a list of all products.
//
// - GET /products/:id: Retrieve a product by its ID.
//
// - POST /products: Create a new product.
//
// - PUT /products/:id: Update an existing product by its ID.
//
// - DELETE /products/:id: Delete a product by its ID.
func productRoutes(app *gin.Engine, db *gorm.DB) {
	productRepository := repositories.NewProductRepository(db)
	productService := services.NewProductService(productRepository)
	controller := NewProductController(productService)

	app.GET("/products", controller.GetAllProducts)
	app.GET("/products/:id", controller.GetProductByID)
	app.POST("/products", controller.CreateProduct)
	app.PUT("/products/:id", controller.UpdateProduct)
	app.DELETE("/products/:id", controller.DeleteProduct)
}

// Sets up the HTTP route handlers for order-related operations.
//
// It initializes the order repository, service, and controller, and binds
// the HTTP endpoints to their corresponding handler functions. The following
// routes are registered:
//
// - GET /orders: Retrieve a list of all orders.
//
// - GET /orders/:id: Retrieve an order by its ID.
//
// - POST /orders: Create a new order.
//
// - PUT /orders/:id: Update an existing order by its ID.
//
// - DELETE /orders/:id: Delete an order by its ID.
func orderRoutes(app *gin.Engine, db *gorm.DB) {
	orderRepository := repositories.NewOrderRepository(db)
	orderService := services.NewOrderService(orderRepository)
	controller := NewOrderController(orderService)

	app.GET("/orders", controller.GetAllOrders)
	app.GET("/orders/:id", controller.GetOrderByID)
	app.POST("/orders", controller.CreateOrder)
	app.PUT("/orders/:id", controller.UpdateOrder)
	app.DELETE("/orders/:id", controller.DeleteOrder)
}

// InitRoutes initializes all routes for the application.
//
// It sets up the routes for customers, suppliers, products, and orders.
func InitRoutes(app *gin.Engine, db *gorm.DB) {
	customerRoutes(app, db)
	supplierRoutes(app, db)
	productRoutes(app, db)
	orderRoutes(app, db)
}
