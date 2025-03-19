package routes

import (
	"github.com/MuhammadHassaan1230/restaurant-management/controllers"
	"github.com/gin-gonic/gin"
)

func OrderItemRoute(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/order-items", controllers.GetOrderItems)
	incomingRoutes.GET("/order-items/:id", controllers.GetOrderItem)
	incomingRoutes.POST("/order-items", controllers.CreateOrderItem)
	incomingRoutes.PATCH("/order-items/:id", controllers.UpdateOrderItem)
}
