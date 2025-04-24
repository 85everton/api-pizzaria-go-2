package handler

import (
	"fmt"
	"net/http"
	"pizzaria/internal/data"
	"pizzaria/internal/models"
	Service "pizzaria/internal/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetPizzas(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"pizzas": data.Pizzas,
	})
}

func PostPizzas(c *gin.Context) {
	var newPizza models.Pizza
	if err := c.ShouldBindJSON(&newPizza); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"erro": err.Error()})
		return
	}

	if err := Service.ValidatePizzaPrice(&newPizza); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"erro": err.Error()})
		return
	}

	newPizza.ID = len(data.Pizzas) + 1
	data.Pizzas = append(data.Pizzas, newPizza)
	data.SavePizzas()
	fmt.Println(newPizza)
	c.JSON(http.StatusCreated, newPizza)
}

func GetPizzasByID(c *gin.Context) {
	idPraram := c.Param("id")
	id, err := strconv.Atoi(idPraram)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"erro": err.Error()})
		return
	}

	for _, p := range data.Pizzas {
		if p.ID == id {
			c.JSON(http.StatusCreated, p)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "Pizza não encontrada"})
}

func DeletePizzaById(c *gin.Context) {
	idPraram := c.Param("id")
	id, err := strconv.Atoi(idPraram)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"erro": err.Error()})
		return
	}
	for i, p := range data.Pizzas {
		if p.ID == id {
			data.Pizzas = append(data.Pizzas[:i], data.Pizzas[i+1:]...)
			data.SavePizzas()
			c.JSON(http.StatusOK, gin.H{"message": "Pizza deletada com sucesso"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"method": "piza não encontrada"})

}

func UpdatePizzaById(c *gin.Context) {
	idPraram := c.Param("id")
	id, err := strconv.Atoi(idPraram)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"erro": err.Error()})
		return

	}
	var updatedPizza models.Pizza
	if err := c.ShouldBindJSON(&updatedPizza); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"erro": err.Error()})
		return
	}

	if err := Service.ValidatePizzaPrice(&updatedPizza); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"erro": err.Error()})
		return
	}

	for i, p := range data.Pizzas {
		if p.ID == id {
			data.Pizzas[i] = updatedPizza
			data.Pizzas[i].ID = id
			data.SavePizzas()
			c.JSON(http.StatusOK, data.Pizzas[i])
			return
		}

	}

	c.JSON(http.StatusNotFound, gin.H{"method": "pizza não encontrada"})

}
