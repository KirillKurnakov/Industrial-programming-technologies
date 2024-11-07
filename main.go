package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Item struct {
	ID            string  `json:"id"`
	Name          string  `json:"name"`
	Category_id   int     `json:"category_id"`
	Description   string  `json:"description"`
	Price         float64 `json:"price"`
	QuantityStock int     `json:"quantityStock"`
}

var items = []Item{
	{ID: "1", Name: "Кружка", Category_id: 1, Description: "Крутой подарок", Price: 45.0, QuantityStock: 5},
	{ID: "2", Name: "Тарелка", Category_id: 2, Description: "Крутой подарок", Price: 45.0, QuantityStock: 5},
	{ID: "3", Name: "Мягкая игрушка", Category_id: 3, Description: "Крутой подарок", Price: 45.0, QuantityStock: 5},
}

var itemsBasket = []Item{}

var router = gin.Default()

func main() {
	//router := gin.Default()

	// Получение всех товаров
	router.GET("/items", getItems)

	// Получение всех товаров из корзины
	router.GET("/id/1/basket", getItemsBasket)

	// Получение товара по ID
	router.GET("/items/:id", getItemByID)

	// Создание нового товара
	router.POST("/items", createItem)

	// Добавление товара в корзину
	router.POST("/id/1/basket", createItemBasket)

	// Обновление существующего товара
	router.PUT("/items/:id", updateItem)

	// Удаление товара
	router.DELETE("/items/:id", deleteItem)

	// Удаление товара из корзины
	router.DELETE("/id/1/basket", deleteItemBasket)

	router.Run(":8080")
}

func getItems(c *gin.Context) {
	c.JSON(http.StatusOK, items)
}

func getItemsBasket(c *gin.Context) {
	c.JSON(http.StatusOK, itemsBasket)
}

func getItemByID(c *gin.Context) {
	id := c.Param("id")

	for _, item := range items {
		if item.ID == id {
			c.JSON(http.StatusOK, item)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "item not found"})
}

func createItem(c *gin.Context) {
	var newItem Item

	if err := c.BindJSON(&newItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		return
	}

	items = append(items, newItem)
	c.JSON(http.StatusCreated, newItem)
}

func createItemBasket(c *gin.Context) {
    
    var body Item
    c.ShouldBindJSON(&body)
	var newItem Item
    id := body.ID

    for _, item := range items {
		if item.ID == id {
            newItem = item
			break
		}
	}

	if newItem.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		return
	}

	itemsBasket = append(itemsBasket, newItem)
	c.JSON(http.StatusOK, newItem)
}

func updateItem(c *gin.Context) {
	id := c.Param("id")
	var updatedItem Item

	if err := c.BindJSON(&updatedItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		return
	}

	for i, item := range items {
		if item.ID == id {
			items[i] = updatedItem
			c.JSON(http.StatusOK, updatedItem)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "item not found"})
}

func deleteItem(c *gin.Context) {
	id := c.Param("id")

	for i, item := range items {
		if item.ID == id {
			items = append(items[:i], items[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "item deleted"})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "item not found"})
}

func deleteItemBasket(c *gin.Context) {
	var body Item
    c.ShouldBindJSON(&body)
    id := body.ID

	for i, item := range itemsBasket {
		if item.ID == id {
			itemsBasket = append(itemsBasket[:i], itemsBasket[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "item from basket deleted"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "item in basket not found"})
}
