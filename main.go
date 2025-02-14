package main

import (
	"log"
	"net/http"
	"store/controllers"
	"store/domain/entities"
	"sync"

	"github.com/gin-gonic/gin"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Main starts the Gin server with the API routes and database connection.
func main() {
	app := gin.Default()
	app.Use(JSONMiddleware())
	db := GetDB()
	controllers.InitRoutes(app, db)
	app.Run(":8080")
}

var (
	db   *gorm.DB
	once sync.Once
)

// GetDB initializes and returns a singleton instance of the database connection.
// It uses the Once pattern to ensure the database connection is established only once.
// The function constructs a DSN (Data Source Name) for a PostgreSQL database and attempts
// to open a connection using GORM. If the connection fails, it logs a fatal error.
// Upon successful connection, it calls AutoMigrate to auto-migrate database tables.
// The function returns a pointer to the gorm.DB instance.
func GetDB() *gorm.DB {
	once.Do(func() {
		dsn := "host=localhost user=postgres password=CYvr9tNwEbalWAZPsMiC dbname=gostore port=5432 sslmode=disable TimeZone=UTC"
		conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatalf("Failed to connect to database: %v", err)
		}
		db = conn
		AutoMigrate()
		log.Println("Database connection established")
	})
	return db
}

// AutoMigrate performs the auto-migration of the tables in the database. It is
// called by the GetDB method when the database connection is established. It
// auto-migrates the tables for the Customer, Supplier, Product, Order, Contact,
// ProductSupplier, and OrderProductSupplier entities. The method checks if the
// database connection is initialized and logs a fatal error if it is not. It
// also logs a fatal error if the migration fails. If the migration is successful,
// it logs a message to the console.
func AutoMigrate() {
	if db == nil {
		log.Fatal("Database connection is not initialized")
	}
	err := db.AutoMigrate(
		&entities.Customer{},             // Add the Customer entity
		&entities.Supplier{},             // Add the Supplier entity
		&entities.Product{},              // Add the Product entity
		&entities.Order{},                // Add the Order entity
		&entities.Contact{},              // Add the Contact entity
		&entities.ProductSupplier{},      // Add the ProductSupplier entity
		&entities.OrderProductSupplier{}, // Add the OrderProductSupplier entity
	)
	if err != nil {
		log.Fatalf("Migration failed: %v", err)
	}
	log.Println("AutoMigrate completed successfully")
}

// JSONMiddleware is a middleware function that sets the Accept header to "application/json"
// and the Content-Type header to "application/json". It is used to ensure that the responses
// are in JSON format.
//
// Its aborts the request if the Content-Type header is not "application/json".
func JSONMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.GetHeader("Content-Type") != "application/json" {
			c.JSON(http.StatusUnsupportedMediaType, gin.H{"error": "Content-Type must be application/json"})
			c.Abort()
			return
		}
		c.Request.Header.Set("Accept", "application/json")
		c.Header("Content-Type", "application/json")
		c.Next()
	}
}
