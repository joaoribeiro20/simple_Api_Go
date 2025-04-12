package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/joaoribeiro20/simplegoapi/config"
	"github.com/joaoribeiro20/simplegoapi/models"
)

func CreateBudget(c *gin.Context) {
	var budget models.Budget
	if err := c.ShouldBindJSON(&budget); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Create(&budget)
	c.JSON(http.StatusCreated, budget)
}

func GetBudgets(c *gin.Context) {
	var budgets []models.Budget
	config.DB.Find(&budgets)
	c.JSON(http.StatusOK, budgets)
}

func GetBudget(c *gin.Context) {
	var budget models.Budget
	if err := config.DB.First(&budget, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Budget não encontrado"})
		return
	}
	c.JSON(http.StatusOK, budget)
}

func UpdateBudget(c *gin.Context) {
	var budget models.Budget
	if err := config.DB.First(&budget, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Budget não encontrado"})
		return
	}
	if err := c.ShouldBindJSON(&budget); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Save(&budget)
	c.JSON(http.StatusOK, budget)
}

func DeleteBudget(c *gin.Context) {
	var budget models.Budget
	if err := config.DB.First(&budget, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Budget não encontrado"})
		return
	}
	config.DB.Delete(&budget)
	c.JSON(http.StatusNoContent, nil)
}
