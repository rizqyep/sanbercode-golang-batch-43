package routers

import (
	"formative-13/controllers"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	router.POST("/cars", controllers.CreateCar)
	router.PUT("/cars/:carID", controllers.UpdateCar)
	router.GET("/cars/:carID", controllers.GetCar)
	router.DELETE("/cars/:carID", controllers.DeleteCar)
	return router
}
