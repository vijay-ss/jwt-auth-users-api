package routes

import (
	controller "github.com/vijay-ss/jwt-auth-users-api/controllers"
	"github.com/vijay-ss/jwt-auth-users-api/middleware"
	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.GET("/users", controller.GetUsers())
	incomingRoutes.GET("/users/:user_id", controller.GetUser())

}