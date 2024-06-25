package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type product struct {
	ID         string  `json:"id"`
	Name       string  `json:"name"`
	Quantity   int     `json:"quantity"`
	Unit_coast float64 `json:"unit_coast"`
	Measure    string  `json:"measure"`
}

var products = []product{
	{ID: "1", Name: "Fish", Quantity: 150, Unit_coast: 149.9, Measure: "4"},
	{ID: "2", Name: "bread", Quantity: 1200, Unit_coast: 25, Measure: "2"},
	{ID: "3", Name: "Apple", Quantity: 10, Unit_coast: 250, Measure: "1"},
	{ID: "4", Name: "RandomProd", Quantity: 100, Unit_coast: 1, Measure: "0"},
	{ID: "5", Name: "RandomProd", Quantity: 100, Unit_coast: 1, Measure: "0"},
	{ID: "6", Name: "RandomProd", Quantity: 100, Unit_coast: 1, Measure: "0"},
	{ID: "7", Name: "RandomProd", Quantity: 100, Unit_coast: 1, Measure: "0"},
	{ID: "8", Name: "RandomProd", Quantity: 100, Unit_coast: 1, Measure: "0"},
}

func main() {
	router := gin.Default()
	router.GET("/products", getProducts)          //Получить список продуктов
	router.GET("/products/:id", getProductByID)   //Получить один продукт
	router.POST("/products", postProducts)        //Создать новый продукт
	router.PUT("/products/:id", putProduct)       // Изменить продукт
	router.DELETE("/products/:id", deleteProduct) // Удалить продукт

	router.GET("/measures/", getMeasureList)      // Получить список полей measure продуктов.
	router.GET("/measures/:id", getMeasureByID)   // Получить поле Measure продукта по id id
	router.PUT("/measures/:id", putMeasure)       // Изменить значение поля measure продукта
	router.DELETE("/measures/:id", deleteMeasure) // Установить значение продукта measure по умолчанию

	router.Run("localhost:3000")
}

func getProducts(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, products)
}

func postProducts(c *gin.Context) {
	var newProduct product

	if err := c.BindJSON(&newProduct); err != nil {
		return
	}

	products = append(products, newProduct)
	c.IndentedJSON(http.StatusCreated, newProduct)
}

func getProductByID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range products {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "product not found"})
}
func putProduct(c *gin.Context) {
	id := c.Param("id")
	var updatedProduct product

	if err := c.BindJSON(&updatedProduct); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "problem with request data"})
		return
	}
	for i, a := range products {
		if a.ID == id {
			products[i] = updatedProduct // Замена существующего альбома новым.
			c.IndentedJSON(http.StatusOK, updatedProduct)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Product not found"})
}

func deleteProduct(c *gin.Context) {
	id := c.Param("id")

	for i, a := range products {
		if a.ID == id {
			products = append(products[:i], products[i+1:]...)
			c.IndentedJSON(http.StatusOK, gin.H{"message": "product deleted"})
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "product not found"})
}

// Получить список Measure всех продуктов:
func getMeasureList(c *gin.Context) {
	var measures []string
	for _, a := range products {
		measures = append(measures, a.Measure)
	}
	c.IndentedJSON(http.StatusOK, measures)
}

func getMeasureByID(c *gin.Context) {
	id := c.Param("id")
	for _, a := range products {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a.Measure)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "product not found"})
}

func putMeasure(c *gin.Context) {
	id := c.Param("id")
	var newMeasure struct {
		Measure string `json:"measure"`
	}

	if err := c.BindJSON(&newMeasure); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "problem with request data"})
		return
	}

	for i, a := range products {
		if a.ID == id {
			products[i].Measure = newMeasure.Measure // Обновление цены продукта.
			c.IndentedJSON(http.StatusOK, products[i])
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "product not found"})
}
func deleteMeasure(c *gin.Context) {
	id := c.Param("id")

	// Ищем продукт по ID и присваиваем его измерению значение 0.
	for i, a := range products {
		if a.ID == id {
			products[i].Measure = "0" // Установка цены продукта в 0.
			c.IndentedJSON(http.StatusOK, products[i])
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "product not found"})
}
