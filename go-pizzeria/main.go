package main

import (
	"encoding/json"
	"fmt"
	"os"
	"pizzeria/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

var pizzas []models.Pizza = []models.Pizza{}

func main() {
	loadPizzas()

	router := gin.Default()
	router.GET("/pizzas", getPizzas)
	router.GET("/pizzas/:id", getPizzaByID)
	router.POST("/pizzas", createPizza)

	router.Run(":8080")
}

func getPizzas(c *gin.Context) {
	c.JSON(200, gin.H{
		"pizzas": pizzas,
	})
}

func getPizzaByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Formato de ID inválido",
		})
		return
	}

	for _, pizza := range pizzas {
		if id == pizza.ID {
			c.JSON(200, gin.H{
				"pizza": pizza,
			})
			return
		}
	}
	c.JSON(404, gin.H{
		"error": "Pizza não encontrada",
	})
}

func createPizza(c *gin.Context) {
	var newPizza models.Pizza
	if err := c.BindJSON(&newPizza); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	newPizza.ID = len(pizzas) + 1
	pizzas = append(pizzas, newPizza)
	c.JSON(201, gin.H{
		"message": "Pizza created successfully",
		"data":    newPizza,
	})

	savePizzas()
}

func loadPizzas() {
	file, err := os.Open("data/pizza.json")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()
	decoder := json.NewDecoder(file)

	if err := decoder.Decode(&pizzas); err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}
}

func savePizzas() {
	file, err := os.Create("data/pizza.json")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	if err := encoder.Encode(pizzas); err != nil {
		fmt.Println("Error encoding JSON:", err)
		return
	}
}
