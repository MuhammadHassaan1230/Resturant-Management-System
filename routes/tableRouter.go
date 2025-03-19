package routes

import (
	"github.com/MuhammadHassaan1230/restaurant-management/controllers"
	"github.com/gin-gonic/gin"
)

func TableRoute(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/tables", controllers.GetTables)
	incomingRoutes.GET("/tables/:id", controllers.GetTable)
	incomingRoutes.POST("/tables", controllers.CreateTable)
	incomingRoutes.PATCH("/tables/:id", controllers.UpdateTable)
}
