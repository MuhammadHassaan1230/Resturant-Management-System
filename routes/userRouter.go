package routes

import (
	"github.com/MuhammadHassaan1230/restaurant-management/controllers"
	"github.com/gin-gonic/gin"
)

func UserRoute(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/users", controllers.GetUsers)
	incomingRoutes.GET("/users/:id", controllers.GetUser)
	incomingRoutes.POST("/signup", controllers.Signup)
	incomingRoutes.POST("/login", controllers.Login)
}
