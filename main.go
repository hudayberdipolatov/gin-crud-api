package main

import (
	"gorm_with_mysql_and_gin/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := setupRouter()
	_ = router.Run(":8080")
}

func setupRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, "pong")
	})

	// user routes
	userRepo := controllers.New()
	router.POST("/users", userRepo.CreateUser)
	router.GET("/users", userRepo.GetUsers)
	router.GET("/users/:id", userRepo.GetUser)
	router.PUT("/users/:id", userRepo.UpdateUser)
	router.DELETE("/users/:id", userRepo.DeleteUser)
	return router
}
