package routers

import (
	"assignment2/controllers"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	router.GET("/orders", controllers.GetAllOrders)
	router.POST("/orders", controllers.PostOrders)
	router.DELETE("/orders/:orderID", controllers.DeleteOrderById)
	router.PUT("/orders/:orderID", controllers.UpdateOrderByID)

	return router
}
