package controllers

import (
	"store/domain/entities"
	"store/services"
	"store/utils"

	"github.com/gin-gonic/gin"
)

// ContactController is an interface that defines the methods for the contact controller.
//
// The methods in this interface are used to create, retrieve, update, and delete
// contacts in the database.
type ContactController interface {
	CreateContact(ctx *gin.Context)     // Create a new contact
	GetAllContacts(ctx *gin.Context)    // Get all contacts
	GetContactByID(ctx *gin.Context)    // Get a contact by ID
	UpdateContact(ctx *gin.Context)     // Update a contact
	DeleteContact(ctx *gin.Context)     // Delete a contact
	DeleteAllContacts(ctx *gin.Context) // Delete all contacts
}

// contactController is a struct that contains a pointer to a contactService and
// implements the ContactController.
//
// The struct contains a pointer to a contactService which is used to interact with
// the contacts table in the database.
type contactController struct {
	contactService services.ContactService
}

// NewContactController creates a new instance of contactController with the provided
// contactService and returns it as a ContactController. This function is used to
// initialize a new contact controller that can handle HTTP requests for CRUD operations
// on contacts by utilizing the contact service.
func NewContactController(contactService services.ContactService) ContactController {
	return &contactController{contactService: contactService}
}

// Handles the HTTP request for creating a new contact.
//
// This method takes a pointer to a *gin.Context as a parameter and binds the JSON
// request body to a contact entity. If the request body is not valid JSON, it returns
// a 400 error response. It then calls the Create method of the contact service to
// create the contact in the database. If the creation fails, it returns a 500 error
// response. On success, it returns a 201 status code along with the created contact
// in the response body.
func (c *contactController) CreateContact(ctx *gin.Context) {
	var contact entities.Contact

	err := ctx.ShouldBindJSON(&contact)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
	}

	err = c.contactService.Create(ctx, &contact)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
	}

	ctx.JSON(201, contact)
}

// Retrieves all contacts from the database.
//
// The method takes a pointer to a *gin.Context as a parameter. It returns a slice
// of pointers to entities.Contact and an error. If something goes wrong, the
// method returns an empty slice and an error.
//
// The method returns a slice of pointers to entities.Contact and an error. If the
// contacts are found, the method returns the contacts and nil. If no contacts are
// found or an error occurs, the method returns an empty slice and an error.
func (c *contactController) GetAllContacts(ctx *gin.Context) {
	contacts, err := c.contactService.GetAll(ctx)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
	}

	ctx.JSON(200, contacts)
}

// Handles the HTTP request for retrieving a contact by its ID.
//
// This method takes a pointer to a *gin.Context as a parameter and extracts the
// contact ID from the URL parameters. It then calls the GetByID method of the
// contact service to retrieve the contact from the database. If the contact is
// found, it returns a 200 status code with the contact in the response body.
// If an error occurs during the retrieval, it returns a 500 error response.
func (c *contactController) GetContactByID(ctx *gin.Context) {
	id := ctx.Param("id")

	contact, err := c.contactService.GetByID(ctx, utils.StringToUint(id))
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
	}

	ctx.JSON(200, contact)
}

// Handles the HTTP request for updating a contact.
//
// This method takes a pointer to a *gin.Context as a parameter and extracts the
// contact data from the request body. It then calls the Update method of the
// contact service to update the contact in the database. If the contact is
// updated successfully, it returns a 200 status code with the updated contact
// in the response body. If an error occurs during the update, it returns a 500
// error response.
func (c *contactController) UpdateContact(ctx *gin.Context) {
	var contact entities.Contact

	err := ctx.ShouldBindJSON(&contact)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
	}

	err = c.contactService.Update(ctx, &contact)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
	}

	ctx.JSON(200, contact)
}

// Handles the HTTP request for deleting a contact by its ID.
//
// This method takes a pointer to a *gin.Context as a parameter and extracts the
// contact ID from the URL parameters. It then calls the Delete method of the
// contact service to delete the contact from the database. If the deletion is
// successful, it returns a 200 status code with a success message in the response
// body. If an error occurs during the deletion, it returns a 500 error response.
func (c *contactController) DeleteContact(ctx *gin.Context) {
	id := ctx.Param("id")

	err := c.contactService.Delete(ctx, utils.StringToUint(id))
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
	}

	ctx.JSON(200, gin.H{"message": "Contact deleted successfully"})
}

// Handles the HTTP request for deleting multiple contacts by their IDs.
//
// This method takes a pointer to a *gin.Context as a parameter and extracts the
// IDs of the contacts to be deleted from the URL query parameters. It then calls
// the DeleteAll method of the contact service to delete the contacts from the
// database. If the deletion is successful, it returns a 200 status code with a
// success message in the response body. If an error occurs during the deletion,
// it returns a 500 error response.
func (c *contactController) DeleteAllContacts(ctx *gin.Context) {
	ids := ctx.QueryArray("ids")

	err := c.contactService.DeleteAll(ctx, utils.StringArrToUintArr(ids))
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
	}

	ctx.JSON(200, gin.H{"message": "All contacts deleted successfully"})
}
