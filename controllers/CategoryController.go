package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/myustanzah/GO-Fiber.git/models"
	"gorm.io/gorm"
)

type CategoryController struct {
	DB *gorm.DB
}

func NewCategoryController(db *gorm.DB) *CategoryController {
	return &CategoryController{DB: db}
}

func (c *CategoryController) GetAllCategories(ctx *fiber.Ctx) error {
	var categories []models.Category
	if err := c.DB.Find(&categories).Error; err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve categories",
		})
	}
	return ctx.JSON(categories)
}

func (c *CategoryController) GetCategoryByID(ctx *fiber.Ctx) error {
	var category models.Category
	id := ctx.Params("id")
	if err := c.DB.First(&category, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Category not found",
			})
		}
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve category",
		})
	}
	return ctx.JSON(category)
}

func (c *CategoryController) CreateCategory(ctx *fiber.Ctx) error {
	var category models.Category
	if err := ctx.BodyParser(&category); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request data",
		})
	}

	if err := c.DB.Create(&category).Error; err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create category",
		})
	}
	return ctx.Status(fiber.StatusCreated).JSON(category)
}

func (c *CategoryController) UpdateCategory(ctx *fiber.Ctx) error {
	var category models.Category
	id := ctx.Params("id")
	if err := c.DB.First(&category, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Category not found",
			})
		}
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve category",
		})
	}

	if err := ctx.BodyParser(&category); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request data",
		})
	}

	if err := c.DB.Save(&category).Error; err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update category",
		})
	}
	return ctx.JSON(category)
}

func (c *CategoryController) DeleteCategory(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	if err := c.DB.Delete(&models.Category{}, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Category not found",
			})
		}
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to delete category",
		})
	}
	return ctx.SendStatus(fiber.StatusNoContent)
}

func (c *CategoryController) GetCategoryProducts(ctx *fiber.Ctx) error {
	var products []models.Product
	categoryID := ctx.Params("id")
	if err := c.DB.Where("category_id = ?", categoryID).Find(&products).Error; err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve products for category",
		})
	}
	if len(products) == 0 {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "No products found for this category",
		})
	}
	return ctx.JSON(products)
}

func (c *CategoryController) GetCategoryOrderItems(ctx *fiber.Ctx) error {
	var orderItems []models.OrderItem
	categoryID := ctx.Params("id")
	if err := c.DB.Where("category_id = ?", categoryID).Find(&orderItems).Error; err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve order items for category",
		})
	}
	if len(orderItems) == 0 {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "No order items found for this category",
		})
	}
	return ctx.JSON(orderItems)
}

func (c *CategoryController) GetCategoryOrders(ctx *fiber.Ctx) error {
	var orders []models.Order
	categoryID := ctx.Params("id")
	if err := c.DB.Where("category_id = ?", categoryID).Find(&orders).Error; err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve orders for category",
		})
	}
	if len(orders) == 0 {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "No orders found for this category",
		})
	}
	return ctx.JSON(orders)
}

func (c *CategoryController) GetCategoryOrderItemsByOrderID(ctx *fiber.Ctx) error {
	orderID := ctx.Params("order_id")
	var orderItems []models.OrderItem
	if err := c.DB.Where("order_id = ?", orderID).Find(&orderItems).Error; err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve order items for order",
		})
	}
	if len(orderItems) == 0 {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "No order items found for this order",
		})
	}
	return ctx.JSON(orderItems)
}

func (c *CategoryController) GetCategoryOrderByID(ctx *fiber.Ctx) error {
	orderID := ctx.Params("order_id")
	var order models.Order
	if err := c.DB.First(&order, orderID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Order not found",
			})
		}
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve order",
		})
	}
	return ctx.JSON(order)
}

func (c *CategoryController) GetCategoryOrderItemsByProductID(ctx *fiber.Ctx) error {
	productID := ctx.Params("product_id")
	var orderItems []models.OrderItem
	if err := c.DB.Where("product_id = ?", productID).Find(&orderItems).Error; err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve order items for product",
		})
	}
	if len(orderItems) == 0 {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "No order items found for this product",
		})
	}
	return ctx.JSON(orderItems)
}

func (c *CategoryController) GetCategoryOrderItemByID(ctx *fiber.Ctx) error {
	orderItemID := ctx.Params("order_item_id")
	var orderItem models.OrderItem
	if err := c.DB.First(&orderItem, orderItemID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Order item not found",
			})
		}
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve order item",
		})
	}
	return ctx.JSON(orderItem)
}

func (c *CategoryController) DeleteCategoryOrderItem(ctx *fiber.Ctx) error {
	orderItemID := ctx.Params("order_item_id")
	if err := c.DB.Delete(&models.OrderItem{}, orderItemID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Order item not found",
			})
		}
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to delete order item",
		})
	}
	return ctx.SendStatus(fiber.StatusNoContent)
}

func (c *CategoryController) DeleteCategoryOrder(ctx *fiber.Ctx) error {
	orderID := ctx.Params("order_id")
	if err := c.DB.Delete(&models.Order{}, orderID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Order not found",
			})
		}
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to delete order",
		})
	}
	return ctx.SendStatus(fiber.StatusNoContent)
}

func (c *CategoryController) DeleteCategoryProduct(ctx *fiber.Ctx) error {
	productID := ctx.Params("product_id")
	if err := c.DB.Delete(&models.Product{}, productID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Product not found",
			})
		}
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to delete product",
		})
	}
	return ctx.SendStatus(fiber.StatusNoContent)
}
