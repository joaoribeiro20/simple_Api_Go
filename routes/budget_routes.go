package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/joaoribeiro20/simplegoapi/controllers"
)

func RegisterBudgetRoutes(r *gin.Engine) {
	budgets := r.Group("/budgets")
	{
		budgets.POST("/", controllers.CreateBudget)
		budgets.GET("/", controllers.GetBudgets)
		budgets.GET("/:id", controllers.GetBudget)
		budgets.PUT("/:id", controllers.UpdateBudget)
		budgets.DELETE("/:id", controllers.DeleteBudget)
	}
}
