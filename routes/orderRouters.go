package routers

import (
 "assignment-2/controllers"
 "github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
    router := gin.Default()
    router.POST("/new-orders", controllers.CreateOrder)
    router.GET("/orders", controllers.GetOrder)
    router.PUT("/orders", controllers.UpdateOrder)
    router.DELETE("/orders", controllers.DeleteOrder)
    return router
}