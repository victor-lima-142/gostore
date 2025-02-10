package repositories

import (
	"store/domain/entities"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// SupplierRepository is an interface that defines the methods that must
// be implemented by any data store that wants to interact with the suppliers
// table in the database.
//
// It provides methods for creating a new supplier, getting a supplier by its ID, getting all suppliers,
// updating a supplier, and deleting a supplier.
type SupplierRepository interface {
	Create(ctx *gin.Context, supplier *entities.Supplier) error    // Create a new supplier
	GetByID(ctx *gin.Context, id uint) (*entities.Supplier, error) // Get a supplier by ID
	GetAll(ctx *gin.Context) ([]*entities.Supplier, error)         // Get all suppliers
	Update(ctx *gin.Context, supplier *entities.Supplier) error    // Update a supplier
	Delete(ctx *gin.Context, id uint) error                        // Delete a supplier
	DeleteAll(ctx *gin.Context, ids []uint) error                  // Delete multiple suppliers
}

// supplierRepository is a struct that contains a pointer to a gorm DB instance and
// implements the SupplierRepository.
//
// The struct contains a pointer to a gorm DB instance which is used to interact
// with the suppliers table in the database.
type supplierRepository struct {
	db *gorm.DB
}

// NewSupplierRepository creates a new instance of supplierRepository with the provided
// database instance and returns it as a SupplierRepository.
// This function is used to initialize a new supplier repository that can perform
// CRUD operations and other queries on the suppliers table.
func NewSupplierRepository(db *gorm.DB) SupplierRepository {
	return &supplierRepository{db: db}
}

// Creates a new supplier in the database.
//
// The method takes a pointer to a *gin.Context and a pointer to a entities.Supplier
// as parameters. It returns an error if something goes wrong.
//
// The method creates a new supplier in the database using the given supplier object.
// The supplier object is passed as a pointer and the method is responsible for creating
// a new supplier in the database with the given attributes.
//
// The method returns an error if something goes wrong. If the supplier is created
// successfully, the method returns nil.
func (r *supplierRepository) Create(ctx *gin.Context, supplier *entities.Supplier) error {
	return r.db.WithContext(ctx).Create(supplier).Error
}

// Retrieves a supplier by its ID from the database.
//
// The method takes a pointer to a *gin.Context and a uint as parameters. It
// returns a pointer to a entities.Supplier and an error. If something goes wrong,
// the method returns nil and an error.
//
// The method gets a supplier by its ID from the database using the given ID.
// The method returns a pointer to a entities.Supplier and an error. If the
// supplier is found, the method returns the supplier and nil. If the supplier is
// not found, the method returns nil and an error.
func (r *supplierRepository) GetByID(ctx *gin.Context, id uint) (*entities.Supplier, error) {
	var supplier entities.Supplier
	err := r.db.WithContext(ctx).First(&supplier, id).Error
	return &supplier, err
}

// Retrieves all suppliers from the database.
//
// The method takes a pointer to a *gin.Context as a parameter. It returns a
// slice of pointers to entities.Supplier and an error. If something goes wrong,
// the method returns an empty slice and an error.
//
// The method retrieves all suppliers from the database. If the suppliers are
// found, the method returns the suppliers and nil. If no suppliers are found or
// an error occurs, the method returns an empty slice and an error.
func (r *supplierRepository) GetAll(ctx *gin.Context) ([]*entities.Supplier, error) {
	var suppliers []*entities.Supplier
	err := r.db.WithContext(ctx).Find(&suppliers).Error
	return suppliers, err
}

// Updates a supplier in the database.
//
// The method takes a pointer to a *gin.Context and a pointer to an entities.Supplier
// as parameters. It returns an error if something goes wrong.
//
// The method updates a supplier in the database using the given supplier object.
// The supplier object is passed as a pointer and the method is responsible for updating
// a supplier in the database with the given attributes.
//
// The method returns an error if something goes wrong. If the supplier is updated
// successfully, the method returns nil.
func (r *supplierRepository) Update(ctx *gin.Context, supplier *entities.Supplier) error {
	return r.db.WithContext(ctx).Save(supplier).Error
}

// Deletes a supplier by its ID from the database.
//
// The method takes a pointer to a *gin.Context and a uint as parameters. It
// returns an error if something goes wrong.
//
// The method deletes a supplier by its ID from the database using the given ID.
// If the supplier is deleted successfully, the method returns nil. If the supplier
// is not found or an error occurs, the method returns an error.
func (r *supplierRepository) Delete(ctx *gin.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&entities.Supplier{}, id).Error
}

// Deletes multiple suppliers from the database by their IDs.
//
// The method takes a pointer to a *gin.Context and a slice of uints as parameters.
// It returns an error if something goes wrong.
//
// The method deletes multiple suppliers from the database using the given slice
// of IDs. The method returns an error if something goes wrong. If the suppliers
// are deleted successfully, the method returns nil.
func (r *supplierRepository) DeleteAll(ctx *gin.Context, ids []uint) error {
	return r.db.WithContext(ctx).Delete(&entities.Supplier{}, ids).Error
}
