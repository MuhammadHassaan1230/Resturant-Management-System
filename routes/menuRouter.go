package routes

import (
	"github.com/MuhammadHassaan1230/restaurant-management/controllers"
	"github.com/gin-gonic/gin"
)

func MenuRouter(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/menus", controllers.GetMenus)
	incomingRoutes.GET("/menus/:id", controllers.GetMenu)
	incomingRoutes.POST("/menus", controllers.CreateMenu)
	incomingRoutes.PATCH("/menus/:id", controllers.UpdateMenu)
}
