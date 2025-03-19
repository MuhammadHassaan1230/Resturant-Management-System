package routes

import (
	"github.com/MuhammadHassaan1230/restaurant-management/controllers"
	"github.com/gin-gonic/gin"
)

func InvoiceRoute(IncomingRoutes *gin.Engine) {
	IncomingRoutes.GET("/invoices", controllers.GetInvoices)
	IncomingRoutes.GET("/invoices/:id", controllers.GetInvoice)
	IncomingRoutes.POST("/invoices", controllers.CreateInvoice)
	IncomingRoutes.PATCH("/invoices/:id", controllers.UpdateInvoice)
}
