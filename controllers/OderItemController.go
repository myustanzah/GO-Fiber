package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/myustanzah/GO-Fiber.git/models"
	"gorm.io/gorm"
)

type OrderItemController struct {
	DB *gorm.DB
}

func NewOrderItemController(db *gorm.DB) *OrderItemController {
	return &OrderItemController{
		DB: db,
	}
}

func (oic *OrderItemController) GetAllOrderItems(c *fiber.Ctx) error {
	var orderItems []models.OrderItem
	if err := oic.DB.Find(&orderItems).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve order items",
		})
	}
	return c.JSON(orderItems)
}

func (oic *OrderItemController) GetOrderItemByID(c *fiber.Ctx) error {
	var orderItem models.OrderItem
	id := c.Params("id")
	if err := oic.DB.First(&orderItem, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Order item not found",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve order item",
		})
	}
	return c.JSON(orderItem)
}

func (oic *OrderItemController) CreateOrderItem(c *fiber.Ctx) error {
	var orderItem models.OrderItem
	if err := c.BodyParser(&orderItem); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request data",
		})
	}

	if err := oic.DB.Create(&orderItem).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create order item",
		})
	}
	return c.Status(fiber.StatusCreated).JSON(orderItem)
}

func (oic *OrderItemController) UpdateOrderItem(c *fiber.Ctx) error {
	var orderItem models.OrderItem
	id := c.Params("id")
	if err := oic.DB.First(&orderItem, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Order item not found",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve order item",
		})
	}

	if err := c.BodyParser(&orderItem); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request data",
		})
	}

	if err := oic.DB.Save(&orderItem).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update order item",
		})
	}
	return c.JSON(orderItem)
}

func (oic *OrderItemController) DeleteOrderItem(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := oic.DB.Delete(&models.OrderItem{}, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Order item not found",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to delete order item",
		})
	}
	return c.SendStatus(fiber.StatusNoContent)
}
