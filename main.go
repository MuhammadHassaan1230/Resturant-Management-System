package main

import (
	"os"

	"github.com/MuhammadHassaan1230/restaurant-management/database"
	"github.com/gin-gonic/gin"
)

var collection = database.GetCollection(database.Client, "restaurantManagement", "food")

func main() {

	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	router := gin.New()
	router.Use(gin.Logger())

	//Routes Group
	//routes.UserRoute(router)
	//routes.FoodRouter(router)
	//routes.MenuRouter(router)
	//routes.TableRoute(router)
	//routes.OrderRoute(router)
	//routes.OrderItemRoute(router)
	//routes.InvoiceRoute(router)

	router.Run(":" + port)
}
