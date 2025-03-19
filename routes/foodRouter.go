package routes

import (
	"github.com/MuhammadHassaan1230/restaurant-management/controllers"
	"github.com/gin-gonic/gin"
)

func FoodRouter(incomingroutes *gin.Engine) {
	incomingroutes.GET("/food", controllers.GetFoods)
	incomingroutes.GET("/food/:id", controllers.GetFood)
	incomingroutes.POST("/food", controllers.CreateFood)
	incomingroutes.PATCH("/food/:id", controllers.UpdateFood)
}
