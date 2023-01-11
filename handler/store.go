package handler

import (
	"ecomm-alpha/database"
	"ecomm-alpha/models"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

// CreateStore creates a store for seller
//
//	@Summary		Register a new store data
//	@Description	Register store
//	@Tags
//	@Accept			json
//	@Produce		json
//	@Param			store body models.Store true "Register store"
//	@Success		200	{object}	ResponseHTTP{data=models.store}
//	Failure			400	{object}	ResponseHTTP{}
//	Failure			422	{object}	ResponseHTTP{}
//	Failure			500	{object}	ResponseHTTP{}
//	@Router			/api/v1/store/ [post]
func CreateStore(c *fiber.Ctx) error {
	db := database.DB
	user := c.Locals("user").(*jwt.Token).Claims.(jwt.MapClaims)
	// log.Println(user)

	store := new(models.Store)

	var statusCode int
	if err := c.BodyParser(store); err != nil {
		statusCode = GetStatusCodeFromError(err)
		log.Println(err)
		return c.Status(statusCode).JSON(ResponseHTTP{Success: false, Message: err.Error(), Data: nil})
	}

	if ok, errorFields := validateStoreInput(store); ok != true {
		return c.Status(fiber.StatusBadRequest).JSON(ResponseHTTP{Success: false, Message: "validation error", Data: errorFields})
	}

	store.SellerID = uint(user["seller_id"].(float64))

	if err := db.Create(store).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(ResponseHTTP{Success: false, Message: err.Error(), Data: nil})
	}

	return c.Status(fiber.StatusCreated).JSON(ResponseHTTP{Success: true, Message: "", Data: store})
}

// UpdateStore updates the store of seller
//
//	@Summary		Updates the store data
//	@Description	Update store
//	@Tags
//	@Accept			json
//	@Produce		json
//	@Param			store body models.Store true "Update store"
//	@Success		200	{object}	ResponseHTTP{data=string}
//	Failure			400	{object}	ResponseHTTP{}
//	Failure			422	{object}	ResponseHTTP{}
//	Failure			500	{object}	ResponseHTTP{}
//	@Router			/api/v1/store/ [put]

func UpdateStore(c *fiber.Ctx) error {
	db := database.DB

	store := new(models.Store)
	var statusCode int
	if err := c.BodyParser(store); err != nil {
		statusCode = GetStatusCodeFromError(err)
		if err != nil {
			log.Println(err)
		}
		return c.Status(statusCode).JSON(ResponseHTTP{Success: false, Message: err.Error(), Data: nil})
	}

	if ok, errorFields := validateStoreInput(store); ok != true {
		return c.Status(400).JSON(ResponseHTTP{Success: false, Message: "validation error", Data: errorFields})
	}

	// log.Println(store)

	var storeId string
	err := db.Raw("update store set name=?,description=? where id =? returning id;", store.Name, store.Description, store.ID).Scan(&storeId).Error
	if err != nil {
		log.Println(err)
		return c.Status(500).JSON(ResponseHTTP{Success: false, Message: "Internal Server Error", Data: nil})
	}

	return c.Status(201).JSON(ResponseHTTP{Success: true, Message: "", Data: storeId})
}

// PatchStoreName updates the name of store
//
//	@Summary		Updates the store name
//	@Description	Update store name
//	@Tags
//	@Accept			json
//	@Produce		json
//	@Param			name body string true "update store name"
//	@Success		200	{object}	ResponseHTTP{data=string}
//	Failure			400	{object}	ResponseHTTP{}
//	Failure			422	{object}	ResponseHTTP{}
//	Failure			500	{object}	ResponseHTTP{}
//	@Router			/api/v1/store/ [patch]
func PatchStoreName(c *fiber.Ctx) error {
	db := database.DB

	store := new(models.Store)

	var statusCode int
	if err := c.BodyParser(store); err != nil {
		statusCode = GetStatusCodeFromError(err)
		if err != nil {
			log.Println(err)
		}
		return c.Status(statusCode).JSON(ResponseHTTP{Success: false, Message: err.Error(), Data: nil})
	}

	if store.Name == "" {
		return c.Status(statusCode).JSON(ResponseHTTP{Success: false, Message: "validation error", Data: []string{"name"}})
	}

	var storeId string

	err := db.Raw("update store set name=?,where id=? returning id", store.Name, store.ID).Scan(&storeId).Error
	log.Println(store)

	if err != nil {
		return c.Status(500).JSON(ResponseHTTP{Success: false, Message: "Internal Server Error", Data: nil})
	}
	return c.Status(200).JSON(ResponseHTTP{Success: true, Message: "", Data: storeId})
}

// PatchStoreDescription updates description of store
//
//	@Summary		Updates the store description
//	@Description	Update store description
//	@Tags
//	@Accept			json
//	@Produce		json
//	@Param			description body string true "update store description"
//	@Success		200	{object}	ResponseHTTP{data=string}
//	Failure			400	{object}	ResponseHTTP{}
//	Failure			422	{object}	ResponseHTTP{}
//	Failure			500	{object}	ResponseHTTP{}
//	@Router			/api/v1/store/ [patch]
func PatchStoreDescription(c *fiber.Ctx) error {
	db := database.DB

	store := new(models.Store)

	var statusCode int
	if err := c.BodyParser(store); err != nil {
		statusCode = GetStatusCodeFromError(err)
		if err != nil {
			log.Println(err)
		}
		return c.Status(statusCode).JSON(ResponseHTTP{Success: false, Message: err.Error(), Data: nil})
	}

	if store.Description == "" {
		return c.Status(statusCode).JSON(ResponseHTTP{Success: false, Message: "validation error", Data: []string{"description"}})
	}

	var storeId string

	err := db.Raw("update store set description=?,where id=? returning id", store.Description, store.ID).Scan(&storeId).Error

	if err != nil {
		return c.Status(500).JSON(ResponseHTTP{Success: false, Message: "Internal Server Error", Data: nil})
	}

	return c.JSON(ResponseHTTP{Success: true, Message: "", Data: store})
}
